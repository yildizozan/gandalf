package detector

import "testing"

func BenchmarkAnalyseRawQueryBestCase(b *testing.B) {
	for n := 0; n < b.N; n++ {
		AnalyseRawQuery("username=yildizozan%27+1%3D1&email=%27SELECT+%2A+FROM+users")
	}
}

func BenchmarkAnalyseRawQueryWorseCase(b *testing.B) {
	for n := 0; n < b.N; n++ {
		AnalyseRawQuery("username=yildizozan=12325123123&email=%27SELECT+%2A+FROM+users")
	}
}
