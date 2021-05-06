package main

import (
	"fmt"
	"strings"

	tm "github.com/buger/goterm"
)

func print_header() {
	tm.Clear()
	tm.MoveCursor(1, 1)
	tm.Println(tm.Color("Dark Fiber Stat Scraper", tm.YELLOW))
	tm.Println(tm.Color("-----------------------------", tm.RED))
}

type TableManager struct {
	TimesRun int
}

func (t TableManager) PrintLoading() {
	tm.Println("\nLoading...")
}

func (t TableManager) MoveToBeginning() {
	tm.MoveCursor(1, 6)
}

func (t TableManager) PrintPayouts(c PayoutStructure) {
	// Payouts
	tm.Println(tm.Color("\nPayments", tm.BLUE))
	p_box := tm.NewBox(75|tm.PCT, 5, 0)
	p_rows := strings.Join(
		[]string{
			strings.Join([]string{"Immature: ", fmt.Sprint(c.Immature)}, ""),
			strings.Join([]string{"Owed    : ", fmt.Sprint(c.Owed)}, ""),
			strings.Join([]string{"Paid    : ", fmt.Sprint(c.Paid)}, ""),
		},
		"\n",
	)

	fmt.Fprint(p_box, p_rows)
	tm.Println(p_box)
}

func (t TableManager) PrintPayoutError(e error) {
	// Payouts
	tm.Println(tm.Color("Payout stats", tm.RED))
	p_box := tm.NewBox(75|tm.PCT, 5, 0)
	p_rows := strings.Join(
		[]string{
			"Error fetching payments:",
			e.Error(),
			"If this continues to happen, please ping @DukeFerdinand",
		},
		"\n",
	)

	fmt.Fprint(p_box, p_rows)
	tm.Println(p_box)
}

func (t TableManager) PrintStats(c StatStructure) {
	t.MoveToBeginning()
	tm.Println(tm.Color("\nMiner stats", tm.BLUE))
	s_box := tm.NewBox(75|tm.PCT, 6, 0)
	s_rows := strings.Join(
		[]string{
			fmt.Sprintf("Total Shares    : %v", c.TotalShares),
			fmt.Sprintf("Your Shares     : %v", c.Shares),
			fmt.Sprintf("Your Percentage : %v", c.UserPercentage),
			fmt.Sprintf("Estimated Payout: %v", c.EstimatedPayout),
		},
		"\n",
	)

	fmt.Fprint(s_box, s_rows)
	tm.Println(s_box)

}

func (t TableManager) PrintStatsError(e error) {
	t.MoveToBeginning()
	// Payouts
	tm.Println("\nPayments -> ERROR!")
	p_box := tm.NewBox(75|tm.PCT, 6, 0)
	p_rows := strings.Join(
		[]string{
			"Error fetching stats:",
			e.Error(),
			"If this continues to happen, please ping @DukeFerdinand",
		},
		"\n",
	)

	fmt.Fprint(p_box, p_rows)
	tm.Println(p_box)
}

func (t TableManager) WriteScreen() {
	tm.Flush()
}

func (t TableManager) PrintError(i interface{}) {
	tm.Println(i)
}
