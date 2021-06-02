package utils

import (
	"reflect"
	"testing"
	"time"
)

func TestTimeToDate(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{name: "1", args: args{t: time.Now()}, want: func() time.Time { t, _ := time.ParseInLocation("20060102", time.Now().Format("20060102"), time.Local); return t }()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimeToDateTime(tt.args.t); !reflect.DeepEqual(got.Unix(), tt.want.Unix()) {
				t.Errorf("TimeToDateTime() = %v,%v, want %v", got, got.Unix(), tt.want.Unix())
			}
		})
	}
}
