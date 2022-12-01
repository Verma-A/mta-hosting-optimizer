package mocksetup

import (
	"reflect"
	"testing"
)

func TestGetData(t *testing.T) {
	type args struct {
		value int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Positive",
			args: args{
				value: 1,
			},
			want: []string{
				"test host",
				"test host1",
			},
		},
	}
	for _, tt := range tests {
		AddData([]DataModel{
			{"1.1.1.1", "test host", true},
			{"1.1.1.1", "test host1", false},
			{"1.1.1.1", "test host2", true},
			{"1.1.1.1", "test host2", true},
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := GetData(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetData() = %v, want %v", got, tt.want)
			}
		})
	}
}
