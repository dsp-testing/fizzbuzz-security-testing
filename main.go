package main

import (
  "encoding/json"
  "fmt"
  "net/http"
)

type Repository struct {
  Name        string `json:"name"`
  Description string `json:"description"`
  Language    string `json:"language"`
}

func main() {
  req, err := http.NewRequest("GET", fmt.Sprintf("https://api.github.com/users/%s/repos", "monalisa"), nil)
  if err != nil {
    fmt.Println("Failed to create request:", err)
    return
  }

  req.Header.Add("Authorization", "token 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")

  resp, err := http.DefaultClient.Do(req)
  if err != nil {
    fmt.Println("Failed to fetch repositories:", err)
    return
  }
  defer resp.Body.Close()

  var repositories []Repository
  err = json.NewDecoder(resp.Body).Decode(&repositories)
  if err != nil {
    fmt.Println("Failed to parse response body:", err)
    return
  }

  for _, repo := range repositories {
    if repo.Description != "" {
      fmt.Printf("Name: %s\nDescription: %s\nLanguage: %s\n\n", repo.Name, repo.Description, repo.Language)
    } else {
      fmt.Printf("Name: %s\nLanguage: %s\n\n", repo.Name, repo.Language)
    }
  }
}
