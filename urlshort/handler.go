package urlshoet

import (
	"net/http"
)

func MapHandler(paths map[string]string , fallback http.Handler) http.HandlerFunc{
	return func( w http.ResponseWriter, r *http.Request ){
		Rpath := r.URL.Path
		val , ok :=paths[Rpath] 

		if ok{
			http.Redirect(w, r, val, http.StatusFound)
			return 
		}

		fallback.ServeHTTP(w, r)
 
	}

}