package analytics

import (
  "fmt"
  "strings"
  "math/rand"
  "net/url"
  "net/http"
)

type UA struct {
  ID string
}

func CreateUA(id string) *UA {
  return &UA{ID: id,}
}

// ver que onda con el CID

func (ua *UA)Event(category string, action string, label string, value int, uuid string) {

  hc := http.Client{}
  
  form := url.Values{}
  form.Add("v", "1")
  form.Add("tid", ua.ID)
  form.Add("cid", uuid)
  form.Add("t", "event")
  form.Add("ec", category)
  form.Add("ea", action)
  form.Add("el", label)
  form.Add("ev", fmt.Sprint(value))

  req, _ := http.NewRequest("POST", "https://www.google-analytics.com/collect", strings.NewReader(form.Encode()))
  req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
  hc.Do(req)
}

func UUID() string {
  b := make([]byte, 16)
  _, err := rand.Read(b)
  if err != nil {
    return ""
  }
  uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
  return uuid
}