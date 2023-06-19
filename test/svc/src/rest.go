package main

import (
	"fmt"
	"log"
	"net/http"
	//"sync"

	//"time"

	//"k8s.io/apimachinery/pkg/api/errors"
	//metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	//"k8s.io/client-go/kubernetes"
	//rest0 "k8s.io/client-go/rest"
	//v1beta1 "k8s.io/api/extensions/v1beta1"
	"github.com/ant0ine/go-json-rest/rest"
	incluster "inClusterServiceDiscovery"
)

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/api/snap/:key", get),
		rest.Put("/api/snap/:key", put),
		rest.Put("/api/svc", svc),
		rest.Post("/api/svc", svc),
		rest.Get("/api/svc/db", dbsvc),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

type Info struct {
	Key   string
	Value string
}

type Svc struct {
	Ip  string
	Err string
}

func get(w rest.ResponseWriter, r *rest.Request) {
	key := r.PathParam("key")
	fmt.Printf("Key: %v\n", key)
	info := Info{}
	value := "test"
	info.Key = key
	info.Value = value
	fmt.Println(info)
	w.WriteJson(info)
}

func put(w rest.ResponseWriter, r *rest.Request) {
	info := Info{}
	key := r.PathParam("key")
	info.Key = key
	err := r.DecodeJsonPayload(&info)
	fmt.Println(info)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(&info)
}

func svc(w rest.ResponseWriter, r *rest.Request) {
	ret := Svc{}
	c := incluster.Config{}
	err := r.DecodeJsonPayload(&c)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	s, err := incluster.CreateSearch(&c)
	if err != nil {
		ret.Ip = ""
		ret.Err = fmt.Sprintf("%v", err)
		w.WriteJson(ret)
		return
	}
	s.Print()
	ips, err := s.Result()
	if err != nil {
		ret.Ip = ""
		ret.Err = fmt.Sprintf("%v", err)
		w.WriteJson(ret)
		return
	}
	ret.Ip = ips
	ret.Err = "0"
	w.WriteJson(ret)
}

func dbsvc(w rest.ResponseWriter, r *rest.Request) {
	ret := Svc{}
	c := incluster.Config{}
	c.Namespace = "database"
	c.ControllerName = "database"
	c.ControllerType = "daemonset"
	c.Service = "database"
	s, err := incluster.CreateSearch(&c)
	if err != nil {
		ret.Ip = ""
		ret.Err = fmt.Sprintf("%v", err)
		w.WriteJson(ret)
		return
	}
	s.Print()
	ips, err := s.Result()
	if err != nil {
		ret.Ip = ""
		ret.Err = fmt.Sprintf("%v", err)
		w.WriteJson(ret)
		return
	}
	ret.Ip = ips
	ret.Err = "0"
	w.WriteJson(ret)
}
