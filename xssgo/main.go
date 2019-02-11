package main

import (
    "cli"
    "data/datatype"
    "fmt"
    "io/ioutil"
    "net/url"
    "request"
)

var Conf datatype.Configure

func main(){
    parsed := cli.Cli()
    Conf.Input = *parsed
    u, err := url.Parse(parsed.Url)
    if err != nil{
        fmt.Println("[-]  Invalid URL !")
        return
    }
    Conf.Scheme = u.Scheme
    Conf.Host = u.Host
    Conf.Path = u.Path
    Conf.Username = u.User.Username()
    Conf.Password, _ = u.User.Password()
    Conf.Input.Url = "http://localhost/xss.php?x=<script>alert(1)</script>"
    resp, _ := request.HttpDo("GET", Conf.Input.Url, "", map[string]string{})
    //doc, err := goquery.NewDocumentFromReader(body)
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil{
        fmt.Println(err)
    }
    fmt.Println(string(body))
    return
}
