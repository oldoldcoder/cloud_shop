package com.cloudshop.gateway.filter;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.cloud.gateway.filter.GatewayFilterChain;
import org.springframework.cloud.gateway.filter.GlobalFilter;
import org.springframework.core.Ordered;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpStatusCode;
import org.springframework.http.server.reactive.ServerHttpRequest;
import org.springframework.http.server.reactive.ServerHttpResponse;
import org.springframework.stereotype.Component;
import org.springframework.web.server.ServerWebExchange;
import reactor.core.publisher.Mono;

import java.net.URI;
import java.time.Instant;

@Component
public class RequestLogFilter implements GlobalFilter, Ordered {
    private static final Logger log = LoggerFactory.getLogger(RequestLogFilter.class);

    @Override
    public Mono<Void> filter(ServerWebExchange exchange, GatewayFilterChain chain) {
        ServerHttpRequest request = exchange.getRequest();
        URI uri = request.getURI();
        HttpHeaders headers = request.getHeaders();
        String method = request.getMethod() != null ? request.getMethod().name() : "UNKNOWN";
        String path = uri.getRawPath();
        String query = uri.getRawQuery();
        String trace = headers.getFirst("X-Request-Id");
        long start = System.currentTimeMillis();

        log.info("GW-REQ method={} path={}{} time={} traceId={} ua={}",
                method,
                path,
                (query != null && !query.isEmpty() ? ("?" + query) : ""),
                Instant.ofEpochMilli(start),
                trace,
                headers.getFirst(HttpHeaders.USER_AGENT));

        return chain.filter(exchange)
                .doOnSuccess(v -> logResponse(exchange, start, null))
                .doOnError(err -> logResponse(exchange, start, err));
    }

    private void logResponse(ServerWebExchange exchange, long start, Throwable err) {
        ServerHttpResponse response = exchange.getResponse();
        HttpStatusCode statusCode = response.getStatusCode();
        int status = statusCode != null ? statusCode.value() : 0;
        long cost = System.currentTimeMillis() - start;
        String path = exchange.getRequest().getPath().value();
        if (err == null) {
            log.info("GW-RESP status={} costMs={} path={}", status, cost, path);
        } else {
            log.warn("GW-RESP-ERR status={} costMs={} path={} err={}", status, cost, path, err.toString());
        }
    }

    @Override
    public int getOrder() {
        return Ordered.LOWEST_PRECEDENCE;
    }
}
