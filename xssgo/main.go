package main

import (
    "cli"
    "fmt"
    "net/url"
    "request"
)

type Result struct{
    Page string
    Param string
    Payload string
    Weakness bool
}

type Conf struct{
    Input cli.Parse
    Scheme string
    Host string
    Path string
    Username string
    Password string
    Dict string
    Results *Result
}

func main(){
    parsed := cli.Cli()
    conf := Conf{}
    conf.Input = *parsed
    u, err := url.Parse(parsed.Url)
    if err != nil{
        fmt.Println("[-]  Invalid URL !")
        return
    }
    conf.Scheme = u.Scheme
    conf.Host = u.Host
    conf.Path = u.Path
    conf.Username = u.User.Username()
    conf.Password, _ = u.User.Password()
    fmt.Println(u)
    body, _ := request.HttpDo("GET", conf.Input.Url, "", map[string]string{})
    fmt.Println(body)
    return
}
