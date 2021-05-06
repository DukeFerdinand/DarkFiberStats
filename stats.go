package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	tm "github.com/buger/goterm"
)

type PayoutStructure struct {
	Error    string `json:"error,omitempty"`
	Immature string `json:"immature"`
	Owed     string `json:"owed"`
	Paid     string `json:"paid"`
}

type StatStructure struct {
	Error           string  `json:"error"`
	Shares          float64 `json:"shares"`
	TotalShares     float64 `json:"totalShares"`
	UserPercentage  string  `json:"userPercentage"`
	EstimatedPayout float64 `json:"estimatedPayout"`
}

type MinerStatConfig struct {
	MinerId     string `json:"minerId"`
	CoinType    string `json:"coinType"`
	BlockReward int    `json:"blockReward"`
}

func (c MinerStatConfig) GetPayouts() (PayoutStructure, error) {
	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(c)

	res, err := http.Post("https://aggrogator.dev/api/coins/payouts", "application/text", payload)

	if err != nil {
		return PayoutStructure{}, err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	data, err := ioutil.ReadAll(res.Body)

	if err != nil || len(data) == 0 {
		return PayoutStructure{}, err
	}

	payouts := PayoutStructure{}
	parse_error := json.Unmarshal(data, &payouts)

	if parse_error != nil {
		return PayoutStructure{}, err
	}

	return payouts, nil
}

func (c MinerStatConfig) GetStats() (StatStructure, error) {
	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(c)

	res, err := http.Post("https://aggrogator.dev/api/coins/stats", "application/text", payload)

	if err != nil {
		return StatStructure{}, err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	data, err := ioutil.ReadAll(res.Body)

	if err != nil || len(data) == 0 {
		return StatStructure{}, err
	}

	stats := StatStructure{}
	parse_error := json.Unmarshal(data, &stats)

	if parse_error != nil {
		return StatStructure{}, nil
	}

	return stats, nil
}

func (c MinerStatConfig) PrintConfig() {

	tm.Println(strings.Join([]string{"Monitoring stats for: ", tm.Color(c.MinerId, tm.CYAN)}, ""))
	tm.Println(strings.Join([]string{"Coin Type           : ", tm.Color(c.CoinType, tm.CYAN)}, ""))

}
