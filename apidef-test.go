package main

import (
	"fmt"
	"github.com/iron-io/jarvis/apidef"
	"os"
	"text/tabwriter"
)

func main() {
	resources, err := apidef.Parse("sample-resources/", "mq")
	if err != nil {
		panic(err)
	}
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	for id, resource := range resources {
		endpoints := resource.BuildEndpoints()
		if len(endpoints) < 1 {
			continue
		}
		fmt.Fprintf(w, "\n%s (%s)\n", resource.Name, id)
		for _, endpoint := range endpoints {
			fmt.Fprintf(w, "\t%s /%s\t%s\n", endpoint.Verb, endpoint.Path, endpoint.Description)
		}
	}
	w.Flush()
}
