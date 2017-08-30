package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type API2STRUCT struct {
	PWD       string
	APP       string
	APPWRITE  string
	ID        int
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

type ItemData struct {
	ItemID     string
	Timecreate string
	Title1     string
	Title2     string
	Text1      string
	Text2      string
	Tags1      string
	Num1       int
	Num2       int
	Num3       int
}

func ApiHandler2(w http.ResponseWriter, r *http.Request) { //THIS ONE IS WORKING WITH jsondataRequests
	startAPI2 := time.Now()

	APILogin := false

	decoder := json.NewDecoder(r.Body)
	var jsondata API2STRUCT
	errSearch := decoder.Decode(&jsondata)
	if errSearch != nil {
		fmt.Fprintf(w, `{"Status":"ERROR"}`)
		fmt.Println(errSearch)
		checkErr(err)
		return
	}

	if conf.AdminHASH == jsondata.PWD {
		APILogin = true
	}

	switch jsondata.APP {
	case "ItemWrite":
		ItemWrite(w, jsondata)
	case "ItemDelete":
		ItemDelete(w, jsondata)
	case "PwdManager":
		PwdManager(w, jsondata)
	case "ProjectRead":
		ProjectRead(w, jsondata)
	default:
		fmt.Fprintf(w, `{"Status":"ERROR"}`)
	}

	fmt.Println("Api2Handler:", time.Since(startAPI2), APILogin)
}

func ItemWrite(w http.ResponseWriter, jsondata API2STRUCT) {
	db.Exec("INSERT INTO items(APP,title1,title2,text1,text2,tags1,num1,num2,num3,needlogin) VALUES(?,?,?,?,?,?,?,?,?,?)", jsondata.APPWRITE, jsondata.Title1, jsondata.Title2, jsondata.Text1, jsondata.Text2, jsondata.Tags1, jsondata.Num1, jsondata.Num2, jsondata.Num3, jsondata.Needlogin)

	foo := "Write " + jsondata.Title1
	Eventloger(foo, jsondata.APPWRITE, 0)

	refreshCache()
	fmt.Fprintf(w, `{"Status":"OK"}`)
	return
}

func ItemDelete(w http.ResponseWriter, jsondata API2STRUCT) {
	id := jsondata.ID
	ItemBackuper(id)

	eventname := "DEL >" + strconv.Itoa(id) + "< from ProjectTask"
	Eventloger(eventname, "ProjectTask", id)

	db.Exec("DELETE from items WHERE ItemID = ?", id)

	fmt.Fprintf(w, `{"Status":"OK"}`)
	fmt.Println("ProjectTaskDELETE")
	return
}

func PwdManager(w http.ResponseWriter, jsondata API2STRUCT) {

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

func ProjectRead(w http.ResponseWriter, jsondata API2STRUCT) {

	type ProjectData struct {
		APPreturn string
		DATA      []ItemData
	}

	var ProjectDataCA ProjectData

	var ids *sql.Rows

	ids, err = db.Query("SELECT ItemID,timecreate,title1,title2,text1,text2,tags1,num1,num2,num3 FROM items WHERE APP='ProjectTask' AND num3 = ? ORDER BY num1 DESC", jsondata.ID)
	defer ids.Close()

	checkErr(err)
	for ids.Next() {
		var ItemID, timecreate, title1, title2, text1, text2, tags1 string
		var num1, num2, num3 int

		_ = ids.Scan(&ItemID, &timecreate, &title1, &title2, &text1, &text2, &tags1, &num1, &num2, &num3)
		checkErr(err)

		data := ItemData{
			ItemID:     ItemID,
			Timecreate: timecreate,
			Title1:     title1,
			Title2:     title2,
			Text1:      text1,
			Text2:      text2,
			Tags1:      tags1,
			Num1:       num1,
			Num2:       num2,
			Num3:       num3,
		}

		ProjectDataCA.DATA = append(ProjectDataCA.DATA, data)
	}

	ProjectDataCA.APPreturn = "ProjectRead"

	foo2, _ := json.Marshal(ProjectDataCA)

	fmt.Fprintf(w, string(foo2))

	fmt.Println("ProjectRead ")
	return
}
