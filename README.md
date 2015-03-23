# A Go RapGenius client

## Install
    go get github.com/mondok/rapgenius

## Searching
    rapgenius := rapgenius.New()
    results, err := rapgenius.Search("notorious")
    fmt.Printf("%s, %s", results[0].Title, results[0].Artist.Name)

## Retrieving Artists
    rapgenius := rapgenius.New()
    artist, err := rapgenius.Artist(22)
    fmt.Printf("Artist is %s", artist.Name)

## Retrieving Songs
    rapgenius := rapgenius.New()
    song, err := rapgenius.Song(22)
    fmt.Printf("Song is %s", song.Title)
