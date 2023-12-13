// package main

package main

import (
	"fmt"
	"html"

	"log"
	"net/http"
	"os"
)

var (
	CertFilePath = ".certs/tls_server.crt"
	KeyFilePath  = ".certs/tls_server.key"
)

func startWebServer(addr string, port string) {
	count := 0
	http.HandleFunc("/ammune/log", func(w http.ResponseWriter, r *http.Request) {
		count += 1
		switch r.Method {
		case "GET":
			log.Printf("Finished GET for log message request #%v", count)
			fmt.Fprintf(w, "success processed log message endpoint, %q", html.EscapeString(r.URL.Path))
		case "POST":
			log.Printf("Finished POST request for log message #%v", count)
			fmt.Fprintf(w, "success processed log message endpoint, %q", html.EscapeString(r.URL.Path))
		default:
			fmt.Fprintf(w, "Sorry, only GET and POST methods are supported for log message.")
		}
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		count += 1
		if (count % 1000) == 0 {
			log.Printf("Finished GET request #%v", count)
		}
		fmt.Fprintf(w, "Hello from default endpoint: welcome to the secure connection that we have made, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		//log.Println("hello endpoint")
		if (count % 1000) == 0 {
			log.Printf("Finished GET request for hello #%v", count)
		}

		fmt.Fprintf(w, "Hello from hello endpoint, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/world", func(w http.ResponseWriter, r *http.Request) {
		count := 0
		count += 1
		if (count % 1000) == 0 {
			log.Printf("Finished GET request for world #%v", count)
		}

		fmt.Fprintf(w, "Hello from world endpoint, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/foo/bar", func(w http.ResponseWriter, r *http.Request) {
		count := 0
		count += 1
		if (count % 1000) == 0 {
			log.Printf("Finished GET request for foo/bar #%v", count)
		}

		fmt.Fprintf(w, "Hello from foo/bar endpoint, %q", html.EscapeString(r.URL.Path))
	})

	if len(addr) != 0 {
		url := addr + ":" + port
		http.ListenAndServeTLS(url, CertFilePath, KeyFilePath, nil)
	} else {
		url := ":" + port
		http.ListenAndServeTLS(url, CertFilePath, KeyFilePath, nil)
	}

}

func main() {
	args := os.Args
	var port = "4443"
	var addr = ""
	if len(args) > 1 {
		port = args[1]
	}
	if len(args) > 2 {
		log.Printf("we have %d args to read from ", len(args))
		addr = args[1]
	}
	log.Printf("starting server at %s", addr)

	startWebServer(addr, port)
	log.Println("server is up and running")
}
