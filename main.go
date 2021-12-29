package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

const URL = "https://script.google.com/macros/s/AKfycbwcV24P8K6peCazJWljy97uIT9LpH3viVqvOCFT9OClP8N-eklHmwvT_Bk60tzwU0EX/exec"

type resData struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "invalid argument")
		return
	}
	v := url.Values{}
	v.Add("text", os.Args[1])
	v.Add("source", "ja")
	v.Add("target", "en")
	resp, err := http.PostForm(URL, v)
	if err != nil {
		log.Fatal(err)
	}
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var d resData
	if err := json.Unmarshal(buf, &d); err != nil {
		log.Fatal(err)
	}
	if d.Code != 200 {
		log.Fatal("")
	}
	fmt.Println(d.Text)
}
