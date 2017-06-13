package main

import (
	"fmt"
	"io/ioutil"
	//"strconv"
	"net/http"
	//"net/http/httputil"
	//"log"

	//"github.com/gorilla/mux"
)


var epicApiKey, clientId, clientSecret, baseUri string =
"API KEY", "paraGon", "", "https://developer-paragon.epicgames.com"

// A person with an Epic account using this API
type Account struct {
	ID string
	Name string
}

// A hero from Paragon
type Hero struct {
	ID string
	Name []string
	Attack string
	Traits []string
	Scale string
	ReleaseDate string
	Affinities []string
	Difficulty string
	//Stats []
	//Images []
	Abilities []Abilities
}

type Abilities struct {
	Name []string
	Description []string
	ShortDescription []string
	Type string
	Binding string
	DamageType string
	MaxLevel int
	//ModifiersByLevel []
	//Images []
}

type Card struct {
	ID string
	Name []string
	SlotType string
	//Images []
	Rarity string
	Affinities []string
	Cost int
	UpgradeSlots int
	Effects []Effect
	MaxedEffects []Effect
}

type Effect struct {
	Description string
	//Value []
	Passive bool
	Unique bool
	Cooldown int
}

// func createUrl(apiVersion int, uri string) string {
// 	return baseUri + "/v" + strconv.Itoa(apiVersion) + "/" + uri
// }

// func makeRequest(httpMethod string, apiVerstion int, uri string, authenticationHeader []string, callback string, requestBody string) {
// 	// Set up headers
// 	var headers = make(map[string]string)
// 	headers["X-Epic-ApiKey"] = epicApiKey


// }


const loginhtml = `<html><body>
<a href="/login">Log in to your Epic account</a>
</body></html>`

func getIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, loginhtml)
}

func getLogin(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}
	req, err1 := http.NewRequest("GET", "https://developer-paragon.epicgames.com/v1/auth/login/paraGon", nil)
	//req, err1 := http.NewRequest("GET", "https://developer-paragon.epicgames.com/v1/account/asdf/stats", nil)
	req.Header["X-Epic-ApiKey"] = []string{"API-KEY"}
	req.Header["Accept"] = []string{"application/json"}
	w.WriteHeader(http.StatusOK)
	if err1 != nil {
		// log.Fatal("")
	}

	for key, value := range req.Header {
		fmt.Println(key, value)
	}
	fmt.Println(req)

	resp, err2 := client.Do(req)
	if err2 != nil {
		// log.Fatal("")
	}

	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		bodyBytes, err3 := ioutil.ReadAll(resp.Body)
		if err3 != nil {
			// log.Fatal("")
		}
		fmt.Fprintln(w, string(bodyBytes))
	}
}