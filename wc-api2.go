package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type API2STRUCT struct {
	PWD     string
	APP     string
	string1 string
	string2 string
	int1    int
	itn2    int
}

func ApiHandler2(w http.ResponseWriter, r *http.Request) { //THIS ONE IS WORKING WITH JSON Requests
	startAPI2 := time.Now()

	APILogin := false

	decoder := json.NewDecoder(r.Body)
	var json API2STRUCT
	errSearch := decoder.Decode(&json)
	if errSearch != nil {
		fmt.Fprintf(w, "ERROR")
		checkErr(err)
		return
	}

	if personalpwd == json.PWD {
		APILogin = true
	}

	switch json.APP {
	case "PwdManager":
		PwdManager(json)
	}

	fmt.Println("Api2Handler:", startAPI2, APILogin)
}

func PwdManager(json API2STRUCT) {
	fmt.Println(json)
}
