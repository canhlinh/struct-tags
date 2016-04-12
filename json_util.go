package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func GenerateJsonResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	b, _ := json.Marshal(data)
	var prettyJSON bytes.Buffer
	json.Indent(&prettyJSON, b, "", "  ")
	w.Write(prettyJSON.Bytes())
}

func ReadFileConfig() JsonObject {
	file, e := ioutil.ReadFile("./config.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	var jsonObject JsonObject
	json.Unmarshal(file, &jsonObject)
	return jsonObject
}
