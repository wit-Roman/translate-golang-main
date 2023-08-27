package router

import "net/http"

func addHeaders(w *http.ResponseWriter, r *http.Request) {
	origin := r.Header.Get("Origin")
	acceptedUrl := "http://localhost:9876"
	if origin == "http://localhost:4200" ||
		origin == "http://localhost:3001" ||
		origin == "http://localhost:8080" {
		acceptedUrl = origin
	}

	(*w).Header().Set("Access-Control-Allow-Origin", acceptedUrl)
	(*w).Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Cache-Control, X-Accel-Buffering, Authorization")
	(*w).Header().Set("Access-Control-Expose-Headers", "Authorization")
	(*w).Header().Set("Content-Type", "application/json")
}
