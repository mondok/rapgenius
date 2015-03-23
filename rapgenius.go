package rapgenius

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func checkErr(e error) bool {
	if e != nil {
		log.Print("Error")
		log.Print(e)
		return true
	}
	return false
}

type searchResponse struct {
	Result *searchResults `json:"response"`
}

type artistResponseRoot struct {
	ArtistResponse *artistResponse `json:"response"`
}

type artistResponse struct {
	Artist *Artist `json:"artist"`
}

type searchResults struct {
	SearchHits []*SearchHit `json:"hits"`
}

// SearchHit is a single item in the
// search results
type SearchHit struct {
	Item *SearchItem `json:"result"`
}

// SearchItem is a single search
// item from API
type SearchItem struct {
	AnnotationCount int     `json:"annotation_count"`
	ID              int     `json:"id"`
	Title           string  `json:"title"`
	Artist          *Artist `json:"primary_artist"`
}

// Artist is a musician
type Artist struct {
	ID    int    `json:"id"`
	Image string `json:"image_url"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

// RapGenius is a new instance of RapGenius
// HTTP client
type RapGenius struct {
	BaseURL string
	Client  *http.Client
}

// New creates new instance of RapGenius
func New() *RapGenius {
	return &RapGenius{
		BaseURL: "http://api.rapgenius.com",
		Client:  &http.Client{},
	}
}

// Executes the query
func (h *RapGenius) execute(path string, response interface{}) (err error) {
	url := h.BaseURL + "/" + path
	req, err := http.NewRequest("GET", url, nil)
	if !checkErr(err) {
		resp, err := h.Client.Do(req)
		if !checkErr(err) {
			defer resp.Body.Close()
			data, err := ioutil.ReadAll(resp.Body)
			if !checkErr(err) {
				err = json.Unmarshal(data, response)
				checkErr(err)
			}
		}
	}
	return
}

// Artist retrieves artist by ID
func (h *RapGenius) Artist(id int) (result *Artist, err error) {
	path := fmt.Sprintf("artists/%d", id)
	response := &artistResponseRoot{}
	err = h.execute(path, response)
	if !checkErr(err) {
		result = response.ArtistResponse.Artist
	}
	fmt.Print(response)
	return
}

// Search searches RapGenius for the specified query
func (h *RapGenius) Search(query string) (result []*SearchHit, err error) {
	path := fmt.Sprintf("search?q=%s", query)
	response := &searchResponse{}
	err = h.execute(path, response)
	if !checkErr(err) {
		result = response.Result.SearchHits
	}
	return
}
