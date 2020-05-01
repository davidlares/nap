// client information
package nap

import (
  "net/http"
  "bytes"
  "encoding/json"
  "strings"
)

type Client struct {
  Client *http.Client
  AuthInfo Authentication // interface
}

func NewClient() *Client {
  return &Client {
      Client: http.DefaultClient, // retrifve a http client (default)
  }
}

func (c *Client) SetAuth(auth Authentication) {
  c.AuthInfo = auth // assinging the interface
}

// processing request
func (c *Client) ProcessRequest(baseURL string, res *RestResource, params map[string]string, payload interface{}) error {
  // URL formation and trimming
  endpoint := strings.TrimLeft(res.RenderEndpoint(params), "/")
  trimmedBaseURL := strings.TrimRight(baseURL, "/")
  url := trimmedBaseURL + "/" + endpoint
  // processing requests
  req := buildClientRequest(res.Method, url, payload)
  // processing request
  // req, err := http.NewRequest(res.Method, url, nil)
  // if err != nil {
  //   return err
  // }

  // check Auth information
  if c.AuthInfo != nil {
    req.Header.Add("Authorization", c.AuthInfo.AuthorizationHeader())
  }

  // proecessing response
  resp, err := c.Client.Do(req)
  if err != nil {
    return err
  }
  // calling function
  // if err := res.Router.CallFunc(resp, nil); err != nil {
  //   return err
  // }
  // return nil
  return res.Router.CallFunc(resp)

}

func buildClientRequest(method, url string, payload interface{}) *http.Request {
  if payload != nil {
    payloadBytes, err := json.Marshal(payload)
    if err != nil {
      return nil
    }
    payloadBuffer := bytes.NewBuffer(payloadBytes)
    req, err := http.NewRequest(method, url, payloadBuffer)
    return req
  }
  req, err := http.NewRequest(method, url, nil)
  if err != nil {
    return nil
  }
  return req
}
