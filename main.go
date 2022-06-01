package main

import (
	"bufio"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
)

// func isP(s string) string {
// 	mid := len(s) / 2
// 	last := len(s) - 1
// 	for i := 0; i < mid; i++ {
// 		if s[i] != s[last-i] {
// 			return "NO. It's not a Palimdrome."
// 		}
// 	}
// 	return "YES! You've entered a Palindrome"
// }
// func change(x, y string) (string, string) {
// 	return y, x
// }

// type httpRequestMessageStruct struct {
// 	requestUseLogin
// 	requestUsePassword
// 	requestUseData
// }

// type httpRequestStruct struct {
// 	requestUseUrl      string
// 	requestUsePort     string
// 	requestUseRout     string
// 	httpRequestMessage httpRequestMessageStruct
// }
var ctx = context.Background()

func main() {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	requestUseUrl := "localhost"
	requestUsePort := "3000"
	requestUseRout := "get-token"
	requestUseLogin := "root1"
	requestUsePassword := "1"
	requestUseData := "21"

	httpRequestString := "http://" + requestUseUrl + ":" + requestUsePort + "/" + requestUseRout + "?login=" + requestUseLogin + "&password=" + requestUsePassword + "&data=" + requestUseData
	// Menu
	fmt.Println("\n********************************/ Menu /****************************************")
	fmt.Println("\nThis is client for sending http requests to server")
	fmt.Println("\nDefault URL: ", requestUseUrl, ";",
		"\nDefault Port: ", requestUsePort, ";",
		"\nDefault Rout: ", requestUseRout, ";",
		"\nDefault Login: ", requestUseLogin, ";",
		"\nDefault Password: ", requestUsePassword, ";",
		"\nDefault Data: ", requestUseData, ";",
		"\nDefault http request: ", httpRequestString, ";")

	// fmt.Println("Логин")
	// fmt.Scanf("%s\n", &Log)

	// fmt.Println("Пароль")
	// fmt.Scanf("%s\n", &Pass)
	// terminalInputSlice := []string{"Param", "value", "value1", "value2"}
	// fmt.Println(terminalInputSlice[0])
	// fmt.Println(terminalInputSlice[1])
	// fmt.Println(terminalInputSlice[2])
	readFromTerminal := bufio.NewScanner(os.Stdin)
	// use `for scanner.Scan()` to keep reading
	// line := readFromTerminal.Text()
	// var terminalParameter string
	// var terminalValue string
	helpChangeDefaults := "\nTo change defaults enter: <value to change> <new value> \nFor example: login user8 \nOr password 12345"
	fmt.Println("\nDo you want to change defaults? (y/n)")
	// fmt.Scanf("%s\n", &readFromTerminal)
	readFromTerminal.Scan()
	if readFromTerminal.Text() == "y" {
		fmt.Println(helpChangeDefaults)
		readFromTerminal.Scan()
		line := readFromTerminal.Text()
		terminalInputSlice := strings.Fields(line)
		fmt.Println("\n<Value to change:>", terminalInputSlice[0], "<New value:>", terminalInputSlice[1])
		node := rdb.Set(ctx, terminalInputSlice[0], terminalInputSlice[1], 0).Err()
		if node != nil {
			panic(node)
		}
		// switch terminalInputSlice[0] {
		// case "url":
		// 	requestUseUrl = terminalInputSlice[1]
		// case "port":
		// 	requestUsePort = terminalInputSlice[1]
		// case "rout":
		// 	requestUseRout = terminalInputSlice[1]
		// case "login":
		// 	requestUseLogin = terminalInputSlice[1]
		// case "password":
		// 	requestUsePassword = terminalInputSlice[1]
		// case "data":
		// 	requestUseData = terminalInputSlice[1]
		// }
		// terminalInputSlice := strings.Fields(helpChangeDefaults)
		// fmt.Println(terminalInputSlice[2])
		// stringfromterm := readFromTerminal
		// terminalInputSlice := strings.Fields(stringfromterm)
		// fmt.Println(terminalInputSlice[2])
		// fmt.Println(terminalInputSlice[0])
		// fmt.Println(terminalInputSlice[1])
		// fmt.Println(terminalInputSlice[2])
	} else {
		fmt.Println("No changes")
	}
	// if readFromTerminal == "y" {
	// 	fmt.Println("\nWhat part of http request need to change?",
	// 		"Write in terminal: 'url' or 'port' or 'rout' or 'login' or 'password' or 'data'")
	// 	fmt.Scanf("%s\n", &readFromTerminal)
	// 	if readFromTerminal == "login" {
	// 		fmt.Println("\nEnter new login")
	// 		fmt.Scanf("%s\n", &readFromTerminal)
	// 		requestUseLogin = "login=" + readFromTerminal
	// 		fmt.Println("\nLogin changed for:", requestUseLogin)
	// 	}
	// 	if readFromTerminal == "password" {
	// 		fmt.Println("\nEnter new password")
	// 		fmt.Scanf("%s\n", &readFromTerminal)
	// 		requestUsePassword = "password=" + readFromTerminal
	// 		fmt.Println("\nPassword changed for:", requestUsePassword)
	// 	}
	// 	if readFromTerminal == "data" {
	// 		fmt.Println("\nEnter new data")
	// 		fmt.Scanf("%s\n", &readFromTerminal)
	// 		requestUseData = "data=" + readFromTerminal
	// 		fmt.Println("\nData changed for:", requestUseData)
	// 	}
	// 	if readFromTerminal == "port" {
	// 		fmt.Println("\nEnter new port")
	// 		fmt.Scanf("%s\n", &readFromTerminal)
	// 		requestUsePort = ":" + readFromTerminal
	// 		fmt.Println("\nPort changed for:", requestUsePort)
	// 	}
	// } else {
	// 	fmt.Println("No changes")
	// }

	requestUseUrl, node := rdb.Get(ctx, "url").Result()
	if node == redis.Nil {
		fmt.Println("url does not exist")
		fmt.Println(helpChangeDefaults)
		readFromTerminal.Scan()
		line := readFromTerminal.Text()
		terminalInputSlice := strings.Fields(line)
		fmt.Println("\n<Value to change:>", terminalInputSlice[0], "<New value:>", terminalInputSlice[1])
		node := rdb.Set(ctx, terminalInputSlice[0], terminalInputSlice[1], 0).Err()
		if node != nil {
			panic(node)
		}
	} else if node != nil {
		panic(node)
	} else {
		fmt.Println("url:", requestUseUrl)
	}

	requestUsePort, node = rdb.Get(ctx, "port").Result()
	if node == redis.Nil {
		fmt.Println("port does not exist")
		fmt.Println(helpChangeDefaults)
		readFromTerminal.Scan()
		line := readFromTerminal.Text()
		terminalInputSlice := strings.Fields(line)
		fmt.Println("\n<Value to change:>", terminalInputSlice[0], "<New value:>", terminalInputSlice[1])
		node := rdb.Set(ctx, terminalInputSlice[0], terminalInputSlice[1], 0).Err()
		if node != nil {
			panic(node)
		}
	} else if node != nil {
		panic(node)
	} else {
		fmt.Println("port:", requestUsePort)
	}

	requestUseRout, node = rdb.Get(ctx, "rout").Result()
	if node == redis.Nil {
		fmt.Println("rout does not exist")
		fmt.Println(helpChangeDefaults)
		readFromTerminal.Scan()
		line := readFromTerminal.Text()
		terminalInputSlice := strings.Fields(line)
		fmt.Println("\n<Value to change:>", terminalInputSlice[0], "<New value:>", terminalInputSlice[1])
		node := rdb.Set(ctx, terminalInputSlice[0], terminalInputSlice[1], 0).Err()
		if node != nil {
			panic(node)
		}
	} else if node != nil {
		panic(node)
	} else {
		fmt.Println("rout:", requestUseRout)
	}

	requestUseLogin, node = rdb.Get(ctx, "login").Result()
	if node == redis.Nil {
		fmt.Println("login does not exist")
		fmt.Println(helpChangeDefaults)
		readFromTerminal.Scan()
		line := readFromTerminal.Text()
		terminalInputSlice := strings.Fields(line)
		fmt.Println("\n<Value to change:>", terminalInputSlice[0], "<New value:>", terminalInputSlice[1])
		node := rdb.Set(ctx, terminalInputSlice[0], terminalInputSlice[1], 0).Err()
		if node != nil {
			panic(node)
		}
	} else if node != nil {
		panic(node)
	} else {
		fmt.Println("login:", requestUseLogin)
	}

	requestUsePassword, node = rdb.Get(ctx, "password").Result()
	if node == redis.Nil {
		fmt.Println("password does not exist")
		fmt.Println(helpChangeDefaults)
		readFromTerminal.Scan()
		line := readFromTerminal.Text()
		terminalInputSlice := strings.Fields(line)
		fmt.Println("\n<Value to change:>", terminalInputSlice[0], "<New value:>", terminalInputSlice[1])
		node := rdb.Set(ctx, terminalInputSlice[0], terminalInputSlice[1], 0).Err()
		if node != nil {
			panic(node)
		}
	} else if node != nil {
		panic(node)
	} else {
		fmt.Println("password:", requestUsePassword)
	}

	requestUseData, node = rdb.Get(ctx, "data").Result()
	if node == redis.Nil {
		fmt.Println("data does not exist")
		fmt.Println(helpChangeDefaults)
		readFromTerminal.Scan()
		line := readFromTerminal.Text()
		terminalInputSlice := strings.Fields(line)
		fmt.Println("\n<Value to change:>", terminalInputSlice[0], "<New value:>", terminalInputSlice[1])
		node := rdb.Set(ctx, terminalInputSlice[0], terminalInputSlice[1], 0).Err()
		if node != nil {
			panic(node)
		}
	} else if node != nil {
		panic(node)
	} else {
		fmt.Println("data:", requestUseData)
	}

	httpRequestString = "http://" + requestUseUrl + ":" + requestUsePort + "/" + requestUseRout + "?login=" + requestUseLogin + "&password=" + requestUsePassword + "&data=" + requestUseData

	fmt.Println("\n******************************/ Menu End /**************************************")
	//Menu end

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
