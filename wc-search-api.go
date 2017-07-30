package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type SearchLoginSTRUCT struct {
	PWD         string
	SearchValue string
}

func ApiSearch(w http.ResponseWriter, r *http.Request) {
	startSearch := time.Now()

	SearchLogin := false

	decoder := json.NewDecoder(r.Body)
	var Sjson SearchLoginSTRUCT
	errSearch := decoder.Decode(&Sjson)
	if errSearch != nil {
		fmt.Fprintf(w, "ERROR")
		checkErr(err)
		return
	}

	if personalpwd == Sjson.PWD {
		SearchLogin = true
	}
	//--------------------------------------------

	fmt.Fprintf(w, strconv.FormatBool(SearchLogin))
	fmt.Println("SearchApi:", time.Since(startSearch), "Login:", SearchLogin)

}
