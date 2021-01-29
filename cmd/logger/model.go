package logger

type Stream struct {
	App     string `json:"app"`
	Malware string `json:"malware"`
	Type    string `json:"type"`
}

type Value [2]string

type Values []Value

type StreamEntry struct {
	Stream Stream `json:"stream"`
	Values Values `json:"values"`
}

type Streams []StreamEntry

type Loki struct {
	Streams Streams `json:"streams"`
}
