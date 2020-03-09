package main

import (
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/micro/clients/go/client"
	registry "github.com/micro/go-micro/registry/service/proto"
)

func main() {
	c := client.NewClient(nil)
	req := registry.ListRequest{}
	rsp := registry.ListResponse{}
	if err := c.Call("go.micro.registry", "Registry.ListServices", req, &rsp); err != nil {
		fmt.Println(err)
		return
	}
	index := map[string]struct{}{}
	for _, service := range rsp.GetServices() {
		_, ok := index[service.GetName()]
		if ok {
			continue
		}
		tmpl, err := template.New("test").Parse(serviceTemplate)
		if err != nil {
			panic(err)
		}
		f, err := os.Create("./" + service.GetName() + ".md")
		if err != nil {
			log.Println("create file: ", err)
			return
		}
		err = tmpl.Execute(f, map[string]interface{}{"service": service})
		if err != nil {
			panic(err)
		}
	}
}
