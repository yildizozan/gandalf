package detector

import (
	"testing"
)

func TestAnalyseRawQuery(t *testing.T) {
	channel := make(chan bool)
	type args struct {
		uri     string
		channel chan bool
	}

	// Valid
	tests := []struct {
		name string
		args args
	}{
		{"unvalid", args{
			"/api/healtz?username=yildizozan",
			channel,
		}},
	}

	for _, tt := range tests {
		go analyseRawQuery(tt.args.uri, tt.args.channel)
		t.Run(tt.name, func(t *testing.T) {
			result := <-channel
			if result {
				t.Error(tt.name, tt.args.uri)
			}
		})

	}

	// Not valid
	tests = []struct {
		name string
		args args
	}{
		{"unvalid", args{
			"username=yildizozan=12325123123&email=%27SELECT+%2A+FROM+users",
			channel,
		}},
	}

	for _, tt := range tests {
		go analyseRawQuery(tt.args.uri, tt.args.channel)
		t.Run(tt.name, func(t *testing.T) {
			result := <-channel
			if !result {
				t.Error(tt.name, tt.args.uri)
			}
		})

	}

}
