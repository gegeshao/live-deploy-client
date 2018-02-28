package utils

import (
  "io"
  "log"
  "net/http"
  "os"
  "regexp"
  "strings"
  "testing"
)

func TestDownload(t *testing.T){
  url := "https://api.github.com/repos/certbot/certbot/tarball/v0.21.1"
  response, e := http.Get(url)
  if e!=nil{
    t.Fail()
  }
  defer response.Body.Close()
  filename := ""
  desposition := response.Header.Get("Content-Disposition")
  reg:=regexp.MustCompile(`(?m)filename[^;=]*=['"]?([^;'"]+)['"]?`)
  regResult := reg.FindStringSubmatch(desposition)
  if len(regResult) == 0{
    pathArr := strings.Split(url, "/")
    filename = pathArr[len(pathArr)-1]
  }else{
    filename = regResult[1]
  }
  //config:= GetConfig()
  //installDir := config.System.InstallPath
  switch response.Header.Get("Content-Type"){
  case  "application/x-gzip":
  default:
    //filename = path.Join(installDir, filename)
    log.Println(filename)
    file, err := os.Create(filename)
    if err != nil {
      log.Println(err)
      t.Fail()
    }
    if _, err := io.Copy(file, response.Body); err!=nil{
      log.Println(err)
      t.Fail()
    }
    file.Close()
  }


}