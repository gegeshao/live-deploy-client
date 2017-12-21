package utils

import (
  "io/ioutil"
)

func GetScript(filename string)(string, error){
  fileContet, err:=ioutil.ReadFile("scripts/"+filename+".lua")
  if err!=nil{
    return "", nil
  }
  return string(fileContet), nil
}
