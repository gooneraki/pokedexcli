package main

import (
	"fmt"
	"io"
	"net/http"
)

func fetchFromUrl(locationAreaUrl string, c *config) ([]byte, error) {

	entry, found := c.cache.Get(locationAreaUrl)

	if found {
		return entry, nil
	} else {

		res, err := http.Get(locationAreaUrl)
		if err != nil {
			return nil, fmt.Errorf("error fetching %v", err)
		}

		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read response body: %v", err)
		}

		c.cache.Add(locationAreaUrl, body)
		return body, nil
	}
}
