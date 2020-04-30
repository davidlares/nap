// working code for APIs
package nap

import (
  "fmt"
)

type API struct {
  BaseURL       string
  Resources     map[string]*RestResource // names for the resources (queries)
  DefaultRouter *CBRouter // pointer to the CBRouter (dynamic, if a resource have it)
  Client        *Client // Client library (wrapper HTTP client)
}

// implementing a new API
func NewAPI(baseURL string) *API {
  return &API {
    BaseURL:        baseURL,
    Resources:      make(map[string]*RestResource),
    DefaultRouter:  NewRouter(),
    Client:         NewClient(),
  }
}

// setting up to the client (if needed)
func (api *API) SetAuth(auth Authentication) {
  api.Client.SetAuth(auth)
}

// adding resources
func (api *API) AddResource(name string, res *RestResource) {
  api.Resources[name] = res
}

// adding API calls
func (api *API) Call(name string, params map[string]string) error {
  res, ok := api.Resources[name]
  if !ok {
    return fmt.Errorf("Resource does not exist %s", name)
  }
  if err := api.Client.ProcessRequest(api.BaseURL, res, params); err != nil {
    return err
  }
  return nil
}

func (api *API) ResourceNames() []string {
  resources := []string{}
  for k := range api.Resources {
    resources = append(resources, k)
  }
  return resources
}
