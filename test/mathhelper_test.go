package test

import (
	mathhelper "DockerGoNginx/app/helper/mathHelper"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Simple testing what different between Fatal and Error
func TestRound(t *testing.T) {
	c := mathhelper.Round(73.75, 1)
	if c == 0 {
		t.Error("number should not be null")
	}

	assert.Exactly(t, 73.8, c)
}

func TestDecimal(t *testing.T) {
	type args struct {
		value float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "no1",
			args: args{
				value: 10.75,
			},
			want: 10.8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mathhelper.Decimal(tt.args.value); got != tt.want {
				t.Errorf("Decimal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecimal2(t *testing.T) {
	type args struct {
		value float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "no1",
			args: args{
				value: 10.75,
			},
			want: 10.75,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mathhelper.Decimal2(tt.args.value); got != tt.want {
				t.Errorf("Decimal2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecimal2Parallel(t *testing.T) {
	type args struct {
		value float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "no1",
			args: args{
				value: 10.755,
			},
			want: 10.76,
		},
		{
			name: "no2",
			args: args{
				value: 10.753,
			},
			want: 10.75,
		},
	}
	for _, tt := range tests {
		tt := tt // 解決 Go 只最後一個情境
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := mathhelper.Decimal2(tt.args.value); got != tt.want {
				t.Errorf("Decimal2() = %v, want %v", got, tt.want)
			}
		})
	}
}
