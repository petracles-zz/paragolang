package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"log"

	//"github.com/gorilla/mux"
)

// 287c233edb7b4fab84d0ef725335ed41 - jack's account id
// d72de5f52c664cfd9d9da4e952aa1a8e - dream's account id

var epicApiKey, clientId, clientSecret, baseUri string =
"API-KEY", "paraGon", "", "https://developer-paragon.epicgames.com"

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

const loginhtml = `<html><body>
<a href="/login">Log in to your Epic account</a>
</body></html>`

func getIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, loginhtml)
}

func getLogin(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}
	req, err1 := http.NewRequest("GET", "https://developer-paragon.epicgames.com/v1/auth/login/paraGon", nil)
	if err1 != nil {
		log.Fatal("http.NewRequest() FAILURE")
	}
	//req, err1 := http.NewRequest("GET", "https://developer-paragon.epicgames.com/v1/account/asdf/stats", nil)
	req.Header["X-Epic-ApiKey"] = []string{"API-KEY"}
	req.Header["Accept"] = []string{"application/json"}
	w.WriteHeader(http.StatusOK)

	// for key, value := range req.Header {
	// 	fmt.Println(key, value)
	// }
	// fmt.Println(req)

	resp, err2 := client.Do(req)
	if err2 != nil {
		log.Println("client.Do() FAILURE")
	}

	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		bodyBytes, err3 := ioutil.ReadAll(resp.Body)
		if err3 != nil {
			log.Println("ioutil.ReadAll() FAILURE")
		}
		fmt.Fprintln(w, string(bodyBytes))
	}
}