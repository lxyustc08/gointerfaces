// Package gohttphandler is the test code in go programming language
package gohttphandler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type dollars float32

// mu the mutex between read and update price
var mu sync.Mutex

// htmlReport is the test html report
const htmlReport = `<h1> Items Prices </h1>
<table>
<tr style='text-align: left'>
<th>items</th>
<th>price</th>
</tr>
{{range $key,$value := .}}
<tr>
<td> {{$key}} </td>
<td> {{$value}} </td>
</tr>
{{end}}
</table>`

// HTMLReport the html report to produce output of list
var HTMLReport = template.Must(template.New("htmlReport").Parse(htmlReport))

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

// basic ServeHTTP interface process
//func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
//	for item, price := range db {
//		fmt.Fprintf(w, "%s: %s\n", item, price)
//	}
//}

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	case "/price":
		item := req.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "no such item: %q\n", item)
			return
		}
		fmt.Fprintf(w, "%s\n", price)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such page: %s\n", req.URL)
	}
}

// use servemux to produce a server
func (db database) list(w http.ResponseWriter, req *http.Request) {
	mu.Lock()
	if err := HTMLReport.Execute(w, db); err != nil {
		log.Fatal(err)
	}
	mu.Unlock()
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	mu.Lock()
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		mu.Unlock()
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
	mu.Unlock()
}

func (db database) updatePrice(w http.ResponseWriter, req *http.Request) {
	mu.Lock()
	values := req.URL.Query()
	item := values.Get("item")
	_, ok := db[item]
	if !ok {
		mu.Unlock()
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	price := values.Get("price")
	afterprice, err := strconv.ParseFloat(price, 32)
	if err != nil {
		mu.Unlock()
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "invalid price: %v\n", err)
		return
	}
	db[item] = dollars(afterprice)
	mu.Unlock()
}

// TestHTTPMux use servemux to create a server
func TestHTTPMux() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/price", http.HandlerFunc(db.price))
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

// TestEcommerce is the function for test
func TestEcommerce() {
	db := database{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}

// TestDefaultHTTPMux is the test function to test default http mux serve
func TestDefaultHTTPMux() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/update", db.updatePrice)
	log.Fatal(http.ListenAndServe("172.22.126.98:8000", nil))
}
