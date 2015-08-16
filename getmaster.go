package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "os"
    )

func main() {
    response, err := http.Get("http://github.com/pnd-tech-club/kazoo-quest/archive/master.zip")
    if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
    } else {
        defer response.Body.Close()
        contents, err := ioutil.ReadAll(response.Body)
        if err != nil {
            fmt.Printf("%s", err)
            os.Exit(1)
        }
        //fmt.Printf("%s\n", string(contents))
	ioutil.WriteFile("kqmaster.zip", contents, 0777)
    }
}

