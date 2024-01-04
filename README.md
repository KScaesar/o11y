# o11y

## prometheus

<https://prometheus.io/docs/introduction/first_steps/>

<https://prometheus.io/docs/prometheus/latest/command-line/prometheus/>

### Notation

```
<metric name>{<label name>=<label value>, ...}

api_http_requests_total{method="POST", handler="/messages"}
```

<https://prometheus.io/docs/concepts/data_model/#notation>

### metric example:

navigate to `http://localhost:9090/metrics`  

```text
# HELP go_gc_duration_seconds A summary of the pause duration of garbage collection cycles.
# TYPE go_gc_duration_seconds summary
go_gc_duration_seconds{quantile="0"} 5.3526e-05
go_gc_duration_seconds{quantile="0.25"} 9.2524e-05
go_gc_duration_seconds{quantile="0.5"} 0.000137984
go_gc_duration_seconds{quantile="0.75"} 0.00020063
go_gc_duration_seconds{quantile="1"} 0.020772341
go_gc_duration_seconds_sum 0.059429859
go_gc_duration_seconds_count 165

# HELP go_goroutines Number of goroutines that currently exist.
# TYPE go_goroutines gauge
go_goroutines 32

# HELP prometheus_engine_query_duration_seconds Query timings
# TYPE prometheus_engine_query_duration_seconds summary
prometheus_engine_query_duration_seconds{slice="inner_eval",quantile="0.5"} NaN
prometheus_engine_query_duration_seconds{slice="inner_eval",quantile="0.9"} NaN
prometheus_engine_query_duration_seconds{slice="inner_eval",quantile="0.99"} NaN
prometheus_engine_query_duration_seconds_sum{slice="inner_eval"} 0.000325868

# HELP prometheus_http_requests_total Counter of HTTP requests.
# TYPE prometheus_http_requests_total counter
prometheus_http_requests_total{code="200",handler="/-/ready"} 3
prometheus_http_requests_total{code="200",handler="/api/v1/label/:name/values"} 7
prometheus_http_requests_total{code="200",handler="/api/v1/query"} 9
prometheus_http_requests_total{code="200",handler="/api/v1/query_range"} 2
```

### METRIC TYPES

- Counter
- Gauge
- Histogram
- Summary

<https://prometheus.io/docs/concepts/metric_types/>

<https://prometheus.io/docs/tutorials/understanding_metric_types/>

### JOBS AND INSTANCES

In Prometheus terms, an endpoint you can scrape is called an instance,  
usually corresponding to a single process.  

For example:  
- job: api-server
    - instance 1: 1.2.3.4:5670
    - instance 2: 1.2.3.4:5671
    - instance 3: 5.6.7.8:5670
    - instance 4: 5.6.7.8:5671

<https://prometheus.io/docs/concepts/jobs_instances/>

### FEATURE FLAGS

```
# Prometheus agent
prometheus --config.file=prometheus.yml --web.listen-address="0.0.0.0:9090" --enable-feature=agent
```

<https://prometheus.io/docs/prometheus/latest/feature_flags/#prometheus-agent>

### reload config

```
curl -XPOST monitor.example.com:9090/-/reload
```

<https://prometheus.io/docs/alerting/latest/management_api/#reload>

### file-based service discovery

When using Prometheus' file-based service discovery mechanism, the Prometheus instance will listen for changes to the file and automatically update the scrape target list, without requiring an instance restart.


在生產環境中，當你需要改變 Prometheus 伺服器中的 file_sd（文件服務發現）目標文件時，確保這種更改是原子的，以避免重新載入時出現錯誤。

最佳的方法是在一個單獨的位置創建更新的文件，然後將其重命名為目標文件名（使用 mv 命令或 rename() 系統調用）。

<https://prometheus.io/docs/guides/file-sd/#use-file-based-service-discovery-to-discover-scrape-targets>

<https://prometheus.io/docs/prometheus/latest/http_sd/>