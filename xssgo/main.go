package main

import "fmt"
import "cli"

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
    conf.input = parsed
    fmt.Println(conf.input.url)
    return
}
