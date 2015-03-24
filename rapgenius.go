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

// Song retrieves song by ID
func (h *RapGenius) Song(id int) (result *Song, err error) {
	path := fmt.Sprintf("songs/%d", id)
	response := &songResponse{}
	err = h.execute(path, response)
	if !checkErr(err) {
		result = response.Response.Song
	}
	fmt.Print(response)
	return
}

// Artist retrieves artist by ID
func (h *RapGenius) Artist(id int) (result *Artist, err error) {
	path := fmt.Sprintf("artists/%d", id)
	response := &artistResponse{}
	err = h.execute(path, response)
	if !checkErr(err) {
		result = response.Response.Artist
	}
	fmt.Print(response)
	return
}

// Search searches RapGenius for the specified query
func (h *RapGenius) Search(query string) (result []*SearchItem, err error) {
	path := fmt.Sprintf("search?q=%s", query)
	response := &searchResponse{}
	err = h.execute(path, response)
	if !checkErr(err) {
		result = make([]*SearchItem, len(response.Response.SearchHits))
		for i, item := range response.Response.SearchHits {
			result[i] = item.Item
		}
	}
	return
}
