package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

var (
	mainPage *template.Template
	app      App
)

func update(w http.ResponseWriter, r *http.Request) {
	app.Time = time.Now().UnixNano()
	mainPage.Execute(w, app)
}

func activateCheat(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["id"]
	cheatActive := 0
	if ok {
		n, err := strconv.Atoi(keys[0])
		if err != nil {
			panic(err)
		}
		cheatActive = n
	}
	for i, _ := range app.Cheats {
		app.Cheats[i].Active = false
	}

	app.Cheats[cheatActive].Active = true
	mainPage.Execute(w, app)
}

func index(w http.ResponseWriter, r *http.Request) {
	mainPage.Execute(w, app)
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	filePath := r.URL.Path
	headers := w.Header()
	if strings.HasSuffix(filePath, "css") {
		headers["Content-Type"] = []string{"text/css"}
	} else if strings.HasSuffix(filePath, "js") {
		headers["Content-Type"] = []string{"application/javascript"}
	} else if strings.HasSuffix(filePath, "woff") {
		headers["Content-Type"] = []string{"application/x-font-woff"}
	}

	// b := MustAsset("public/static/app.css")
	println(filePath)
	b, _ := ioutil.ReadFile("public" + filePath)
	w.Write(b)
}

func server() {
	router := mux.NewRouter().StrictSlash(true)

	router.PathPrefix("/static/").HandlerFunc(staticHandler)
	router.HandleFunc("/update", update)
	router.HandleFunc("/activate/", activateCheat)
	router.HandleFunc("/", index)

	log.Println("Listen on port 8080...")
	http.ListenAndServe(":8080", router)
}

func main() {

	// fb := MustAsset("public/index.html")
	fb, _ := ioutil.ReadFile("public/index.html")

	mainPage = template.Must(template.New("app").Parse(string(fb)))

	app = App{
		Version: "0.0.1",
		Time:    time.Now().Unix(),
		Cheats: []Cheats{
			Cheats{
				Name:   "Git",
				Active: false,
				Notes: []Cheat{
					Cheat{
						Name:        "Teste",
						Description: "testing this",
					},
				},
			},
			Cheats{
				Name:   "Bash",
				Active: false,
				Notes: []Cheat{
					Cheat{
						Name:        "Teste bash",
						Description: "testing bash",
					},
				},
			},
		},
	}

	// w := webview.New(webview.Settings{
	// 	Title:     "MyCheatSheet",
	// 	URL:       "http://192.168.0.113:8080/",
	// 	Resizable: true,
	// 	Width:     1367,
	// 	Height:    766,
	// })

	go server()

	time.Sleep(time.Second)

	// w.Run()
	// os.Exit(0)
	d := make(chan bool)
	<-d

}
