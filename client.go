package picsmize

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type Client struct {
	*Picsmize

	compress  map[string]interface{}
	filter    map[string]interface{}
	resize    map[string]interface{}
	scale     map[string]interface{}
	crop      map[string]interface{}
	flip      map[string]interface{}
	watermark map[string]interface{}
	toJSON    bool
}

/**
 * Compress an image
 *
 * @param {Array} data
 */

func (c *Client) Compress(compressLevel Options) *Client {
	c.compress = compressLevel
	return c
}

/*
 * Resizes an image
 *
 * @param {String} Mode Type, {Array} Options
 */

func (c *Client) Resize(modeType string, resizeOptions Options) *Client {

	resizeOptions["mode"] = modeType
	c.resize = resizeOptions
	return c
}

/*
 * Scales an image
 *
 * @param {Number} Scale Size
 */

func (c *Client) Scale(scaleSize float64) *Client {

	scaleJson := map[string]interface{}{}
	scaleJson["size"] = scaleSize
	c.scale = scaleJson
	return c
}

/*
 * Crop an image
 *
 * @param {String} Mode Type, {Array} Options
 */

func (c *Client) Crop(modeType string, cropOptions Options) *Client {

	cropOptions["mode"] = modeType
	c.crop = cropOptions
	return c
}

/*
 * Flip an image
 *
 * @param {String} Type
 */

func (c *Client) Flip(flipType string) *Client {

	flipJson := map[string]interface{}{}
	flipJson[flipType] = true
	c.flip = flipJson
	return c
}

/*
 * Filter an image
 *
 * @param {String} filter type, {Array} options
 */

func (c *Client) Filter(filterType string, filterOptions Options) *Client {

	filterJson := map[string]interface{}{}
	filterJson[filterType] = filterOptions
	c.filter = filterJson
	return c
}

/*
 * Watermark an image
 *
 * @param {String} filter type, {Array} options
 */

func (c *Client) Watermark(watermarkOptions Options) *Client {

	c.watermark = watermarkOptions
	return c
}

/*
 * Sends a standard request to the API
 * and returns a JSON response
 */

func (c *Client) ToJSON() ([]interface{}, error) {

	c.toJSON = true
	if !c.inputFetch {
		c.errorMessage = "fetch(string) method requires a valid file URL passed as an argument"
	}

	res, err := c.request()
	if err != nil {
		return nil, err
	}

	return res, nil
}

/*
 * Sent a request form Picsmize
 *
 * @param {Array} response
 */

func (c *Client) request() ([]interface{}, error) {

	if c.errorMessage != "" {
		return nil, errors.New(c.errorMessage)
	}

	if c.apiKey == "" {
		return nil, errors.New("requires a valid API key for image processing")
	}

	process := c.generateProcess()
	options := map[string]interface{}{}
	if c.inputFetch {
		options["img_url"] = c.imgUrl
	}

	options["process"] = process
	opjn, err := json.Marshal(options)
	if err != nil {
		return nil, err
	}

	payload := bytes.NewReader(opjn)
	req, _ := http.NewRequest("POST", apiEndpoint+"/image/process", payload)

	req.Header.Add("content-type", "application/json")
	req.Header.Add("apikey", c.apiKey)

	res, resErr := http.DefaultClient.Do(req)
	if resErr != nil {
		return nil, resErr
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	textBytes := []byte(string(body))

	var bodyRes interface{}
	errv := json.Unmarshal(textBytes, &bodyRes)
	if err != nil {
		return nil, errv
	}
	result := bodyRes.(map[string]interface{})

	if !result["status"].(bool) {
		return nil, errors.New(result["message"].(string))
	}

	apiCall := make(map[string]string)
	apiCall["Limit"] = res.Header.Get("x-ratelimit-limit")
	apiCall["Remaining"] = res.Header.Get("x-ratelimit-remaining")
	response := append([]interface{}{result}, apiCall)

	return response, nil
}

func (c *Client) generateProcess() map[string]interface{} {
	process := map[string]interface{}{}

	if c.compress != nil {
		process["compress"] = c.compress
	}

	if c.resize != nil {
		process["resize"] = c.resize
	}

	if c.scale != nil {
		process["scale"] = c.scale
	}

	if c.crop != nil {
		process["crop"] = c.crop
	}

	if c.flip != nil {
		process["flip"] = c.flip
	}

	if c.filter != nil {
		process["filter"] = c.filter
	}

	if c.watermark != nil {
		process["watermark"] = c.watermark
	}
	return process
}
