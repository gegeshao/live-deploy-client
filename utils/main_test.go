package utils

import (
  "os"
  "testing"
)

var CWD = "/Users/hyh/go-work/src/live-deploy-client"

func TestMain(m *testing.M) {
  SetCWD(CWD)
  InitConfig("./conf/config.yaml")
  // call flag.Parse() here if TestMain uses flags
  os.Exit(m.Run())
}
