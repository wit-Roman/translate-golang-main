package router

import "net/http"

func addHeaders(w *http.ResponseWriter, r *http.Request) {
	origin := r.Header.Get("Origin")
	acceptedUrl := "https://translate.apps-web.xyz"
	if 
		origin == "https://apps-web.xyz" ||
		origin == "http://localhost:4200" ||
		origin == "http://localhost:3001" ||
		origin == "http://localhost:8080" ||
		origin == "http://localhost:8081" {
		acceptedUrl = origin
	}

	(*w).Header().Set("Access-Control-Allow-Origin", acceptedUrl)
	(*w).Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Cache-Control, X-Accel-Buffering, Authorization")
	(*w).Header().Set("Access-Control-Expose-Headers", "Authorization")
	(*w).Header().Set("Content-Type", "application/json")
}

func addHeadersSSE(w *http.ResponseWriter) {
	(*w).Header().Set("Connection", "keep-alive")
	(*w).Header().Set("Content-Type", "text/event-stream;charset=utf-8")
	(*w).Header().Set("Cache-Control", "no-cache, no-transform")
	(*w).Header().Set("X-Accel-Buffering", "no")

}

func addTokenHeader(w *http.ResponseWriter, token string) {
	(*w).Header().Set("Authorization", token)
}
