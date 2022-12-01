package main

import (
	"errors"
	mocksetup "mta-hosting-optimizer/mockSetup"
	"net/http"
	"os"
	"testing"
)

type testResWriter struct{}

func (t *testResWriter) Header() http.Header {
	return http.Header{}
}

func (t *testResWriter) Write(data []byte) (int, error) {
	if string(data) == "1" {
		return 0, errors.New("test-error")
	}
	return 0, nil
}

func (t *testResWriter) WriteHeader(statusCode int) {}

func TestHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
		mock func() *http.Request
	}{
		{
			name: "Positive, get method",
			args: args{
				w: &testResWriter{},
			},
			mock: func() *http.Request {
				r, _ := http.NewRequest("GET", "something.com", nil)
				os.Setenv("tempKey", "1")
				return r
			},
		},
		{
			name: "Negative, get method, error from write function",
			args: args{
				w: &testResWriter{},
			},
			mock: func() *http.Request {
				r, _ := http.NewRequest("GET", "something.com", nil)
				os.Setenv("tempKey", "1")
				mocksetup.AddData([]mocksetup.DataModel{
					{IP: "1", HostName: "1", Active: false},
				})
				return r
			},
		},
		{
			name: "Negative, get method, error from env variable",
			args: args{
				w: &testResWriter{},
			},
			mock: func() *http.Request {
				r, _ := http.NewRequest("GET", "something.com", nil)
				os.Setenv("tempKey", "")
				return r
			},
		},
		{
			name: "Negative, wrong http method",
			args: args{
				w: &testResWriter{},
			},
			mock: func() *http.Request {
				r, _ := http.NewRequest("POST", "something.com", nil)
				os.Setenv("tempKey", "")
				return r
			},
		},
	}
	for _, tt := range tests {
		tt.args.r = tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			Handler(tt.args.w, tt.args.r)
		})
	}
}
