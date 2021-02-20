package client

import (
	"crypto/tls"
	"go-api-automated-testing/common"
	"io"
	"net/http"
	"time"
)

func HTTPrequest(mothod string, url string, body io.Reader, timeout time.Duration) (*http.Response, uint64, error) {

	// 跳过证书验证
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Transport: tr,
		Timeout:   timeout,
	}

	reqest, err := http.NewRequest(mothod, url, body)
	startTime := time.Now()
	resp, err := client.Do(reqest)
	requestTime := uint64(common.DiffNano(startTime))

	return resp, requestTime, err

}
