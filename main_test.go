package main

import "testing"

func Test_queryPath(t *testing.T) {
	type args struct {
		url   string
		xpath string
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
		wantErr    bool
	}{
		{
			name:       "query_positive_1",
			args:       args{url: "https://example.com/", xpath: "//div//h1"},
			wantResult: "Example Domain",
			wantErr:    false,
		},
		{
			name:       "query_negative_1",
			args:       args{url: "https://example.com/", xpath: "//div//h2"},
			wantResult: "",
			wantErr:    true,
		},
		{
			name:       "query_negative_2",
			args:       args{url: "https://example.co/", xpath: "//div//h1"},
			wantResult: "",
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := queryPath(tt.args.url, tt.args.xpath)
			if (err != nil) != tt.wantErr {
				t.Errorf("queryPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResult != tt.wantResult {
				t.Errorf("queryPath() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
