package main

import (
	"bufio"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

// Leak sturct represents a slice of bucket leaks in yas3bl.json file
type Leak struct {
	Count        string `json:"count"`
	Data         string `json:"data"`
	Organization string `json:"organization"`
	URL          string `json:"url"`
}

// YAS3BL (yet another s3 bucket leak) is a struct that holds all leak instances
type YAS3BL struct {
	Leaks []Leak `json:"yas3bl"`
}

func main() {
	const tmpl = `# YAS3BL (Yet Another S3 Bucket Leak)

> 🔓 Enumerating all the AWS S3 bucket leaks that have been discovered to date.

| Company | Link | Records Exposed | Data |
| ------- | ---- | --------------- | ---- |
{{range .Leaks}}| <h4>{{.Organization}}</h4> | [🔗]({{.URL}}) | {{.Count}} | {{.Data}} |
{{end}}
`
	jsonBytes, err := ioutil.ReadFile("yas3bl.json")
	if err != nil {
		log.Fatalf("could not read file: %+v\n", err)
	}

	var bucketsLeaked YAS3BL
	err = json.Unmarshal(jsonBytes, &bucketsLeaked)
	if err != nil {
		log.Fatalf("could not unmarshal JSON: %+v\n", err)
	}

	sort.Slice(bucketsLeaked.Leaks, func(i, j int) bool {
		return strings.ToUpper(bucketsLeaked.Leaks[i].Organization) < strings.ToUpper(bucketsLeaked.Leaks[j].Organization)
	})

	f, err := os.Create("README.md")
	if err != nil {
		log.Fatalf("could not create README.md file: %+v\n", err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	t := template.Must(template.New("tmpl").Parse(tmpl))
	err = t.Execute(w, bucketsLeaked)
	if err != nil {
		log.Fatalf("could not merge data from JSON into README.md: %+v\n", err)
	}
	w.Flush()
}
