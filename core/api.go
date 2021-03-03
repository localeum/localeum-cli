package core

import (
	"github.com/go-resty/resty/v2"
)

const HostUrl = "https://api.localeum.com/api"

type Api struct {
	client *resty.Client
}

func NewApiClient() *Api {
	a := &Api{}
	a.client = resty.New()
	a.client.SetHostURL(HostUrl)

	a.client.SetHeaders(map[string]string{
		"Content-Type": "application/json",
		"User-Agent": "Localeum CLI",
	})

	return a
}

func (a *Api) SetAuthToken(token string) {
	a.client.SetAuthToken(token)
}

func (a *Api) Get(path string, query map[string]string) (*resty.Response, error) {
	return a.client.R().
		SetQueryParams(query).
		Get(path)
}