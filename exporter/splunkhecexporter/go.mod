module github.com/open-telemetry/opentelemetry-collector-contrib/exporter/splunkhecexporter

go 1.15

require (
	github.com/census-instrumentation/opencensus-proto v0.3.0
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/splunk v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.7.0
	go.opentelemetry.io/collector v0.26.1-0.20210517223805-ce829467ebea
	go.uber.org/zap v1.16.0
	google.golang.org/protobuf v1.26.0

)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/splunk => ../../internal/splunk
