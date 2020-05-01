package main

import (
  "flag"
  "io/ioutil"
  "fmt"
  "log"
  "net/http"
  "os"
  "github.com/davidlares/nap"
)

var api = nap.NewAPI("https://httpbin.org")

func main() {
  list := flag.Bool("list", false, "Get list of all API resources")
  flag.Parse()
  if *list {
    for _, name := range api.ResourceNames() {
      fmt.Println(name)
    }
    return
  }
  resource := os.Args[1]
  if err := api.Call(resource, nil, nil); err != nil {
    log.Fatalln(err)
  }
}

func init() {
  router := nap.NewRouter()
  // interface removed
  router.RegisterFunc(200, func(resp *http.Response) error {
    defer resp.Body.Close()
    content, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      return err
    }
    fmt.Println(string(content))
    return nil
  })

  api.AddResource("get", nap.NewResource("/get", "GET", router))
}
