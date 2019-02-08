package main

import (
    "cli"
    "fmt"
    "request"
)

type Result struct{
    page string
    param string
    payload string
    weakness bool
}

type Conf struct{
    input cli.Parse
    url string
    dict string
    spider bool
    results *Result
}

func main(){
    parsed := cli.Cli()
    conf := Conf{}
    conf.input = *parsed
    body, _ := request.HttpDo("GET", conf.input.Url, "", map[string]string{})
    fmt.Println(body)
    return
}
