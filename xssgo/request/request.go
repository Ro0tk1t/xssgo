package request

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "strings"
)

func make_header(req *http.Request, key string, value string)(req *http.Request){
    req.Header.Set(key, value)
    return
}

func init_header(req *http.Request)(req *http.Request){
    make_header(req, "User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.53 Safari/537.36")
    return
}

func HttpDo(method, url, data string, headers map){
    client := &http.Client{}
    req, err := http.NewRequest(method, url, strings.NewReader(data))
    if err != nil{
        return nil, err
    }
    init_header(req)
    for key, value := range headers{
        make_header(req, key, value)
    }
    resp, err := client.Do(req)
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil{
        return nil, err
    }
    return string(body)
}
