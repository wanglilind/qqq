package tracing

import (
	"io"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"github.com/wanglilind/qqq/pkg/logger"
)

type Tracer struct {
	closer io.Closer
}

func NewTracer(serviceName string, jaegerEndpoint string) (*Tracer, error) {
	cfg := &config.Configuration{
			ServiceName: serviceName,
			Sampler: &config.SamplerConfig{
					Type:  "const",
					Param: 1,
			},
			Reporter: &config.ReporterConfig{
					LogSpans:           true,
					LocalAgentHostPort: jaegerEndpoint,
			},
	}

	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
			return nil, err
	}

	opentracing.SetGlobalTracer(tracer)

	return &Tracer{
			closer: closer,
	}, nil
}

func (t *Tracer) Close() error {
	return t.closer.Close()
} 
