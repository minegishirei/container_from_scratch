// main 関数はコマンドライン引数を解析し、適切な関数を実行します。
// run 関数は現在の実行可能ファイルを使用して特定の名前空間設定で子プロセスを作成します。
// child 関数はホスト名を設定し、指定されたコマンドを別のプロセスで実行します。
package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// main 関数はコマンドライン引数を解析し、適切な関数を実行します。
func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("無効なコマンドです！！")
	}
}

// run 関数は現在の実行可能ファイルを使用して特定の名前空間設定で子プロセスを作成します。
func run() {
	fmt.Printf("PID %d として %v を実行中\n", os.Getpid(), os.Args[2:])

	// 子プロセス用の引数を準備します。
	args := append([]string{"child"}, os.Args[2:]...)
	cmd := exec.Command("/proc/self/exe", args...)

	// 子プロセスの標準入力、出力、エラーを設定します。
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// 子プロセスが新しいUTS名前空間を使用するように設定します。
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS,
	}

	// 子プロセスを実行します。
	cmd.Run()
}

// child 関数はホスト名を設定し、指定されたコマンドを別のプロセスで実行します。
func child() {
	fmt.Printf("PID %d として %v を実行中\n", os.Getpid(), os.Args[2:])

	// コンテナのホスト名を設定します。
	syscall.Sethostname([]byte("container-demo"))

	// 別のプロセスで指定されたコマンドを実行します。
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}