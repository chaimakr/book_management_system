package utils

import (
	"context"
	"net/http"
	"os"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

func GetRequestID(r *http.Request) string {
	// No request id found.
	hv := r.Header.Get("X-Request-ID")
	if hv != "" {
		return hv
	}

	return uuid.New().String()
}

type Service struct {
	Logger *zap.SugaredLogger
	Tracer trace.Tracer
}

func BuildTracer() trace.Tracer {
	// Environment variables

	traceExporter, err := otlptracegrpc.New(context.Background())
	if err != err {
		os.Exit(1)
	}

	provider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()), // 100% of the traces
		sdktrace.WithBatcher(traceExporter),           //
	)

	otel.SetTracerProvider(provider)

	return otel.Tracer("fact-service")
}

func BuildLogger() *zap.SugaredLogger {
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)

	l, _ := config.Build()
	return l.Sugar()
}
