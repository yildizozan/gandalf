package detector

import (
	"github.com/yildizozan/gandalf/cmd/config"
	"testing"
)

var rules = config.Path{
	Prefix: "/admin",
	Exact:  "/admin",
	Match:  "/admin",
}

func BenchmarkAnalyseRawQueryBestCase(b *testing.B) {
	path := "username=yildizozan%27+1%3D1&email=%27SELECT+%2A+FROM+users"
	c := make(chan bool)
	for n := 0; n < b.N; n++ {
		analysePath(&rules, &path, c)
		<-c
	}
}

func BenchmarkAnalyseRawQueryWorseCase(b *testing.B) {
	path := "username=yildizozan=12325123123&email=%27SELECT+%2A+FROM+users"
	c := make(chan bool)
	for n := 0; n < b.N; n++ {
		analysePath(&rules, &path, c)
		<-c
	}
}
