package main

import (
	"net/http"
)

type Client interface {
	Fetch(url string) *http.Response
}

type client struct{}

func NewClient() Client {
	return &client{}
}

func (c *client) Fetch(url string) *http.Response {
	return &http.Response{}
}
