package main

import (
	// "errors"
  "encoding/json"
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
)

func main() {
  // url := "https://a.4cdn.org/g/catalog.json"
  url := "http://127.0.0.1:8080/catalog.json"
  res, err := http.Get(url)

  if err != nil {
    log.Fatal(err)
  }

  body, rerr := ioutil.ReadAll(res.Body)
  bodyBytes := []byte(body)

  if rerr != nil {
    log.Fatal(rerr)
    return
  }

  if bodyBytes == nil {
    log.Fatal("Unexpected: bodyBytes is nil")
    return
  }

  var catalog []map[string]interface{}
  jerr := json.Unmarshal(bodyBytes, &catalog)

  if jerr != nil {
    log.Fatal(jerr)
    return
  }

  for _, page := range catalog {
    fmt.Println(page["page"])
  }

}
