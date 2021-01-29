package detector

import (
	"github.com/yildizozan/gandalf/cmd/config/v2"
	"testing"
)

func TestAnalyseIp(t *testing.T) {

	whitelist := []string{
		"127.0.0.1",
		"10.0.0.1",
	}
	blacklist := []string{
		"192.168.1.0",
		"192.168.1.1",
		"192.168.1.2",
		"192.168.1.3",
	}

	rules := v2.Ip{
		Whitelist: whitelist,
		Blacklist: blacklist,
	}

	validIp := "127.0.0.1"
	unvalidIp := "192.168.1.0"

	channel := make(chan bool)

	type args struct {
		rules      *v2.Ip
		remoteAddr *string
		channel    chan bool
	}
	tests := []struct {
		name string
		args args
	}{
		{"Valid IP", args{&rules, &validIp, channel}},
	}
	for _, tt := range tests {
		go analyseIp(tt.args.rules, tt.args.remoteAddr, tt.args.channel)
		t.Run(tt.name, func(t *testing.T) {
			result := <-channel
			if result {
				t.Error(tt.name, tt.args.remoteAddr)
			}
		})
	}

	tests = []struct {
		name string
		args args
	}{
		{"NotValid IP", args{&rules, &unvalidIp, channel}},
	}
	for _, tt := range tests {
		go analyseIp(tt.args.rules, tt.args.remoteAddr, tt.args.channel)
		t.Run(tt.name, func(t *testing.T) {
			result := <-channel
			if !result {
				t.Error(tt.name, tt.args.remoteAddr)
			}
		})
	}
}
