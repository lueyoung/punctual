package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	//"sync"
	"encoding/json"
	"strings"

	//"time"

	//"k8s.io/apimachinery/pkg/api/errors"
	//metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	//"k8s.io/client-go/kubernetes"
	//rest0 "k8s.io/client-go/rest"
	//v1beta1 "k8s.io/api/extensions/v1beta1"
	"dbuntil"
	"github.com/ant0ine/go-json-rest/rest"
	incluster "inClusterServiceDiscovery"
	table "tableFormat"
)

func main() {
	gettest()
}

func gettest() {
	key := "getModelDef"
	s, err := dbutil.CreateSession(key)
	if err != nil {
		msg := "err: create session"
		log.Println("msg")
		w.WriteJson(map[string]string{"Err": msg})
		return
	}
	fmt.Printf("Key: %v\n", key)
	raw, err := s.Get()
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	str := strings.Trim(raw, key+",")
	fmt.Printf("Json str: %v\n", str)
	var ret jsonresp
	json.Unmarshal([]byte(str), &ret)
	fmt.Printf("ret: %v\n", ret)
}

type field struct {
	DataType string `json:"DataType"`
	Desc     string `json:"Desc"`
	ID       string `json:"ID"`
	Name     string `json:"Name"`
	Size     string `json:"Size"`
	Flag     string `json:"flag"`
	Index    string `json:"index"`
	Keygroup string `json:"keygroup"`
	Shift    string `json:"shift"`
}

type tablex struct {
	Alias    string  `json:"Alias"`
	ID       string  `json:"ID"`
	Name     string  `json:"Name"`
	PriKey   string  `json:"PriKey"`
	Fieldnum string  `json:"fieldnum"`
	Fields   []field `json:"fields"`
	Flag     string  `json:"flag"`
	Index    string  `json:"index"`
	Path     string  `json:"path"`
}

type db struct {
	Alias string   `json:"Alias"`
	ID    string   `json:"ID"`
	Name  string   `json:"Name"`
	Table []tablex `json:"table"`
	Type  string   `json:"type"`
}

type setting struct {
	Database []db `json:"database"`
}

type result struct {
	Setting setting `json:"Setting"`
}

type replyCode struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type jsonresp struct {
	ReplyCode replyCode `json:"replyCode"`
	Result    result    `json:"result"`
}
