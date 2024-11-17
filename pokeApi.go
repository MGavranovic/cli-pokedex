package main

type PokeResults struct {
	Name string `json: "name"`
	Url  string `json: "url"`
}

type PokeData struct {
	Count    int           `json: "count"`
	Next     string        `json: "next"`
	Previous string        `json: "previous"`
	Results  []PokeResults `json: "results"`
}
