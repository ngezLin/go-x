package ctxdata

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/metadata"
)

func Test_Sets(t *testing.T) {
	tests := []struct {
		name     string
		ctx      context.Context
		setters  []Set
		expected context.Context
	}{
		{
			name:     "Given empty context, when setting it with empty setters, then should return empty context",
			ctx:      context.Background(),
			setters:  []Set{},
			expected: context.Background(),
		},
		{
			name:     "Given empty context, when setting it with pid setter, then should return context with pid",
			ctx:      context.Background(),
			setters:  []Set{SetPid("test")},
			expected: context.WithValue(context.Background(), pidKey{}, "test"),
		},
		// TODO: Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Sets(tt.ctx, tt.setters...)

			assert.Equal(t, tt.expected, result)
		})
	}
}

func Test_SetsMD(t *testing.T) {
	tests := []struct {
		name     string
		md       metadata.MD
		setters  []SetMD
		expected metadata.MD
	}{
		{
			name:     "Given empty metadata, when setting it with empty setters, then should return empty metadata",
			md:       metadata.New(nil),
			setters:  []SetMD{},
			expected: metadata.New(nil),
		},
		{
			name:     "Given empty metadata, when setting it with pid, then should return metadata with pid",
			md:       metadata.New(nil),
			setters:  []SetMD{SetMDPid("test")},
			expected: map[string][]string{pidMDKey: []string{"test"}},
		},
		// TODO: Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SetsMD(tt.md, tt.setters...)

			assert.Equal(t, tt.expected, result)
		})
	}
}
