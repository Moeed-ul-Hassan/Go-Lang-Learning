package main

//Purpose of making this file is "Defining How our data will look like, Here we will define Structs and other."
type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}
