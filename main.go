package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	tm "github.com/buger/goterm"
	godotenv "github.com/joho/godotenv"
)

type MinerStatInput struct {
	MinerId string `json:"minerId"`
  CoinType string `json:"coinType"`
  BlockReward int `json:"blockReward"`
}

func print_payouts(i float64, o float64, p float64) {
		// Payouts
	tm.Println("\nPayments -> OK")
	p_box := tm.NewBox(75|tm.PCT, 5, 0)
	p_rows := strings.Join(
		[]string{
			strings.Join([]string{ "Immature: ", fmt.Sprint(i),  }, ""),
			strings.Join([]string{ "Owed    : ", fmt.Sprint(o),  }, ""),
			strings.Join([]string{ "Paid    : ", fmt.Sprint(p),  }, ""),
		},
		"\n",
	)

	fmt.Fprint(p_box, p_rows)
	tm.Println(p_box)
}

func print_stats(i float64, o float64, p float64) {
	tm.Println("\nMiner stats -> OK")
	s_box := tm.NewBox(75|tm.PCT, 5, 0)
	s_rows := strings.Join(
		[]string{
			strings.Join([]string{ "Immature: ", fmt.Sprint(i),  }, ""),
			strings.Join([]string{ "Owed    : ", fmt.Sprint(o),  }, ""),
			strings.Join([]string{ "Paid    : ", fmt.Sprint(p),  }, ""),
		},
		"\n",
	)

	fmt.Fprint(s_box, s_rows)
	tm.Println(s_box)

}

func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

	tm.Clear()
	tm.MoveCursor(1,1)
	tm.Println("Dark Fiber Stat Scraper")
	tm.Println("-----------------------------")

	input := MinerStatInput {
		MinerId: os.Getenv("DARK_FIBER_WALLET"),
		CoinType: os.Getenv("DARK_FIBER_COIN"),
	}

	tm.Println(strings.Join([]string { "Monitoring stats for: ", input.MinerId }, ""))
	tm.Println(strings.Join([]string { "Coin Type           : ", input.CoinType }, ""))

	// Payouts
	print_payouts(55.78, 45.00987, 20000.56)

	// Stats
	print_stats(16939.78, 45.23948, 9000.68726)

	tm.Flush()

	
	// payload := new(bytes.Buffer)
	// json.NewEncoder(payload).Encode(input)
	
	// fmt.Println(payload)

	// res, err := http.Post("https://aggrogator.dev/api/coins/payouts", "application/json", payload)


	// if (err != nil) {
	// 	fmt.Println(err.Error())
	// 	os.Exit(1)
	// }
	
	// fmt.Println(res)
}