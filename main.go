package main

import (
	"github.com/Remydeme/iurgence/api"
	"github.com/Remydeme/iurgence/api/middleware"
	"github.com/Remydeme/iurgence/api/middleware/logger"
	"github.com/Remydeme/iurgence/config"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"log"
	"net/http"
	"time"
)

const (
	SCHEMES = "http"
)

func main() {

	// use alice
	// context.ClearHandler => Used to remove the resuest in the gorilla context
	//

	middlewares := alice.New(logger.Log, middleware.RecoverHandler, context.ClearHandler)
	authentifiedMiddlewares := alice.New(logger.Log, middleware.JWTMiddleware, middleware.RecoverHandler, context.ClearHandler) // middleware + check JWT token

	// create new router
	r := mux.NewRouter()

	r.Handle("/ping/", middlewares.ThenFunc(api.Ping)).
		Methods("GET")

	r.Handle("/v1/signup/", middlewares.ThenFunc(api.SignUp)).
		Methods("POST")

	r.Handle("/v1/signin/", middlewares.ThenFunc(api.SignIn)).
		Methods("POST")

	r.Handle("/v1/user/add/", authentifiedMiddlewares.ThenFunc(api.CreateUser)).
		Methods("POST")

	r.Handle("/v1/user/delete/", authentifiedMiddlewares.ThenFunc(api.DeleteUser)).
		Methods("PUT")

	srv := &http.Server{
		Handler: r,
		// Good practice: enforce timeouts
		//for servers you create!
		Addr:         config.Main.Server.Port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
