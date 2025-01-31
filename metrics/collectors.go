package metrics

import "context"

type Collectors struct {
	PapaStreamCollector *PapaStreamCollector
	HttpClientCollector *HttpClientCollector
}

func NewMetricsCollector(ctx context.Context, reg Registry) *Collectors {
	papaStreamCollector := NewPapaStreamCollector(ctx, reg)
	// httpClientCollector := NewHttpClientCollector(ctx, reg)
	return &Collectors{
		PapaStreamCollector: papaStreamCollector,
		// HttpClientCollector: httpClientCollector,
	}
}
