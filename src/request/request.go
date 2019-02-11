package request

import (
    "errors"
    "fmt"
    "net/http"
    "strings"
)

func makeHeader(req *http.Request, key string, value string){
    req.Header.Set(key, value)
    return
}

func initHeader(req *http.Request){
    makeHeader(req, "User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.53 Safari/537.36")
    return
}

func HttpDo(method, url, data string, headers map[string]string)(*http.Response, error) {
    methodList := []string{"GET", "POST"}
    method = strings.ToUpper(method)
    flag := false
    for _, value := range methodList{
        if value == method{
            flag = true
        }
    }
    if flag == false{
        return nil, errors.New("Invalide method")
    }
    client := &http.Client{}
    req, err := http.NewRequest(method, url, strings.NewReader(data))
    if err != nil {
        fmt.Println(err)
        return nil, err
    }
    initHeader(req)
    for key, value := range headers {
        makeHeader(req, key, value)
    }
    resp, err := client.Do(req)
    if err != nil{
        return nil, err
    }
    //defer resp.Body.Close()
    return resp, nil
}