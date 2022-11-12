package entity

import "time"

type NameDetail struct {
	Title string `json:"title"`
	First string `json:"first"`
	Last  string `json:"last"`
}

type StreetDetail struct {
	Number int    `json:"number"`
	Name   string `json:"name"`
}

type LocationDetail struct {
	Street      StreetDetail `json:"street"`
	City        string       `json:"city"`
	State       string       `json:"state"`
	Country     string       `json:"country"`
	Postcode    interface{}  `json:"postcode"`
	Coordinates struct {
		Latitude  string `json:"latitude"`
		Longitude string `json:"longitude"`
	} `json:"coordinates"`
	Timezone struct {
		Offset      string `json:"offset"`
		Description string `json:"description"`
	} `json:"timezone"`
}

type LoginDetail struct {
	UUID     string `json:"uuid"`
	Username string `json:"username"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
	Md5      string `json:"md5"`
	Sha1     string `json:"sha1"`
	Sha256   string `json:"sha256"`
}

type DobDetail struct {
	Date time.Time   `json:"date"`
	Age  interface{} `json:"age"`
}

type RegisteredDetail struct {
	Date time.Time   `json:"date"`
	Age  interface{} `json:"age"`
}

type IDDetail struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type PictureDetail struct {
	Large     string `json:"large"`
	Medium    string `json:"medium"`
	Thumbnail string `json:"thumbnail"`
}

type ResultsData struct {
	Gender     string           `json:"gender"`
	Name       NameDetail       `json:"name"`
	Location   LocationDetail   `json:"location"`
	Email      string           `json:"email"`
	Login      LoginDetail      `json:"login"`
	Dob        DobDetail        `json:"dob"`
	Registered RegisteredDetail `json:"registered"`
	Phone      string           `json:"phone"`
	Cell       string           `json:"cell"`
	ID         IDDetail         `json:"id"`
	Picture    PictureDetail    `json:"picture"`
	Nat        string           `json:"nat"`
}

type info struct {
	Seed    string      `json:"seed"`
	Results interface{} `json:"results"`
	Page    interface{} `json:"page"`
	Version string      `json:"version"`
}

type Randomuser struct {
	Results []ResultsData `json:"results"`
	Info    info          `json:"info"`
}
