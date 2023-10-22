package movies

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type OverseerrResponse struct {
    ExternalIds struct {
        ImdbId string `json:"imdbId"`
    } `json:"externalIds"`
}

func GetDetails(id string) Movie {
    host := os.Getenv("OVERSEERR_HOST")
    token := os.Getenv("OVERSEERR_TOKEN")

    fmt.Println("Getting details for:", id)

    url := fmt.Sprintf("%s/api/v1/movie/%s", host, id)

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        fmt.Println(err)
        fmt.Println("Failed to create request")
    }

    req.Header.Add("X-Api-Key", token)

    client := &http.Client{}

    response, err := client.Do(req)
    if err != nil {
        fmt.Println(err)
        fmt.Println("Failed to send request")
    }

    defer response.Body.Close()

    var details Movie
    err = json.NewDecoder(response.Body).Decode(&details)
    if err != nil {
        body, err := io.ReadAll(response.Body)
        if err != nil {
            fmt.Println(err)
            fmt.Println("Failed to read response body")
        }

        fmt.Println(string(body))

        fmt.Println(err)
        fmt.Println("Failed to decode overseerr details response")
    }

    return details
}
