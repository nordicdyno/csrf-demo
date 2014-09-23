package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"go/build"

	"path/filepath"
	//"github.com/carbocation/gotogether"
	"bitbucket.org/kardianos/osext"
	// "github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"

	"github.com/aarzilli/golua/lua"
	"github.com/davecgh/go-spew/spew"
)

var (
	bind    = flag.String("bind", ":8080", "[host]:port where to serve on")
	asserts = flag.String("asserts", getWorkDir(), "path to assers")
	verbose = flag.Bool("verbose", false, "path to assers")
)

func getWorkDir() string {
	//filename, _ := osext.Executable()
	// fmt.Println(filename)
	//return filepath.Join(filepath.Dir(filename), "resources")

	p, err := build.Default.Import("github.com/nordicdyno/csrf-demo", "", build.FindOnly)
	if err != nil {
		filename, _ := osext.Executable()
		return filepath.Join(filepath.Dir(filename), "resources")
	}

	return filepath.Join(p.Dir, "resources")
}

func init() {
	flag.Parse()
}

type appHandler struct {
	router *mux.Router
}

// www.evil.ro
func extractHost(s string) string {
	if !strings.Contains(s, ":") {
		return s
	}

	host, _, _ := net.SplitHostPort(s)
	return host
}

func extractPort(s string) string {
	if !strings.Contains(s, ":") {
		return ""
	}

	_, port, _ := net.SplitHostPort(s)
	return ":" + port
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Host: %s\n", extractHost(req.Host))
	fmt.Fprintf(w, "URI: %s\n", req.RequestURI)
}

func main() {
	//gotogether.Handle("/resources/")
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler)

	h := appHandler{router: r}
	http.Handle("/", h)

	//http.HandleFunc("/", indexHandler)
	if err := http.ListenAndServe(*bind, nil); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}

func (h appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if rec := recover(); rec != nil {
			log.Println("Catch error. Recovering...")
			var doc bytes.Buffer
			err := errorTemplate.Execute(&doc, &ErrorPage{
				Code:    http.StatusInternalServerError,
				Message: rec,
			})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "text/html;charset=utf-8")
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, doc.String())
		}
	}()

	dir := filepath.Join(*asserts, extractHost(r.Host))
	if *verbose {
		log.Println("Check dir ", dir)
	}
	//fmt.Println("Check dir ", dir)
	if _, err := os.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			panic("Dir" + dir + " not exists")
		}
		panic(err)
	}

	if strings.HasSuffix(r.URL.Path, ".lua") {
		serveLua(dir, w, r)
		return
	}

	if r.URL.Path == "/" || strings.HasSuffix(r.URL.Path, ".html") {
		serveTemplate(dir, w, r)
		return
	}
	// TODO: check if file exists, in other case try template
	if *verbose {
		log.Println("Static serve ", r.URL.Path, "on", r.Host)
	}
	fs := http.FileServer(http.Dir(dir))
	fs.ServeHTTP(w, r)
	return

	//h.router.ServeHTTP(w, r)
}

type HtmlContext struct {
	Port string
}

func serveTemplate(dir string, w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path
	if name == "/" {
		name = "index.html"
	}
	file := filepath.Join(dir, name)
	if *verbose {
		log.Println("Serve template file: ", file)
	}
	content, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	var t = template.Must(template.New(file).Parse(string(content)))

	ctx := HtmlContext{Port: extractPort(r.Host)}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t.ExecuteTemplate(w, file, &ctx)
	//t.ExecuteTemplate(os.Stdout, file, &ctx)
	//io.WriteString(w, s)
	//io.WriteString(w, "\n")
}

func serveLua(dir string, w http.ResponseWriter, r *http.Request) {
	file := filepath.Join(dir, r.URL.Path)
	/*
		fmt.Fprintf(w, "Dir: %s\n", dir)
		fmt.Fprintf(w, "URL.Path: %s\n", r.URL.Path)
		fmt.Fprintf(w, "file: %s\n", file)
	*/
	if *verbose {
		log.Println("Serve lua file: ", file)
	}
	content, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	log.Println("For file", file, "Host is: ", r.Host)

	luaPrint := func(L *lua.State) int {
		s := L.ToString(1)
		io.WriteString(w, s)
		io.WriteString(w, "\n")
		if *verbose {
			log.Println(s)
		}
		if f, ok := w.(http.Flusher); ok {
			f.Flush()
		}
		return 0
	}
	luaGetHeaderValue := func(L *lua.State) int {
		s := L.ToString(1)
		log.Println("Get header", s)
		var val string
		if s == "Host" {
			val = r.Host
		} else {
			val = r.Header.Get(s)
			log.Println("got header value", val)
		}
		L.PushString(val)
		return 1
	}
	luaSetHeaderValue := func(L *lua.State) int {
		s1 := L.ToString(1)
		s2 := L.ToString(2)
		w.Header().Set(s1, s2)
		return 0
	}

	luaGetCookieValue := func(L *lua.State) int {
		s := L.ToString(1)
		log.Println("Get ", s)
		c, err := r.Cookie(s)
		if err != nil {
			L.PushString("")
			return 0
		}
		log.Println("got value", c.Value)
		L.PushString(c.Value)
		return 1
	}
	luaSetCookieValue := func(L *lua.State) int {
		s1 := L.ToString(1)
		s2 := L.ToString(2)
		//log.Println(s1, " ... ", s2)
		c := http.Cookie{
			Name:  s1,
			Value: s2,
			Path:  "/",
		}
		http.SetCookie(w, &c)
		return 0
	}
	luaGetOriginValue := func(L *lua.State) int {
		log.Println("Get origin", r.Proto)
		spew.Dump(strings.SplitN(r.Proto, "/", 2))
		//val := r.RemoteAddr
		val := strings.ToLower(strings.SplitN(r.Proto, "/", 2)[0]) + "://" + r.Host
		L.PushString(val)
		return 1
	}

	L := lua.NewState()
	defer L.Close()
	L.OpenLibs()
	L.Register("print", luaPrint)
	L.Register("cookie_get", luaGetCookieValue)
	L.Register("cookie_set", luaSetCookieValue)
	L.Register("header_get", luaGetHeaderValue)
	L.Register("header_set", luaSetHeaderValue)
	L.Register("origin_get", luaGetOriginValue)
	L.MustDoString(string(content))
}

func DumpStack(lS *lua.State) {
	top := lS.GetTop()
	dumpTxtMessage := "---- Begin Stack ----\n"
	dumpTxtMessage += fmt.Sprintf("Stack size: %v\n\n", top)
	for i := top; i >= 1; i-- {
		t := int(lS.Type(i))
		dumpTxtMessage += fmt.Sprintf("  %v -- (%v) -- %s\n", i, i-top-1, lS.Typename(t))
	}
	log.Println(dumpTxtMessage, "\n")
}

type ErrorPage struct {
	Code    int
	Message interface{}
}

var errorTemplate = template.Must(template.New("").Parse(`
<html><body>
<h2>This app is crashed with error:</h2>
<h2>Code: {{.Code}}<br>
Message: «{{.Message}}»
</h2>
</body></html>
`))
