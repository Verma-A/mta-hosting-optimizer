package mocksetup

import (
	"testing"
)

func TestInitData(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Positive",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitData()
		})
	}
}

func TestAddData(t *testing.T) {
	type args struct {
		data []DataModel
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Positive",
			args: args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddData(tt.args.data)
		})
	}
}
