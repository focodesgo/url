package main

import (    
    "net/http"	
)

var urlMap map[string]string

func main() {	
	setUpRoutes()	
	http.HandleFunc("/", handler)	
    http.ListenAndServe(":8080", nil)	
}

func setUpRoutes() {
	urlMap = make(map[string]string)
	urlMap["a"] = "https://google.com"	
	urlMap["b"] = "https://github.com"	
	urlMap["default"] = "https://twitter.com"		
}

func handler(w http.ResponseWriter, r *http.Request) {
    keys, ok := r.URL.Query()["key"]

    if !ok || len(keys[0]) < 1 {
		http.Redirect(w, r, urlMap["default"], 301)	
        return
    }

    key := keys[0]	

	if url, exists := urlMap[key]; exists {		
		http.Redirect(w, r, url, 301)
		return
	}

    http.Redirect(w, r, urlMap["default"], 301)	
}