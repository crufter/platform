package main

import (
	"fmt"
	"log"
	"os"
	"strings"
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
		tmpl, err := template.New("test").Funcs(template.FuncMap{"toJSON": toJSON}).Parse(serviceTemplate)
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

func typeToDefaultValue(typ string) string {
	switch typ {
	case "string":
		return `""`
	case "int":
	case "int32":
	case "int64":
		return "0"
	case "bool":
		return "false"
	}
	return "{}"
}

// similarly incredibly ugly implementation can be found at https://github.com/micro/platform/blob/master/web/app/src/app/endpoint-list/endpoint-list.component.ts#L68
func toJSON(input *registry.Value, level int) string {
	if input == nil {
		return "no data"
	}
	indent := strings.Repeat("  ", level)
	const fieldSeparator = ",\n"
	if len(input.GetValues()) == 0 {
		if level == 1 {
			return "{}"
		}
		return indent + "\"" + input.GetName() + "\"" + ": " + typeToDefaultValue(input.GetType())
	}
	lines := []string{}
	for _, value := range input.Values {
		lines = append(lines, toJSON(value, level+1))
	}
	if level == 1 {
		return indent + "{\n" + strings.Join(lines, fieldSeparator) + "\n" + indent + "}"
	}
	return indent + "\"" + input.GetName() + "\"" + ": {\n" + strings.Join(lines, fieldSeparator) + "\n" + indent + "}"
}
