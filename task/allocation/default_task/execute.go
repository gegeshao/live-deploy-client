package default_task

import (
  "bytes"
  "os/exec"
)

func execute(cwd string, command string, params ...string)(success bool, result string){
  cmd := exec.Command(command, params...)
  cmd.Dir = cwd
  var errBuf bytes.Buffer
  var outBuf bytes.Buffer
  cmd.Stdout = &outBuf
  cmd.Stderr = &errBuf
  startErr := cmd.Start()
  waitErr := cmd.Wait()
  err := errBuf.String()
  out := outBuf.String()
  if err != ""{ result = result + err}
  if out != ""{ result = result + out}
  if startErr != nil || waitErr != nil {
    if result == ""{
      result = startErr.Error() + "\n" + waitErr.Error()
    }
    return  false, result
  }
  return  true, result
}
