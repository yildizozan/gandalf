package detector

import (
	"fmt"
	"net/http"
)

func Analyse(req *http.Request) bool {

	chanQuery := make(chan bool)
	chanHeader := make(chan bool)
	chanIp := make(chan bool)
	chanPath := make(chan bool)

	go analyseRawQuery(req.URL.RawQuery, chanQuery)
	go analyseHeaders(req.Header, chanHeader)
	go analyseIp(&req.RemoteAddr, chanIp)
	go analysePath(req.URL.Path, chanPath)

	query, header, ip, path := <-chanQuery, <-chanHeader, <-chanIp, <-chanPath
	fmt.Println("query\t", query)
	fmt.Println("header\t", header)
	fmt.Println("ip\t", ip)
	fmt.Println("path\t", path)

	return query || header || ip || path
}
