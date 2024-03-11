package main
import (
  "fmt"
  "os"
  "os/exec"
  "syscall"
)
// go run container.go run <cmd> <args>
// docker run <cmd> <args>
func main() {
  switch os.Args[1] {
    case "run":
      run()
    default:
      panic("invalid command!!")
  }
}
func run() {
  fmt.Printf("Running %v as PID %d \n", os.Args[2:], os.Getpid())
  cmd := exec.Command(os.Args[2], os.Args[3:]...)
  cmd.Stdin = os.Stdin
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  cmd.SysProcAttr = &syscall.SysProcAttr{
    Cloneflags: syscall.CLONE_NEWUTS
  }
  cmd.Run()
}