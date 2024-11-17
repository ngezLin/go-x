package redis

import (
	"context"
	"testing"
	"time"

	"github.com/go-redis/redismock/v9"
	"github.com/stretchr/testify/assert"
	"github.com/super-saga/go-x/log"
	cacheMock "github.com/super-saga/go-x/redis/mock"
	"go.uber.org/mock/gomock"
)

func redisTestHelper(t *testing.T) *cacheMock.MockRedis {
	t.Helper()
	t.Parallel()
	log.InitForTest()
	mockGen := gomock.NewController(t)
	newMockRedis := cacheMock.NewMockRedis(mockGen)

	return newMockRedis
}

func TestGet(t *testing.T) {
	ctx := context.Background()

	rc := redisTestHelper(t)

	type args struct {
		key string
	}

	tests := []struct {
		name           string
		args           args
		doMock         func(args args)
		want           string
		wantErr        bool
		expectedOutput string
		expectedError  error
	}{
		{
			name: "Success",
			args: args{
				key: "access_token",
			},
			want:    "granted",
			wantErr: false,
			doMock: func(args args) {
				rc.EXPECT().Get(ctx, args.key).Return("granted", nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.doMock != nil {
				tt.doMock(tt.args)
			}

			got, err := rc.Get(ctx, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Redis.Get() error = %v, wantErr %v", err, tt.wantErr)
			}

			if got != tt.want {
				t.Errorf("Redis.Get() error = %v, want %v", got, tt.want)
			}

			db, mock := redismock.NewClientMock()
			s, _ := New(db)
			output, err := s.Get(ctx, tt.args.key)
			if (err != nil) == tt.wantErr {
				assert.Equal(t, tt.expectedError, err)
			}
			assert.Equal(t, tt.expectedOutput, output)
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Error(err)
			}
			mock.ClearExpect()
		})
	}
}

func TestSet(t *testing.T) {
	ctx := context.Background()

	rc := redisTestHelper(t)

	type args struct {
		key   string
		value interface{}
		ttl   time.Duration
	}

	tests := []struct {
		name          string
		args          args
		doMock        func(args args)
		want          string
		wantErr       bool
		expectedError error
	}{
		{
			name: "Success",
			args: args{
				key:   "access_token",
				value: "success",
				ttl:   30 * time.Second,
			},
			wantErr: false,
			doMock: func(args args) {
				rc.EXPECT().Set(ctx, args.key, args.value, args.ttl).Return(nil).AnyTimes()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.doMock != nil {
				tt.doMock(tt.args)
			}

			db, mock := redismock.NewClientMock()
			s, _ := New(db)
			err := s.Set(ctx, tt.args.key, tt.args.value, tt.args.ttl)
			if (err != nil) == tt.wantErr {
				assert.Equal(t, tt.expectedError, err)
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Error(err)
			}
			mock.ClearExpect()
		})
	}
}
