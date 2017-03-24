package rest

import (
  "strings"
)

/* private */

func split(pattern, url string) ([]string, []string) {
  patternZones := strings.Split(pattern, "/")[1:]
  urlZones := strings.Split(url, "/")[1:]

  return patternZones, urlZones
}

/* public */

func Match(pattern, url string) bool {
  match := true

  patternZones, urlZones := split(pattern, url)

  if len(patternZones) != len(urlZones) {
    match = false
  }

  for i := 0; i < len(urlZones); i++ {
    hasPrefix := strings.HasPrefix(patternZones[i], ":")
    if patternZones[i] != urlZones[i] && !hasPrefix {
      match = false
    }
  }

  return match
}

func GetParams(pattern, url string) map[string]string {
  params := make(map[string]string)

  patternZones, urlZones := split(pattern, url)

  for i, zone := range patternZones {
    if strings.HasPrefix(zone, ":") {
      params[zone[1:]] = urlZones[i]
    }
  }

  return params
}
