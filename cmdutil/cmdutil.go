package cmdutil

import (
    "bufio"
    "fmt"
    "os"
    "os/exec"
    "strings"
)

var (
    InputFile   *os.File = os.Stdin
    inputBuffer *bufio.Reader
)

func ReadLine() string {
    buf := buffer()
    line, err := buf.ReadString('\n')
    if err != nil {
        fmt.Println(err)
    }
    return strings.TrimSpace(string(line))
}

func Silence() {
    runCommand(exec.Command("stty", "-echo"))
}

func Unsilence() {
    runCommand(exec.Command("stty", "echo"))
}

func runCommand(command *exec.Cmd) {
    command.Stdin = os.Stdin
    command.Stdout = os.Stdout
    command.Run()
}

func buffer() *bufio.Reader {
    if inputBuffer == nil {
        inputBuffer = bufio.NewReader(InputFile)
    }
    return inputBuffer
}
