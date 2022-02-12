module github.com/layasugar/otel-collector

go 1.16

require (
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da
	github.com/google/uuid v1.3.0
	github.com/jaegertracing/jaeger v1.29.0
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal v0.41.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/jaeger v0.41.0
	github.com/stretchr/testify v1.7.0
	go.opencensus.io v0.23.0
	go.opentelemetry.io/collector v0.41.0
	go.opentelemetry.io/collector/model v0.41.0
	go.uber.org/zap v1.19.1
	google.golang.org/grpc v1.42.0
)
