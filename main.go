package main

import (
	"log"
	"os"
	"sync"
	"time"

	godotenv "github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Dark fiber header/app identifier
	print_header()

	// Setup meat of app, init main data structure
	sc := MinerStatConfig{
		MinerId:  os.Getenv("DARK_FIBER_WALLET"),
		CoinType: os.Getenv("DARK_FIBER_COIN"),
	}
	sc.PrintConfig()
	tm := TableManager{
		TimesRun: 0,
	}
	tm.PrintLoading()
	tm.WriteScreen()

	var wg sync.WaitGroup
	//! Nothing changeable should happen outside of this event loop!
	for {
		wg.Add(2)
		// Stats
		go func() {
			defer wg.Done()
			stats, stat_err := sc.GetStats()
			// Stats
			if stat_err != nil {
				tm.PrintStatsError(stat_err)
			} else {
				tm.PrintStats(stats)
			}
			tm.WriteScreen()
		}()

		// Payouts
		go func() {
			defer wg.Done()
			payouts, payout_err := sc.GetPayouts()

			// TODO: find a way to remove this and still print tables properly
			time.Sleep(time.Millisecond * 500)
			// Payouts
			if payout_err != nil {
				tm.PrintPayoutError(payout_err)
			} else {
				tm.PrintPayouts(payouts)
			}
			tm.WriteScreen()
		}()

		wg.Wait()
		time.Sleep(time.Second * 10)

		tm.TimesRun += 1
	}
}
