package main

import (
	"flag"
	"fmt"
	"log"
	"mymodule/golang/yourownadventure"
	"net/http"
	"os"
	"strings"
	"text/template"
)

func main() {

	port := flag.Int("port", 3000, "the port to start the server")
	fileName := flag.String("file", "gopher.json", "json file where the story is stored.")
	flag.Parse()

	os.ReadFile(*fileName)

	f, err := os.Open(*fileName)
	if err != nil {
		fmt.Println("Couldnt open the provided file.")
	}

	story, err := yourownadventure.JsonStory(f)
	if err != nil {
		fmt.Println("Couldnt decode the story provided.")
	}

	tpl := template.Must(template.New("").Parse(myTpl))

	h := yourownadventure.NewHandler(story, yourownadventure.WithPathFunc(pathFn), yourownadventure.WithTemplate(tpl))

	mux := http.NewServeMux()
	mux.Handle("/story/", h)
	fmt.Printf("Starting server on port %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h), mux)
}

func pathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "/story" || path == "/story/" {
		path = "/story/intro"
	}
	return path[len("/story/"):]
}

var myTpl = `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Choose your own story.</title>
  </head>
  <body>
  <section class="page">
    <h1>{{.Title}}</h1>
    {{range .Paragraphs}}
    <p>{{.}}</p>
    {{end}}
    <ul>
      {{range .Options}}
      <li><a href="/story/{{.Chapter}}">{{.Text}}</a></li>
      {{end}}
    </ul>
	</section>
	<style>
      body {
        font-family: helvetica, arial;
      }
      h1 {
        text-align:center;
        position:relative;
      }
      .page {
        width: 80%;
        max-width: 500px;
        margin: auto;
        margin-top: 40px;
        margin-bottom: 40px;
        padding: 80px;
        background: #FFFCF6;
        border: 1px solid #eee;
        box-shadow: 0 10px 6px -6px #777;
      }
      ul {
        border-top: 1px dotted #ccc;
        padding: 10px 0 0 0;
        -webkit-padding-start: 0;
      }
      li {
        padding-top: 10px;
      }
      a,
      a:visited {
        text-decoration: none;
        color: #6295b5;
      }
      a:active,
      a:hover {
        color: #7792a2;
      }
      p {
        text-indent: 1em;
      }
    </style>
  </body>
</html>`
