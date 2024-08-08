package tracing

import (
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

// ErrorTracer is
func ErrorTracer(span trace.Span, err error) {
	span.RecordError(err)
}

// EventErrorTracer is
func EventErrorTracer(span trace.Span, err error, name string) {
	span.RecordError(err)
	span.SetStatus(codes.Error, err.Error())
	span.AddEvent(name, trace.WithAttributes(
		attribute.String("error", err.Error()),
	))
}
