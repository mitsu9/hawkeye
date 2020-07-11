package toggl

import (
	"io/ioutil"
	"log"
	"net/http"
)

type Client struct {
	apiToken string
}

func NewClient(apiToken string) *Client {
	return &Client{apiToken: apiToken}
}

func (c Client) getRequest(url string, params map[string]string) []byte {
	req, _ := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(c.apiToken, "api_token")

	p := req.URL.Query()
	for k, v := range params {
		p.Add(k, v)
	}
	req.URL.RawQuery = p.Encode()

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)

	return bytes
}
