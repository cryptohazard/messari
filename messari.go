package messari

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jedib0t/go-pretty/table"
)

var (
	apiurl        = "https://data.messari.io/api/v1"
	apiurlmarkets = fmt.Sprintf("%v/markets", apiurl)
	apiurlassets  = fmt.Sprintf("%v/assets", apiurl)
)

// Market hold information about exchange and pair.
type Market struct {
	Exchange string `json:"exchange"`
	Base     string `json:"base"`
	Quote    string `json:"quote"`
	Pair     string `json:"pair"`
}

// Markets is the list of all exchanges and pairs.
type Markets []Market

func (m Markets) String() string {
	tablewriter.AppendHeader(table.Row{"EXCHANGE", "BASE", "QUOTE", "PAIR"})
	for _, v := range m {
		tablewriter.AppendRow([]interface{}{v.Exchange, v.Base, v.Quote, v.Pair})
	}
	return tablewriter.Render()
}

// Asset is a crypto asset.
type Asset struct {
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
	ID     string `json:"id"`
}

// Assets is the list of all crypto assets.
type Assets []Asset

func (a Assets) String() string {
	tablewriter.AppendHeader(table.Row{"Symbol", "Name", "ID"})
	for _, v := range a {
		tablewriter.AppendRow([]interface{}{v.Symbol, v.Name, v.ID})
	}
	return tablewriter.Render()
}

// Client is an HTTP client.
type Client struct {
	http.Client
}

// Markets returns the list of all exchanges and pairs.
func (c *Client) Markets() (Markets, error) {
	var markets struct {
		Data []Market `json:"data"`
	}

	response, err := c.Get(apiurlmarkets)
	if err != nil {
		return Markets{}, fmt.Errorf("messari api: %v", err)
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case 200:
		break
	case 400, 401, 403, 429, 500:
		var e map[string]string
		err = json.NewDecoder(response.Body).Decode(&e)
		if err != nil {
			return Markets{}, fmt.Errorf("messari api: failed to parse error body: %v", err)
		}
		return Markets{}, fmt.Errorf("messari api: %v", e["error_message"])
	default:
		return Markets{}, fmt.Errorf("messari api: returns unknown error")
	}

	err = json.NewDecoder(response.Body).Decode(&markets)
	if err != nil {
		return Markets{}, fmt.Errorf("messari api: failed to parse body: %v", err)
	}

	return markets.Data, nil
}

// Assets returns the list of all crypto assets.
func (c *Client) Assets() (Assets, error) {
	return c.assets("")
}

// AssetsWithMetrics returns the list of all crypto assets.
func (c *Client) AssetsWithMetrics() (Assets, error) {
	return c.assets("?with-metrics")
}

// AssetsWithProfiles returns the list of all crypto assets.
func (c *Client) AssetsWithProfiles() (Assets, error) {
	return c.assets("?with-profiles")
}

// AssetsWithMetricsAndProfiles returns the list of all crypto assets.
func (c *Client) AssetsWithMetricsAndProfiles() (Assets, error) {
	return c.assets("?with-metrics&with-profiles")
}

func (c *Client) assets(queryparam string) (Assets, error) {
	var assets struct {
		Data []Asset `json:"data"`
	}

	response, err := c.Get(fmt.Sprintf("%v%v", apiurlassets, queryparam))
	if err != nil {
		return Assets{}, fmt.Errorf("messari api: %v", err)
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case 200:
		break
	case 400, 401, 403, 429, 500:
		var e map[string]string
		err = json.NewDecoder(response.Body).Decode(&e)
		if err != nil {
			return Assets{}, fmt.Errorf("messari api: failed to parse error body: %v", err)
		}
		return Assets{}, fmt.Errorf("messari api: %v", e["error_message"])
	default:
		return Assets{}, fmt.Errorf("messari api: returns unknown error")
	}

	err = json.NewDecoder(response.Body).Decode(&assets)
	if err != nil {
		return Assets{}, fmt.Errorf("messari api: failed to parse body: %v", err)
	}

	return assets.Data, nil
}
