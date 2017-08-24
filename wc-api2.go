package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type API2STRUCT struct {
	PWD       string
	APP       string
	APPWRITE  string
	Title1    string
	Title2    string
	Text1     string
	Text2     string
	Tags1     string
	Num1      int
	Num2      int
	Num3      int
	Needlogin bool
	//ADDATIONAL
	String1 string
	String2 string
	String3 string
	Int1    int
	Int2    int
	Int3    int
}

func ApiHandler2(w http.ResponseWriter, r *http.Request) { //THIS ONE IS WORKING WITH JSON Requests
	startAPI2 := time.Now()

	APILogin := false

	decoder := json.NewDecoder(r.Body)
	var json API2STRUCT
	errSearch := decoder.Decode(&json)
	if errSearch != nil {
		fmt.Fprintf(w, "ERROR")
		fmt.Println(errSearch)
		checkErr(err)
		return
	}

	if personalpwd == json.PWD {
		APILogin = true
	}

	switch json.APP {
	case "ItemWrite":
		ItemWrite(w, json)
	case "PwdManager":
		PwdManager(w, json)
	default:
		fmt.Fprintf(w, "ERROR")
	}

	fmt.Println("Api2Handler:", time.Since(startAPI2), APILogin)
}

func ItemWrite(w http.ResponseWriter, json API2STRUCT) {
	db.Exec("INSERT INTO items(APP,title1,title2,text1,text2,tags1,num1,num2,num3,needlogin) VALUES(?,?,?,?,?,?,?,?,?,?)", json.APPWRITE, json.Title1, json.Title2, json.Text1, json.Text2, json.Tags1, json.Num1, json.Num2, json.Num3, json.Needlogin)

	foo := "Write " + json.Title1
	Eventloger(foo, json.APPWRITE, 0)

	refreshCache()
	fmt.Fprintf(w, `{"Status":"OK"}`)
	return
}

func PwdManager(w http.ResponseWriter, json API2STRUCT) {

	var ids *sql.Rows

	ids, err = db.Query("SELECT title1,title2,text1 FROM items WHERE APP='PwdManager'")
	defer ids.Close()

	var jsonoutputtmp []string

	checkErr(err)
	for ids.Next() {
		var title1 string
		var title2 string
		var text1 string
		_ = ids.Scan(&title1, &title2, &text1)
		checkErr(err)

		jsonoutputtmp2 := `{ "title1" :"` + title1 + `" , "title2" :" ` + title2 + `","text1":"` + text1 + `" }`

		jsonoutputtmp = append(jsonoutputtmp, jsonoutputtmp2)
	}

	jsonoutput := `{ "PwdResult" : [` + strings.Join(jsonoutputtmp, ",") + "]}"

	fmt.Fprintf(w, jsonoutput)

	fmt.Println("PwdMa")
	return
}
