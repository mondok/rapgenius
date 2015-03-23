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
	Response struct {
		SearchHits []*SearchHit `json:"hits"`
	} `json:"response"`
}

type artistResponse struct {
	Response struct {
		Artist *Artist `json:"artist"`
	} `json:"response"`
}

type songResponse struct {
	Response struct {
		Song *Song `json:"song"`
	} `json:"response"`
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

// Song is a song result
type Song struct {
	AnnotationCount float64 `json:"annotation_count"`
	APIPath         string  `json:"api_path"`
	BopURL          string  `json:"bop_url"`

	DescriptionAnnotation struct {
		Annotatable struct {
			ID    float64 `json:"id"`
			Title string  `json:"title"`
			Type  string  `json:"type"`
			URL   string  `json:"url"`
		} `json:"annotatable"`
		Annotations []struct {
			Authors []struct {
				Attribution float64 `json:"attribution"`
				User        struct {
					Avatar struct {
						Thumb struct {
							BoundingBox struct {
								Height float64 `json:"height"`
								Width  float64 `json:"width"`
							} `json:"bounding_box"`
							URL string `json:"url"`
						} `json:"thumb"`
						Tiny struct {
							BoundingBox struct {
								Height float64 `json:"height"`
								Width  float64 `json:"width"`
							} `json:"bounding_box"`
							URL string `json:"url"`
						} `json:"tiny"`
					} `json:"avatar"`
					ID             float64     `json:"id"`
					Iq             float64     `json:"iq"`
					Name           string      `json:"name"`
					RoleForDisplay interface{} `json:"role_for_display"`
				} `json:"user"`
			} `json:"authors"`
			Body struct {
				Dom struct {
					Children []interface{} `json:"children"`
					Tag      string        `json:"tag"`
				} `json:"dom"`
			} `json:"body"`
			CosignedBy          []interface{} `json:"cosigned_by"`
			CurrentUserMetadata struct {
				Interactions struct {
					Cosign bool        `json:"cosign"`
					Vote   interface{} `json:"vote"`
				} `json:"interactions"`
				Permissions []interface{} `json:"permissions"`
			} `json:"current_user_metadata"`
			ID         float64     `json:"id"`
			Pinned     bool        `json:"pinned"`
			ShareURL   string      `json:"share_url"`
			State      string      `json:"state"`
			URL        string      `json:"url"`
			VerifiedBy interface{} `json:"verified_by"`
			VotesTotal float64     `json:"votes_total"`
		} `json:"annotations"`
		APIPath        string      `json:"api_path"`
		Classification string      `json:"classification"`
		EmbedURL       string      `json:"embed_url"`
		Featured       bool        `json:"featured"`
		Fragment       string      `json:"fragment"`
		ID             float64     `json:"id"`
		Path           string      `json:"path"`
		Range          interface{} `json:"range"`
		SongID         float64     `json:"song_id"`
		TrackingPaths  struct {
			Aggregate  string `json:"aggregate"`
			Concurrent string `json:"concurrent"`
		} `json:"tracking_paths"`
		TwitterShareMessage string `json:"twitter_share_message"`
		URL                 string `json:"url"`
	} `json:"description_annotation"`
	FeaturedArtists []interface{} `json:"featured_artists"`
	ID              float64       `json:"id"`
	Lyrics          struct {
		Dom struct {
			Children []struct {
				Children []interface{} `json:"children"`
				Tag      string        `json:"tag"`
			} `json:"children"`
			Tag string `json:"tag"`
		} `json:"dom"`
	} `json:"lyrics"`
	LyricsUpdatedAt float64 `json:"lyrics_updated_at"`
	Media           []struct {
		Provider string `json:"provider"`
		Type     string `json:"type"`
		URL      string `json:"url"`
	} `json:"media"`
	PrimaryArtist struct {
		ID       float64 `json:"id"`
		ImageURL string  `json:"image_url"`
		Name     string  `json:"name"`
		URL      string  `json:"url"`
	} `json:"primary_artist"`
	ProducerArtists []struct {
		ID       float64 `json:"id"`
		ImageURL string  `json:"image_url"`
		Name     string  `json:"name"`
		URL      string  `json:"url"`
	} `json:"producer_artists"`
	Stats struct {
		Hot       bool    `json:"hot"`
		Pageviews float64 `json:"pageviews"`
	} `json:"stats"`
	Title         string `json:"title"`
	TrackingPaths struct {
		Aggregate  string `json:"aggregate"`
		Concurrent string `json:"concurrent"`
	} `json:"tracking_paths"`
	URL                   string        `json:"url"`
	VerifiedAnnotationsBy []interface{} `json:"verified_annotations_by"`
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
func (h *RapGenius) Search(query string) (result []*SearchHit, err error) {
	path := fmt.Sprintf("search?q=%s", query)
	response := &searchResponse{}
	err = h.execute(path, response)
	if !checkErr(err) {
		result = response.Response.SearchHits
	}
	return
}
