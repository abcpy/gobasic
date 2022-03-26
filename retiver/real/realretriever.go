package real

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"time"
)

func (r *Retriever) String() string {
	return fmt.Sprintf("xxx")
}

type Retriever struct {
	UserAgent string
	Timeout   time.Duration
}

func (r *Retriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	result, err := httputil.DumpResponse(resp, true)

	resp.Body.Close()

	if err != nil {
		panic(err)
	}
	return string(result)
}
