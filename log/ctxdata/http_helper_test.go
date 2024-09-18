package ctxdata

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/metadata"
)

func Test_SetContextAndMetadataFromHTTP(t *testing.T) {
	type args struct {
		ctx          context.Context
		req          *http.Request
		gcpProjectID string
		cIPs         []ClientIP
	}
	type out struct {
		ctx context.Context
		md  metadata.MD
	}

	tests := []struct {
		name string
		args args
		out  out
	}{
		{
			name: "Given complete values, when setting context and metadata, should return a complete context and metadata",
			args: args{
				ctx: context.Background(),
				req: &http.Request{
					Header: map[string][]string{
						headerXCorrelationId: []string{"correlationId"},
						headerTraceparent:    []string{"trace-parent"},
						headerTrace:          []string{"105445aa7843bc8bf206b120001000/1;o=1"},
						headerXForwardedFor:  []string{"127.0.0.1"},
						"User-Agent":         []string{"unittesting"},
					},
					Host:       "http://example.net",
					RemoteAddr: "127.0.0.1",
					RequestURI: "http://example.net",
				},
				gcpProjectID: "amartha-local",
				cIPs:         []ClientIP{},
			},
			out: out{
				ctx: context.WithValue(
					context.WithValue(
						context.WithValue(
							context.WithValue(
								context.WithValue(
									context.WithValue(
										context.WithValue(
											context.WithValue(
												context.WithValue(
													context.WithValue(
														context.Background(),
														correlationIdKey{},
														"correlationId",
													),
													traceParentKey{},
													"trace-parent",
												),
												traceIDKey{},
												"projects/amartha-local/traces/105445aa7843bc8bf206b120001000",
											),
											spanIDKey{},
											"1",
										),
										traceSampledKey{},
										true,
									),
									userAgentKey{},
									"unittesting",
								),
								hostKey{},
								"http://example.net",
							),
							ipKey{},
							"127.0.0.1",
						),
						forwardedForKey{},
						"127.0.0.1",
					),
					pidKey{},
					fmt.Sprintf("%d", os.Getpid()),
				),
				md: map[string][]string{
					correlationIdMDKey: {"correlationId"},
					traceParentMDKey:   {"trace-parent"},
					traceIDMDKey:       {"projects/amartha-local/traces/105445aa7843bc8bf206b120001000"},
					spanIDMDKey:        {"1"},
					traceSampledMDKey:  {"true"},
					userAgentMDKey:     {"unittesting"},
					hostMDKey:          {"http://example.net"},
					ipMDKey:            {"127.0.0.1"},
					forwardedForMDKey:  {"127.0.0.1"},
					pidMDKey:           {fmt.Sprintf("%d", os.Getpid())},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, md := SetContextAndMetadataFromHTTP(tt.args.ctx, tt.args.req, tt.args.gcpProjectID, tt.args.cIPs...)

			assert.Equal(t, tt.out.ctx, ctx)
			assert.Equal(t, tt.out.md, md)
		})
	}
}

func Test_SetContextFromGRPC(t *testing.T) {
	type args struct {
		ctx          context.Context
		gcpProjectID string
	}
	type out struct {
		ctx context.Context
	}

	tests := []struct {
		name string
		args args
		out  out
	}{
		{
			name: "Given complete grpc metadata, when setting context, should return complete context",
			args: args{
				ctx: metadata.NewIncomingContext(context.Background(), map[string][]string{
					correlationIdMDKey: {"correlationId"},
					traceParentMDKey:   {"trace-parent"},
					traceIDMDKey:       {"projects/amartha-local/traces/105445aa7843bc8bf206b120001000"},
					spanIDMDKey:        {"1"},
					traceSampledMDKey:  {"true"},
					userAgentMDKey:     {"unittesting"},
					hostMDKey:          {"http://example.net"},
					ipMDKey:            {"127.0.0.1"},
					forwardedForMDKey:  {"127.0.0.1"},
					pidMDKey:           {fmt.Sprintf("%d", os.Getpid())},
				}),
				gcpProjectID: "amartha-local",
			},
			out: out{
				ctx: context.WithValue(
					context.WithValue(
						context.WithValue(
							context.WithValue(
								context.WithValue(
									context.WithValue(
										context.WithValue(
											context.WithValue(
												context.WithValue(
													context.WithValue(
														metadata.NewIncomingContext(context.Background(), map[string][]string{
															correlationIdMDKey: {"correlationId"},
															traceParentMDKey:   {"trace-parent"},
															traceIDMDKey:       {"projects/amartha-local/traces/105445aa7843bc8bf206b120001000"},
															spanIDMDKey:        {"1"},
															traceSampledMDKey:  {"true"},
															userAgentMDKey:     {"unittesting"},
															hostMDKey:          {"http://example.net"},
															ipMDKey:            {"127.0.0.1"},
															forwardedForMDKey:  {"127.0.0.1"},
															pidMDKey:           {fmt.Sprintf("%d", os.Getpid())},
														}),
														correlationIdKey{},
														"correlationId",
													),
													traceParentKey{},
													"trace-parent",
												),
												traceIDKey{},
												"projects/amartha-local/traces/105445aa7843bc8bf206b120001000",
											),
											spanIDKey{},
											"1",
										),
										traceSampledKey{},
										true,
									),
									userAgentKey{},
									"unittesting",
								),
								hostKey{},
								"http://example.net",
							),
							ipKey{},
							"127.0.0.1",
						),
						forwardedForKey{},
						"127.0.0.1",
					),
					pidKey{},
					fmt.Sprintf("%d", os.Getpid()),
				),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := SetContextFromGRPC(tt.args.ctx, tt.args.gcpProjectID)

			assert.Equal(t, tt.out.ctx, ctx)
		})
	}
}
