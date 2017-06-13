package main

import "net/http"


type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	// - - - - - Index route - - - - -
	Route{
		"Index",
		"GET",
		"/",
		getIndex,
	},
	// // - - - - - Auth routes - - - - -
	Route{
		"Login",
		"GET",
		"/login",
		getLogin,
	},
	// Route{
	// 	"Logout",
	// 	"GET",
	// 	"/logout",
	// 	getLogout,
	// },
	// Route{
	// 	"AuthToken",
	// 	"GET",
	// 	"/auth",
	// 	getAuthToken,
	// },
	// // - - - - - Accounts routes - - - - -
	// Route{
	// 	"Account",
	// 	"GET",
	// 	"/account",
	// 	getAccount,
	// },
	// Route{
	// 	"AccountCards",
	// 	"GET",
	// 	"/account/cards",
	// 	getAccountCards,
	// },
	// // - - - - - Cards routes - - - - -
	// Route{
	// 	"TodoShow",
	// 	"GET",
	// 	"/todos/{todoId}",
	// 	TodoShow,
	// },
	// // - - - - - Heroes routes - - - - -
	// Route{
	// 	"TodoShow",
	// 	"GET",
	// 	"/todos/{todoId}",
	// 	TodoShow,
	// },
}