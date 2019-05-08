package main

import (
  "fmt"
  "testing"

  "github.com/webability-go/analytics"
)

var UA *analytics.UA

func TestAnalytics(t *testing.T) {
  
  UA = analytics.CreateUA("UA-11441155-8")
  UA.SetUUID(analytics.UUID())
  UA.SetGeoID("FR")
  UA.Event("alexa-go", "test-unit", "country", 0)
  
  fmt.Println("OK")
}



