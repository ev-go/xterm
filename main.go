package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	requestUseUrl := "http://localhost"
	requestUsePort := ":3000"
	requestUseRout := "/get-token"
	requestUseLogin := "login=root1"
	requestUsePassword := "password=1"
	requestUseData := "data=21"

	httpRequestString := requestUseUrl + requestUsePort + requestUseRout + "?" + requestUseLogin + "&" + requestUsePassword + "&" + requestUseData

	bearer := "Bearer " + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBVFRFTlRJT04hIjoi0J_RgNC40LLQtdGCLCDQnNCw0LrRgSA6KSIsIkRhdGEgYW5zd2VyIGlzIjoiMjExIiwiVG9rZW4gcmVxdWVzdCBhdCI6IjIwMjItMDUtMjVUMjM6NDg6MzYuODAwNDU4MiswNTowMCIsImFkbWluIHBlcm1pc3Npb25zPyI6Im1heWJlIiwiZXhwIjoxNjUzNTY5MzE3LCJsb2dpbiI6InJvb3QxIn0.C6FekKeToH0j-G8GyiMegaoLtWODi9rOK-OM7ModS5Y"
	cli := http.Client{Timeout: 5 * time.Second}
	request, err := http.NewRequest("GET", httpRequestString, nil)
	request.Header.Add("Authorization", bearer)
	request.Header.Add("Content-Type", `application/json`)
	if err != nil {
		panic(err)
	}
	// defer request.Body.Close() /where to add this?

	response, err := cli.Do(request)
	if err != nil {
		panic(err)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("\nresponse?:", string(responseData))

}
