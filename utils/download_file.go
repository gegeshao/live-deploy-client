package utils

import (
  "io"
  "net/http"
  "os"
  "strings"
)

func Download(url string, filename string) error{
  if filename == "" {
    pathArr := strings.Split(url, "/")
    filename = pathArr[len(pathArr)-1]
  }
  response, e := http.Get(url)
  if e != nil {
    return e
  }

  defer response.Body.Close()
  file, err := os.Create(filename)
  if err != nil {
    return err
  }
  _, err = io.Copy(file, response.Body)
  if err != nil {
    return err
  }
  file.Close()
  return nil
}
