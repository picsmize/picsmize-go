package picsmize

import (
	"errors"
	"net/url"
)

var Version string = "1.0"

const apiEndpoint = "https://api.picsmize.com"

type Options map[string]interface{}

type Picsmize struct {
	apiKey string

	errorMessage string
	inputFetch   bool
	imgUrl       string
}

func Init(key string) (*Picsmize, error) {

	if key == "" {
		return nil, errors.New("requires a valid API key for image processing")
	}

	return &Picsmize{
		apiKey: key,
	}, nil
}

/*
 * Provides a URL of the image for processing
 *
 * @param {String} url
 */

func (p *Picsmize) Fetch(imgUrl string) *Client {

	_, err := url.ParseRequestURI(imgUrl)
	if err != nil && p.errorMessage == "" {
		p.errorMessage = "fetch(string) method requires a valid file URL passed as an argument"
	}

	p.inputFetch = true
	p.imgUrl = imgUrl
	return &Client{
		Picsmize: p,
		compress: map[string]interface{}{},
		toJSON:   false,
	}
}
