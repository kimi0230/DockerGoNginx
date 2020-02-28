package redisservices

import (
	"fmt"
	"testing"
)

func TestPing(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "test ping pong4",
			want: "PONG",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ping(); got != tt.want {
				t.Errorf("Ping() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_demoHash(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "get my name",
			want: "kimi",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := demoHash(); got != tt.want {
				fmt.Println("got1 ", got)
				t.Errorf("demoHash() = %v, want %v", got, tt.want)
			}
			/* 測試是否會過期
			time.Sleep(time.Duration(2) * time.Second)
			rc := RedisClient.Get()
			defer rc.Close()
			if res, _ := rc.Do("hget", "me", "name"); res != nil {
				t.Errorf("demoHash() = %v, want %v", res, "")
			}
			*/
		})
	}
}
