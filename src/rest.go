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
	table "tableFormat"
)

var port = flag.String("p", "8080", "Port ton listen")

func init() {
	flag.Parse()
}

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/api/snap/:key", get),
		rest.Put("/api/snap/:key", put),
		rest.Get("/api/test", test),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	addr := fmt.Sprintf(":%s", *port)
	log.Fatal(http.ListenAndServe(addr, api.MakeHandler()))
}

type Info struct {
	Key   string
	Value string
}

func test(w rest.ResponseWriter, r *rest.Request) {
	info := Info{}
	key := "hello"
	value := "world"
	info.Key = key
	info.Value = value
	w.WriteJson(info)
}

func get1(w rest.ResponseWriter, r *rest.Request) {
	key := r.PathParam("key")
	s, err := dbutil.CreateSession(key)
	if err != nil {
		msg := "err: create session"
		log.Println("msg")
		w.WriteJson(map[string]string{"Err": msg})
		return
	}
	fmt.Printf("Key: %v\n", key)
	str, err := s.Get()
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t, err := table.FromStringToTable(str)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(*t)
}

func get(w rest.ResponseWriter, r *rest.Request) {
	key := r.PathParam("key")
	if key == "getModelDef" {
		getmodeldef(w, r)
		return
	}
	s, err := dbutil.CreateSession(key)
	if err != nil {
		msg := "err: create session"
		log.Println("msg")
		w.WriteJson(map[string]string{"Err": msg})
		return
	}
	//fmt.Printf("Key: %v\n", key)
	str, err := s.Get()
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(str)
}

func gettest(w rest.ResponseWriter, r *rest.Request) {
	key := r.PathParam("key")
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
	var ret table.ModelDef 
	json.Unmarshal([]byte(str), &ret)
	fmt.Printf("ret: %v\n", ret)
	w.WriteJson(ret)
}

func getmodeldef(w rest.ResponseWriter, r *rest.Request) {
	key := "getModelDef"
	s, err := dbutil.CreateSession(key)
	if err != nil {
		msg := "err: create session"
		log.Println("msg")
		w.WriteJson(map[string]string{"Err": msg})
		return
	}
	//fmt.Printf("Key: %v\n", key)
	raw, err := s.Get()
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	str := strings.Trim(raw, key+",")
	//fmt.Printf("Json str: %v\n", str)
	var ret table.ModelDef 
	json.Unmarshal([]byte(str), &ret)
	//fmt.Printf("ret: %v\n", ret)
	//w.Header().Del("X-Powered-By")
	//w.Header().Del("Content-Type")
	//w.Header().Del("Date")
	//w.Header().Del("Transfer-Encoding")
	w.Header().Set("Entity",fmt.Sprintf("Content-Length: %v", len(str)))	
	w.Header().Add("Entity","Content-Type: json")	
	w.Header().Set("Miscellaneous","WebGISServer")	
	w.Header().Set("Security","Access-Control-Allow-Headers: Content-Type,Access-Control-Allow-Headers,Authorization,SOAPAction")	
	w.Header().Add("Security","Access-Control-Allow-Methods: GET,POST,PUT,DELETE,OPTIONS")	
	w.Header().Add("Security","Access-Control-Allow-Origin: *")	
	w.Header().Set("Transport", "Connection: close")	
	w.WriteHeader(200)
	w.WriteJson(ret)
}

func put(w rest.ResponseWriter, r *rest.Request) {
	t := table.Table{}
	key := r.PathParam("key")
	err := r.DecodeJsonPayload(&t)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	value, err := table.FromTableToString(&t)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	s, err := dbutil.CreateSession(key, value)
	if err != nil {
		msg := "err: create session"
		log.Println("msg")
		w.WriteJson(map[string]string{"Err": msg})
		return
	}
	err = s.Put()
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(map[string]string{"Err": "0"})
}
