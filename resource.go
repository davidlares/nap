// resources
package nap

import (
  "log"
  "io/ioutil"
  "html/template"
  "bytes"
)


type RestResource struct {
  Endpoint string // URI
  Method string // HTTP verb
  Router *CBRouter // for dynamic resources (re-usage instead of recreation)
}

func NewResource(endpoint, method string, router *CBRouter) *RestResource {
  return &RestResource {
    Endpoint: endpoint,
    Method: method,
    Router: router,
  }
}

// returning the endpoint string
func (r *RestResource) RenderEndpoint(params map[string]string) string {
  if params == nil {
    return r.Endpoint
  }
  t, err := template.New("resource").Parse(r.Endpoint)
  if err != nil {
    log.Fatalln("Unable to parse endpoint")
  }
  buffer :=  &bytes.Buffer{}
  t.Execute(buffer, params)
  endpoint, err := ioutil.ReadAll(buffer)
  if err != nil {
    log.Fatalln("Unable to read endpoint")
  }
  return string(endpoint)
}
