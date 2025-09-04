package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"

	"oldoldcoder.com/cloudshop/product-service/internal/config"
	"oldoldcoder.com/cloudshop/product-service/internal/database"
	"oldoldcoder.com/cloudshop/product-service/internal/middleware"
	"oldoldcoder.com/cloudshop/product-service/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

func getenv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}

func main() {
	// 获取可执行文件所在目录，然后查找配置文件
	execPath, _ := os.Executable()
	execDir := filepath.Dir(execPath)

	// 尝试多个可能的配置文件路径
	configPaths := []string{
		"config/config.yaml",                            // 开发环境：相对于项目根目录
		filepath.Join(execDir, "../config/config.yaml"), // 打包后：相对于可执行文件
		filepath.Join(execDir, "config.yaml"),           // 配置文件在可执行文件同目录
	}

	var cfg *config.Config
	var err error
	for _, configPath := range configPaths {
		cfg, err = config.LoadConfig(configPath)
		if err == nil {
			log.Printf("Loaded config from: %s", configPath)
			break
		}
	}

	if err != nil {
		log.Fatalf("Failed to load config from any path: %v", err)
	}
	// 若启用 Nacos，再从远程覆盖
	if cfg.Nacos.Enabled {
		log.Printf("Nacos enabled, trying to load remote config...")
		log.Printf("Nacos config: server=%s, namespace=%s, group=%s, dataId=%s",
			cfg.Nacos.ServerAddr, cfg.Nacos.Namespace, cfg.Nacos.Group, cfg.Nacos.DataId)

		// 解析 server_addr，支持 localhost:8848 格式
		serverAddr := cfg.Nacos.ServerAddr
		if serverAddr == "localhost:8848" {
			serverAddr = "127.0.0.1:8848" // Go SDK 需要 IP 地址
		}

		remote, rerr := config.LoadConfigFromNacos(
			serverAddr,
			cfg.Nacos.Username,
			cfg.Nacos.Password,
			cfg.Nacos.Namespace,
			cfg.Nacos.Group,
			cfg.Nacos.DataId,
		)
		if rerr != nil {
			log.Printf("Nacos config failed: %v, using local config", rerr)
		} else {
			log.Printf("Successfully loaded config from Nacos")
			// 只覆盖业务配置，保留本地的 nacos 和 service 配置
			originalNacos := cfg.Nacos
			originalService := cfg.Service
			cfg = remote
			cfg.Nacos = originalNacos     // 保留本地 nacos 配置
			cfg.Service = originalService // 保留本地 service 配置
		}
	} else {
		log.Printf("Nacos disabled, using local config only")
	}

	// 打印关键配置信息
	log.Printf("Using config - Server: %s:%s, Database: %s:%s, Service: %s",
		cfg.Server.Host, cfg.Server.Port, cfg.Database.Host, cfg.Database.Port, cfg.Service.Name)

	// 初始化数据库等
	log.Printf("Initializing database connection...")
	if err := database.InitDatabase(cfg); err != nil {
		log.Fatalf("DB init failed: %v", err)
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(middleware.Logger(), middleware.CORS())

	routes.SetupRoutes(r, cfg.JWT.Secret)
	r.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"status": "ok", "service": "product-service"}) })

	if cfg.Server.Host == "" {
		cfg.Server.Host = "0.0.0.0"
	}
	addr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)

	// http.Server 支持优雅关闭
	srv := &http.Server{Addr: addr, Handler: r}

	// 可选：注册到 Nacos（通过配置开关控制）
	shouldRegister := cfg.Service.Register
	serviceName := cfg.Service.Name
	serviceIP := cfg.Service.IP
	if serviceIP == "" || serviceIP == "127.0.0.1" || serviceIP == "localhost" {
		if ip := detectServiceIP(); ip != "" {
			serviceIP = ip
		}
	}
	servicePort := cfg.Server.Port

	var namingClient naming_client.INamingClient
	registered := false
	if shouldRegister {
		// 解析 serverAddr，支持 localhost:8848 格式
		serverAddr := cfg.Nacos.ServerAddr
		ip := serverAddr
		port := uint64(8848)
		if strings.Contains(serverAddr, ":") {
			parts := strings.Split(serverAddr, ":")
			ip = parts[0]
			if p, err := strconv.ParseUint(parts[1], 10, 64); err == nil {
				port = p
			}
		}

		serverConfs := []constant.ServerConfig{{IpAddr: ip, Port: port}}
		clientConf := constant.ClientConfig{
			NamespaceId:  cfg.Nacos.Namespace,
			Username:     cfg.Nacos.Username,
			Password:     cfg.Nacos.Password,
			TimeoutMs:    10000,
			BeatInterval: 5000,
			LogLevel:     "info",
		}
		cli, nerr := clients.NewNamingClient(vo.NacosClientParam{ClientConfig: &clientConf, ServerConfigs: serverConfs})
		if nerr != nil {
			log.Printf("Nacos naming client init failed: %v", nerr)
		} else {
			namingClient = cli

			// 等待客户端连接完成
			log.Printf("Waiting for Nacos client to connect...")
			time.Sleep(2 * time.Second)

			p, _ := strconv.Atoi(servicePort)
			reg := vo.RegisterInstanceParam{
				Ip:          serviceIP,
				Port:        uint64(p),
				ServiceName: serviceName,
				GroupName:   cfg.Nacos.Group,
				Enable:      true,
				Healthy:     true,
				Weight:      100,
				Ephemeral:   true,
			}
			if ok, rerr := namingClient.RegisterInstance(reg); rerr != nil || !ok {
				log.Printf("Nacos register failed: %v", rerr)
			} else {
				registered = true
				log.Printf("Registered to Nacos: %s %s:%s", serviceName, serviceIP, servicePort)
			}
		}
	}

	// 启动服务
	go func() {
		log.Printf("Product service starting on %s", addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %v", err)
		}
	}()

	// 等待退出信号，优雅关停
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Server shutdown error: %v", err)
	}

	// 反注册 Nacos 实例
	if shouldRegister && registered && namingClient != nil {
		p, _ := strconv.Atoi(servicePort)
		_, derr := namingClient.DeregisterInstance(vo.DeregisterInstanceParam{
			Ip:          serviceIP,
			Port:        uint64(p),
			ServiceName: serviceName,
			GroupName:   cfg.Nacos.Group,
			Ephemeral:   true,
		})
		if derr != nil {
			log.Printf("Nacos deregister failed: %v", derr)
		} else {
			log.Println("Deregistered from Nacos")
		}
	}

	log.Println("Server exiting")
}

// detectServiceIP 返回一个非环回的局域网 IP，优先选择 192.168.x.x；
// 若不存在，再选择任意私网地址（10.x 或 172.16-31.x）。找不到则返回空串。
func detectServiceIP() string {
	var fallback string
	ifaces, err := net.Interfaces()
	if err != nil {
		return ""
	}
	for _, iface := range ifaces {
		if (iface.Flags & net.FlagUp) == 0 {
			continue
		}
		addrs, _ := iface.Addrs()
		for _, a := range addrs {
			var ip net.IP
			switch v := a.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue
			}
			// 优先匹配 192.168.x.x
			if ip[0] == 192 && ip[1] == 168 {
				return ip.String()
			}
			// 记录一个可用的私网地址作为备选
			if fallback == "" {
				if ip[0] == 10 || (ip[0] == 172 && ip[1] >= 16 && ip[1] <= 31) {
					fallback = ip.String()
				}
			}
		}
	}
	return fallback
}
