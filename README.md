# Monitoring Metrics

### Setup Go App

- Run app: `make run`
- Open URL: `http://localhost:5000/`


### Setup Prometheus

1. Start Prometheus
```
docker compose up -d prometheus
docker-compose logs -f prometheus
```

2. To verify it works: 
```
http://localhost:9090/-/ready

http://localhost:9090/metrics
```


### Prometheus Web UI

1. Graph dashboard: `http://localhost:9090/graph`

- Search for
```
prometheus_http_requests_total{code="200",handler="/metrics"}
prometheus_http_requests_total{code="200",handler="/api/v1/alerts"}
```

2. Alerts: `http://localhost:9090/alerts`

3. Configuration: `http://localhost:9090/config`

4. Rules: `http://localhost:9090/rules`

5. Targets: `http://localhost:9090/targets`

6. Service Discovery: `http://localhost:9090/service-discovery`

7. APIs query
- GET `http://localhost:9090/api/v1/query?query={__name__=~'prom.%2B'}`



### Metric names and labels

- Metrics URL: `http://localhost:9090/metrics`
- Graph URL: `http://localhost:9090/graph`

1. TYPE `Counter`

- HELP `prometheus_http_requests_total` Counter of HTTP requests
```
prometheus_http_requests_total{code="200",handler="/metrics"} 67
prometheus_http_requests_total{code="200",handler="/api/v1/alerts"} 0
```

- HELP `promhttp_metric_handler_requests_total` Total number of scrapes by HTTP status code
```
promhttp_metric_handler_requests_total{code="200"} 91
promhttp_metric_handler_requests_total{code="500"} 0
promhttp_metric_handler_requests_total{code="503"} 0
```

- HELP `process_cpu_seconds_total` Total user and system CPU time spent in seconds
```
process_cpu_seconds_total 1.61
```

- HELP `prometheus_sd_kubernetes_events_total` The number of Kubernetes events handled
```
prometheus_sd_kubernetes_events_total{event="add",role="pod"} 0
prometheus_sd_kubernetes_events_total{event="delete",role="pod"} 0
```


2. TYPE `Gauge`

- HELP `prometheus_build_info` A metric with a constant '1' value labeled by version, revision, branch, goversion from which prometheus was built, and the goos and goarch for the build
```
prometheus_build_info{branch="HEAD",goarch="arm64",goos="linux",goversion="go1.22.3",revision="879d80922a227c37df502e7315fad8ceb10a986d",tags="netgo,builtinassets,stringlabels",version="2.52.0"} 1
```

- HELP `prometheus_target_metadata_cache_bytes` The number of bytes that are currently used for storing metric metadata in the cache
```
prometheus_target_metadata_cache_bytes{scrape_job="prometheus"} 11459
```

3. TYPE `Histogram`

- HELP `prometheus_http_request_duration_seconds` Histogram of latencies for HTTP requests
```
prometheus_http_request_duration_seconds_bucket{handler="/-/ready",le="0.1"} 4

prometheus_http_request_duration_seconds_bucket{handler="/alerts",le="0.1"} 1

prometheus_http_request_duration_seconds_bucket{handler="/graph",le="0.1"} 2

prometheus_http_request_duration_seconds_bucket{handler="/metrics",le="0.1"} 471
```

4. TYPE `Summary`

- HELP `prometheus_engine_query_duration_seconds` Query timings
```
prometheus_engine_query_duration_seconds{slice="inner_eval",quantile="0.5"} 5.275e-05
prometheus_engine_query_duration_seconds{slice="inner_eval",quantile="0.9"} 0.000433875
prometheus_engine_query_duration_seconds{slice="inner_eval",quantile="0.99"} 0.0015605
prometheus_engine_query_duration_seconds_sum{slice="inner_eval"} 0.011728418
prometheus_engine_query_duration_seconds_count{slice="inner_eval"} 53
```


### Setup Node Exporter

1. Start Node Exporter
```
docker compose up -d node-exporter
docker-compose logs -f node-exporter
```

2. To verify it works: 
```
http://localhost:9100/metrics
```

### Setup Grafana

1. Start Grafana
```
docker compose up -d grafana
docker-compose logs -f grafana
```

2. To verify it works: 
```
http://localhost:3100
```

### References
- [Metric types](https://prometheus.io/docs/concepts/metric_types/)

- [Node Exporter](https://prometheus.io/docs/guides/node-exporter/)
