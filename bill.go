package main

import (
	"fmt"
	"math"
)

type Plays map[string]Play

type Play struct {
	Name string
	Type string
}

type Performance struct {
	PlayID   string `json:"playID"`
	Audience int    `json:"audience"`
}

type Invoice struct {
	Customer     string        `json:"customer"`
	Performances []Performance `json:"performances"`
}

func playType(play Play) string {
	return play.Type
}

func playName(play Play) string {
	return play.Name
}

func playFor(plays Plays,perf Performance)Play{
	return plays[perf.PlayID]
}

func getAmount(perf Performance, play Play) float64 {
	result := 0.0

	switch playType(play) {
	case "tragedy":
		result = 40000
		if perf.Audience > 30 {
			result += 1000 * (float64(perf.Audience - 30))
		}
	case "comedy":
		result = 30000
		if perf.Audience > 20 {
			result += 10000 + 500*(float64(perf.Audience-20))
		}
		result += 300 * float64(perf.Audience)
	default:
		panic(fmt.Sprintf("unknow type: %s", play.Type))
	}
	return result
}


func volumeCreditsFor(perf Performance, plays Plays) float64 {
	result := 0.0
	result += math.Max(float64(perf.Audience-30), 0)
	// add extra credit for every ten comedy attendees
	if "comedy" == playType(playFor(plays, perf)) {
		result += math.Floor(float64(perf.Audience / 5))
	}
	return result
}


func totalAmount(perfs []Performance,plays Plays)float64{
	result := 0.0
	for _, perf := range perfs {
		play := playFor(plays,perf)
		result += getAmount(perf, play)
	}
	return result
}

func totalVolumeCredits(perfs []Performance,plays Plays)float64{
	result := 0.0
	for _, perf := range perfs {
		result += volumeCreditsFor(perf,plays)
	}
	return result
}

func renderPlainText(invoice Invoice,plays Plays)string{
	result := fmt.Sprintf("Statement for %s\n", invoice.Customer)
	
	for _, perf := range invoice.Performances {
		result += fmt.Sprintf("  %s: $%.2f (%d seats)\n", playName(playFor(plays,perf)), getAmount(perf, playFor(plays,perf))/100, perf.Audience)
	}

	result += fmt.Sprintf("Amount owed is $%.2f\n", totalAmount(invoice.Performances,plays)/100)
	result += fmt.Sprintf("you earned %.0f credits\n", totalVolumeCredits(invoice.Performances,plays))
	return result
}

func statement(invoice Invoice, plays Plays) string {
	return renderPlainText(invoice,plays)
}



func main() {
	inv := Invoice{
		Customer: "Bigco",
		Performances: []Performance{
			{PlayID: "hamlet", Audience: 55},
			{PlayID: "as-like", Audience: 35},
			{PlayID: "othello", Audience: 40},
		}}
	plays := map[string]Play{
		"hamlet":  {Name: "Hamlet", Type: "tragedy"},
		"as-like": {Name: "As You Like It", Type: "comedy"},
		"othello": {Name: "Othello", Type: "tragedy"},
	}

	bill := statement(inv, plays)
	fmt.Println(bill)
}
