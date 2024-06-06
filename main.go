package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func init() {
	result := GenerateShortUrl()
	fmt.Println(result)
}
func main() {

	http.HandleFunc("/", root)
	http.HandleFunc("/short", short)
	http.HandleFunc("/long", long)

	fmt.Println("server running")
	log.Printf("%s", http.ListenAndServe("localhost:9000", nil))

}

var rootHtmlTmpl = template.Must(template.New("rootHtml").Parse(`
       <html>
           <body>
               <h1>URL SHORTERNER</h1>
                {{if .}}{{.}}<br />{{end}}
                <form action="/short" type="POST">
                    shorten this: <input type="text" name="longUrl" />
                    <input  type="submit" value="give me the long url boy!!"  />
                </form>
                <br  />
                <form action=”/long” type=“POST”>
					Expand this: http://goo.gl/<input type=“text” name=“shortUrl” />
                    <input type=“submit” value=“Give me the long URL” />
                </form>
           </body>
        </html>
`))

func root(w http.ResponseWriter, r *http.Request) {
	err := rootHtmlTmpl.Execute(w, nil)
	if err != nil {
		log.Fatalf("%s", err)
	}
}

func short(w http.ResponseWriter, r *http.Request) {

}

func long(w http.ResponseWriter, r *http.Request) {}
