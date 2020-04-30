// client information
package nap

import (
  "net/http"
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
func (c *Client) ProcessRequest(baseURL string, res *RestResource, params map[string]string) error {
  // URL formation and trimming
  endpoint := strings.TrimLeft(res.RenderEndpoint(params), "/")
  trimmedBaseURL := strings.TrimRight(baseURL, "/")
  url := trimmedBaseURL + "/" + endpoint
  // processing request
  req, err := http.NewRequest(res.Method, url, nil)
  if err != nil {
    return err
  }

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
  if err := res.Router.CallFunc(resp, nil); err != nil {
    return err
  }
  return nil
}
