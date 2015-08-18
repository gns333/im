package service

import (
	"fmt"
	"net/http"
	"time"
)

func friendHandle(rw http.ResponseWriter, req *http.Request) {
	fmt.Printf("%v friend handle\n", time.Now())
	rw.Write([]byte("friend"))
}
