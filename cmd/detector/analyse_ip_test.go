package detector

import (
	"github.com/yildizozan/gandalf/cmd/config"
	"testing"
)

func Test_analyseIp(t *testing.T) {

	whitelist := []string{
		"127.0.0.1",
	}
	blacklist := []string{
		"192.168.1.0",
		"192.168.1.1",
		"192.168.1.2",
		"192.168.1.3",
	}

	rules := config.Ip{
		Whitelist: whitelist,
		Blacklist: blacklist,
	}

	validIp := "127.0.0.1"
	unvalidIp := "192.168.1.0"

	channel := make(chan bool)

	type args struct {
		rules      *config.Ip
		remoteAddr *string
		channel    chan bool
	}
	tests := []struct {
		name string
		args args
	}{
		{"Valid ip", args{&rules, &validIp, channel}},
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
		{"Unalid ip", args{&rules, &unvalidIp, channel}},
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
