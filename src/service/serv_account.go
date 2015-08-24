package service

import (
	"fmt"
	"net/http"
	"time"
)

func accountHandle(rw http.ResponseWriter, req *http.Request) {
	fmt.Printf("%v account handle\n", time.Now())
	rw.Write([]byte("account"))
}
