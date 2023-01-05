package logger

import (
	"reflect"
	"testing"

	"go.uber.org/zap/zapcore"
)

func Test_getLevel(t *testing.T) {
	type args struct {
		lvl string
	}
	tests := []struct {
		name string
		args args
		want zapcore.Level
	}{
		{
			name: "space",
			args: args{""},
			want: zapcore.InfoLevel,
		},
		{
			name: "error",
			args: args{"error"},
			want: zapcore.ErrorLevel,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLevel(tt.args.lvl); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}
