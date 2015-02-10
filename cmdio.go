/* http://stackoverflow.com/questions/14803425/properly-passing-data-on-stdin-to-a-command-and-receiving-data-from-stdout-of-th */
package main

import (
    "bytes"
    "io"
    "log"
    "os"
    "os/exec"
    "fmt"
)

func main() {
    usrinput := ""
    fmt.Scan(&usrinput) // Scan is insufficient, exits after a "\n"
    runCatFromStdinWorks(populateStdin(usrinput))
    runCatFromStdinWorks(populateStdin("bbb\n"))
}

func populateStdin(str string) func(io.WriteCloser) {
    return func(stdin io.WriteCloser) {
        defer stdin.Close()
        io.Copy(stdin, bytes.NewBufferString(str))
    }
}

func runCatFromStdinWorks(populate_stdin_func func(io.WriteCloser)) {
    cmd := exec.Command("cat")
    stdin, err := cmd.StdinPipe()
    if err != nil {
        log.Panic(err)
    }
    stdout, err := cmd.StdoutPipe()
    if err != nil {
        log.Panic(err)
    }
    err = cmd.Start()
    if err != nil {
        log.Panic(err)
    }
    populate_stdin_func(stdin)
    io.Copy(os.Stdout, stdout)
    err = cmd.Wait()
    if err != nil {
        log.Panic(err)
    }
}
