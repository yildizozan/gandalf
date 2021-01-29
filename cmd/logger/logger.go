package logger

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
	"log"
	"time"
)

//var strPost = []byte("POST")
//var strRequestURI = []byte("http://localhost:3100/loki/api/v1/push")

func Warning(format string, v ...interface{}) {
	log.Printf(format, v)
}

func Error(format string, v ...interface{}) {
	log.Fatalf(format, v)
}

func Log(format string, v ...interface{}) {
	log.Printf(format, v)
}

func CreateLog(ip string, method string, path string, scheme string, code string) string {
	return fmt.Sprintf("%s %s %s %s %s %s", ip, time.Now(), method, path, scheme, code)
}

func Logger(service string, malware bool, malwareType string, log string) {
	entry := StreamEntry{
		Stream: Stream{
			App:     service,
			Malware: fmt.Sprintf("%t", malware),
			Type:    malwareType,
		},
		Values: Values{
			{fmt.Sprintf("%d", time.Now().UnixNano()), log},
		},
	}
	logLoki := Loki{Streams: []StreamEntry{entry}}

	// Convert from json to string
	prettyJSON, _ := json.MarshalIndent(logLoki, "", "  ")
	//fmt.Println(string(prettyJSON))

	// Send request to Loki api
	req := fasthttp.AcquireRequest()
	req.Header.SetContentType("application/json")
	req.SetBody([]byte(fmt.Sprintf("%s\n", prettyJSON)))
	req.Header.SetMethodBytes([]byte("POST"))
	req.SetRequestURIBytes([]byte(viper.GetString("app.logger.loki")))
	res := fasthttp.AcquireResponse()
	if err := fasthttp.Do(req, res); err != nil {
		Log("Log not send!\n")
	}
	fasthttp.ReleaseRequest(req)

	//fmt.Println(res)

	fasthttp.ReleaseResponse(res) // Only when you are done with body!
}
