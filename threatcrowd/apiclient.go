package threatcrowd

import (
	"encoding/json"
	"log"
	"net/http"
)

type ThreatCrowdClient struct {
	Client  *http.Client
	BaseURL string
	Delay   int64
}

func NewClient() *ThreatCrowdClient {
	BaseURL := "https://www.threatcrowd.org/searchApi/v2"
	Delay := int64(10)

	client := &ThreatCrowdClient{
		Client:  &http.Client{},
		BaseURL: BaseURL,
		Delay:   Delay,
	}

	return client
}

func (c *ThreatCrowdClient) IPReport(ip string) IP {

	req, err := http.NewRequest(http.MethodGet, c.BaseURL+"/ip/report/", nil)

	if err != nil {
		log.Fatalln("error building query")
	}

	q := req.URL.Query()
	q.Add("ip", ip)
	req.URL.RawQuery = q.Encode()

	resp, err := c.Client.Do(req)

	if err != nil {
		log.Fatalf("error querying %s", req.URL.String())
	}
	defer resp.Body.Close()

	ipReport := IP{}
	err = json.NewDecoder(resp.Body).Decode(&ipReport)

	if err != nil {
		log.Fatalln("error reading response")
	}

	return ipReport
}
