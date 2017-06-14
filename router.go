package main

import (
	"net/http"

	"github.com/gorilla/mux"
)


func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	// router.Headers(
	// 	"Context-Type", "application/json",
	// 	"X-Epic-ApiKey", "API-KEY",
	// )
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}















// Index route
//router.HandleFunc("/", getIndex)

// Auth
// router.HandleFunc("/v1/auth/login/{clientId}", getLogin).Methods("GET")
// router.HandleFunc("/v1/auth/logout/{clientId}", getLogout).Methods("GET")
// router.HandleFunc("/v1/auth/token/{code}", getAuthToken).Methods("GET")

// // Accounts routes
// router.HandleFunc("/v1/account/{accountId}", getAccount).Methods("GET")
// router.HandleFunc("/v1/account/{accountId}/cards", getAccountCards).Methods("GET")
// router.HandleFunc("/v1/account/{accountId}/deck/{deckId}", deleteAccountDeck).Methods("DELETE")
// router.HandleFunc("/v1/account/{accountId}/deck/{deckid}", getAccountDeck).Methods("GET")
// router.HandleFunc("/v1/account/{accountId}/deck/{deckId}", postAccountDeck).Methods("POST")
// router.HandleFunc("/v1/account/{accountId}/decks", getAccountAllDecks).Methods("GET")
// router.HandleFunc("/v1/account/{accountId}/stats", getAccountStats).Methods("GET")
// router.HandleFunc("/v1/account/{accountId}/stats/hero/{heroId}", getAccountHeroStats).Methods("GET")
// router.HandleFunc("/v1/accounts/find", getFindAllAccounts).Methods("GET")
// router.HandleFunc("/v1/accounts/find/{displayname}", getFindAccount).Methods("GET")
// router.HandleFunc("/v1/accounts/find/{platform}/{displayname}", getFindPlatformAccount).Methods("GET")

// // Cards routes
// router.HandleFunc("/v1/card/{id}", getCard).Methods("GET")
// router.HandleFunc("/v1/cards", getAllCards).Methods("GET")
// router.HandleFunc("/v1/cards/complete", getCompleteCards).Methods("GET")

// // Heroes routes
// router.HandleFunc("/v1/hero/{id}", getHero).Methods("GET")
// router.HandleFunc("/v1/heroes", getAllHeroes).Methods("GET")
// router.HandleFunc("/v1/heroes/complete", getCompleteHeroes).Methods("GET")