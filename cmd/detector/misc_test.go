package detector

import "testing"

func TestRegex(t *testing.T) {
	const model = `(?i)(\%27)|(\')|(\-\-)|(\%23)|(\#)`

	type args struct {
		str   string
		regex string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"valid", args{
			str:   "/api/healtz?username=yildizozan",
			regex: model,
		}, false},
		{"not valid", args{
			str:   "username=yildizozan%27+1%3D1&email=%27SELECT+%2A+FROM+users",
			regex: model,
		}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := regex(tt.args.str, tt.args.regex); got != tt.want {
				t.Errorf("regex() = %v, want %v", got, tt.want)
			}
		})
	}
}
