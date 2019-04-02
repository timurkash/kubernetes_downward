package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	yaml2 "gopkg.in/yaml.v2"
)

const (
	PORT = "8002"
	KUBERNETES_PROXY = "http://localhost:8001"
)

func main() {
	http.HandleFunc("/", postHandler)
	log.Printf("Listening on port %s ...", PORT)
	http.ListenAndServe(":" + PORT, nil)
}
func postHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	log.Println(r.URL.Path)
	url := KUBERNETES_PROXY + r.URL.Path
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	w.Write(json2Yaml(body))

}

func json2Yaml(body []byte)[]byte {
	m := make(map[string]interface{})
	err := json.Unmarshal(body, &m)
	if err != nil {
		panic(err)
	}
	d, err2 := yaml2.Marshal(&m)
	if err2 != nil {
		panic(err2)
	}
	return d
}