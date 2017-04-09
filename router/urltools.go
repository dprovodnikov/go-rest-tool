package router

import (
  "strings"
)

func split(pattern, url string) ([]string, []string) {
  if strings.HasSuffix(url, "/") {
    url = string(url[:len(url)-1])
  }
  if strings.HasSuffix(pattern, "/") {
    pattern = string(pattern[:len(pattern)-1])
  }

  patternZones := strings.Split(pattern, "/")[1:]
  urlZones := strings.Split(url, "/")[1:]

  return patternZones, urlZones
}

func Match(pattern, url string) bool {
  patternZones, urlZones := split(pattern, url)
  if len(patternZones) != len(urlZones) {
    return false
  }
  for i := 0; i < len(urlZones); i++ {
    hasPrefix := strings.HasPrefix(patternZones[i], ":")
    if patternZones[i] != urlZones[i] && !hasPrefix {
      return false
    }
  }
  return true
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
