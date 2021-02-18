package router

import (
	"../controllers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	//serving
	Route{
		"Index",
		"GET",
		"/clients",
		controllers.GetClients,
	},
	Route{
		"New Client",
		"POST",
		"/clients",
		controllers.NewClient,
	},
	Route{
		"Upload file",
		"POST",
		"/upload",
		controllers.UploadFile,
	},
	Route{
		"Upload multiple file",
		"POST",
		"/upload-multiple",
		controllers.MultipleFileUpload,
	},
	Route{
		"Get uploaded images",
		"GET",
		"/public/uploads/{name}",
		controllers.GetUploadedImages,
	},
}

func NewRouter() http.Handler {
	router := mux.NewRouter().StrictSlash(true)
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // All origins
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	})
	//router.PathPrefix("/public/js/").Handler(http.StripPrefix("/public/js/", http.FileServer(http.Dir("./public/js/"))))
	//router.PathPrefix("/public/css/").Handler(http.StripPrefix("/public/css/", http.FileServer(http.Dir("./public/css/"))))
	//router.PathPrefix("/public/uploads/").Handler(http.FileServer(http.Dir("/public/uploads/")))

	router.NotFoundHandler = http.HandlerFunc(controllers.NotFound)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		router.
			Methods(route.Method).
			Path("/api" + route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	var imgServer = http.FileServer(http.Dir("./public/uploads/"))
	router.PathPrefix("/public/uploads").Handler(http.StripPrefix("/public/uploads/", imgServer))
	return c.Handler(router)
}
