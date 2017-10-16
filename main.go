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

const tmpl = `# YAS3BL (Yet Another S3 Bucket Leak)

> 🔓 Enumerating all the AWS S3 bucket leaks that have been discovered to date.

| Company | Link | Records Exposed | Data |
| ------- | ---- | --------------- | ---- |
{{range .Leaks}}| <h4>{{.Organization}}</h4> | [🔗]({{.URL}}) | {{.Count}} | {{.Data}} |
{{end}}
`

const htmlTmpl = `
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<meta http-equiv="X-UA-Compatible" content="ie=edge">
	<title>YAS3BL</title>
	<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-beta/css/bootstrap.min.css" integrity="sha384-/Y6pD6FV/Vv2HJnA6t+vslU6fwYXjCFtcEpHbNJ0lyAFsXTsjBbfaDjzALeQsN6M" crossorigin="anonymous">
</head>
<style>
td {
	font-family: monospace;
	font-size: 1.2em;
}
a {
	text-decoration: none;
}
</style>
<body>
	<div class="container">
		<div class="header clearfix">
			<h1 class="text-muted text-center">Yet Another S3 Bucket Leak</h1>
		</div>
		<table class="table table-striped">
			<thead>
				<tr>
					<th>Organization</th>
					<th>Count</th>
					<th>Data</th>
				</tr>
			</thead>
			<tbody>
			{{range .Leaks}}<tr>
				<td><a href="{{.URL}}">{{.Organization}}</a></td>
				<td>{{.Count}}</td>
				<td>{{.Data}}</td>
			</tr>
			{{end}}
			</tbody>
		</table>
	</div>
	<footer class="footer fixed-bottom">
		<div class="container-fluid"><span class="text-muted">MIT &copy; 2017 <a href="https://github.com/petermbenjamin/YAS3BL">Peter Benjamin</a></span></div>
	</footer>
	<script src="https://code.jquery.com/jquery-3.2.1.slim.min.js" integrity="sha384-KJ3o2DKtIkvYIK3UENzmM7KCkRr/rE9/Qpg6aAZGJwFDMVNA/GpGFF93hXpG5KkN" crossorigin="anonymous"></script>
	<script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.11.0/umd/popper.min.js" integrity="sha384-b/U6ypiBEHpOf/4+1nzFpr53nxSS+GLCkfwBdFNTxtclqqenISfwAzpKaMNFNmj4" crossorigin="anonymous"></script>
	<script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-beta/js/bootstrap.min.js" integrity="sha384-h0AbiXch4ZDo7tp9hKZ4TsHbi047NrKGLO3SEJAg45jXxnGIfYzk4Si90RDIqNm1" crossorigin="anonymous"></script>
</body>
</html>
`

func main() {
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

	hf, err := os.Create("docs/index.html")
	if err != nil {
		log.Fatalf("could not create index.html file: %+v\n", err)
	}
	defer hf.Close()

	hw := bufio.NewWriter(hf)
	ht := template.Must(template.New("tmpl").Parse(htmlTmpl))
	err = ht.Execute(hw, bucketsLeaked)
	if err != nil {
		log.Fatalf("could not merge data from JSON into index.html: %+v\n", err)
	}
	hw.Flush()
}
