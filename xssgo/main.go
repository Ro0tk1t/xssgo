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
    input map
    url string
    dict string
    spider bool
    results *Result
}

pased := cli.Cli()
conf := Conf()
conf.input = pased

func main(){
    return
}
