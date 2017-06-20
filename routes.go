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
		"/prime/epic/login",
		getLogin,
	},
	Route{
		"Login",
		"GET",
		"/prime/epic/auth/?code={code}",
		getAuthToken,
	},
	// Route{
	// 	"Logout",
	// 	"GET",
	// 	"/logout",
	// 	getLogout,
	// },
	// // - - - - - Accounts routes - - - - -
	Route{
		"Account",
		"GET",
		"/prime/account/{accountId}",
		getAccount,
	},
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