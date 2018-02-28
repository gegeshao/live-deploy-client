package utils

import (
  "testing"
)


func TestGetConfig(t *testing.T) {
  config:=GetConfig()
  if config.LuaScriptsDir == ""{
    t.Fail()
  }
}