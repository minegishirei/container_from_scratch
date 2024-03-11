package main
import (
  "fmt"
  "os"
  "os/exec"
  "syscall"
)

// コマンドのオプションによって実行内容を変更。
func main(){
  switch os.Args[1] {
    case "run":
      run()
    case "child":
      child()
    default:
      panic("invalid command")
  }
}

func run(){
  // os.GetPIDはプロセスIDを取得（現在のプロセスIDと同じ）
  fmt.Printf("Runnning %v as PID %d \n", os.Args[2:], os.Getpid())
  args := append([]string{"child"}, os.Args[2:]...)
  cmd := exec.Command("/proc/self/exe", args...)
  cmd.Stdin = os.Stdin
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  cmd.SysProcAttr = &syscall.SysProcAttr{
    Cloneflags : syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
  }
  cmd.Run()
}
func child(){
  fmt.Printf("Running %v as PID %d \n", os.Args[2:], os.Getpid())
  syscall.Sethostname([]byte("container-demo"))
  cmd := exec.Command(os.Args[2], os.Args[3:]...)
  cmd.Stdin = os.Stdin
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  syscall.Chroot("/containerfs")
  os.Chdir("/")
  //syscall.Mount("proc", "proc", "proc", 0, "")
  cmd.Run()
}


