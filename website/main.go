package main

import(
  "github.com/thinkerou/favicon"
  "github.com/gin-gonic/gin"
  "html/template"
  "net/http"
  "log"
)

type row struct {
	Data []string
}

type tableData struct {
	NotEmpty bool
	Rows     []row
}

var(
  webApp *gin.Engine
  tableTemplate  *template.Template
)

func init() {
  tableTemplate = template.Must(template.ParseFiles("template/index.html", "template/header.html"))
}

func renderTable() {}

func test(w http.ResponseWriter, r *http.Request) {
  defer r.Body.Close()
	data := struct {
		IP   string
	}{
		r.RemoteAddr,
	}

  w.Header().Set("Content-Type", "text/html")

  if err := tableTemplate.ExecuteTemplate(w, "main", data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Failed to execute table template: %v", err)
	}
}

func main() {
  webApp = gin.Default()
  webApp.Delims("{{", "}}")
  webApp.SetFuncMap(template.FuncMap{

	})

  webApp.Use(favicon.New("static/img/favicon.ico"))

  webApp.LoadHTMLGlob("template/*")

  webApp.GET("/", func(c *gin.Context) {

  // Call the HTML method of the Context to render a template
  c.HTML(
      // Set the HTTP status to 200 (OK)
      http.StatusOK,
      // Use the index.html template
      "index.html",
      // Pass the data that the page uses (in this case, 'title')
      gin.H{
          "title": "Home Page",
      },
  )

})

  webApp.Run(":8080")
}
