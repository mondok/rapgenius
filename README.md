# A Go RapGenius client

## Install
    go get github.com/mondok/rapgenius

### Searching
    rapgenius := rapgenius.New()
    results, err := rapgenius.Search("notorious")
    fmt.Printf("%s, %s", results[0].Item.Title, results[0].Item.Artist.Name)
