package main

import (
	"fmt"
	"math"
)

type Plays map[string]Play

type Play struct {
	Name string
	Type string
	BillCalculator BillCalculate
}


func NewPlay(name, playType string) (Play) {
	var calculator BillCalculate
	switch playType {
	case "tragedy":
		calculator = Tragedy{}
	case "comedy":
		calculator = Comedy{}
	default:
		panic(fmt.Sprintf("unknow type: %s", playType))
	}
	return Play{Name: name, Type: playType, BillCalculator: calculator}
}


type Tragedy struct{}
type Comedy struct{}
type BillCalculate interface {
	CalculateAmount(perf Performance) float64
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

func playFor(plays Plays, perf Performance) Play {
	return plays[perf.PlayID]
}

func amountFor(perf Performance, play Play) float64 {
	return play.BillCalculator.CalculateAmount(perf)
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

func totalAmountFor(perfs []Performance, plays Plays) float64 {
	result := 0.0
	for _, perf := range perfs {
		play := playFor(plays, perf)
		result += amountFor(perf, play)
	}
	return result
}

func totalVolumeCreditsFor(perfs []Performance, plays Plays) float64 {
	result := 0.0
	for _, perf := range perfs {
		result += volumeCreditsFor(perf, plays)
	}
	return result
}

type Bill struct {
	Customer           string
	Rates              []Rate
	TotalAmount        float64
	TotalVolumeCredits float64
}

type Rate struct {
	Play          Play
	Amount        float64
	VolumeCredits float64
	Audience      int
}

func renderPlainText(bill Bill) string {

	result := fmt.Sprintf("Statement for %s\n", bill.Customer)

	for _, rate := range bill.Rates {
		result += fmt.Sprintf("  %s: $%.2f (%d seats)\n", playName(rate.Play), rate.Amount/100, rate.Audience)
	}

	result += fmt.Sprintf("Amount owed is $%.2f\n", bill.TotalAmount/100)
	result += fmt.Sprintf("you earned %.0f credits\n", bill.TotalVolumeCredits)
	return result
}

func statement(invoice Invoice, plays Plays) string {

	rates := []Rate{}
	for _, perf := range invoice.Performances {

		rate := Rate{
			Play:          playFor(plays, perf),
			Amount:        amountFor(perf, playFor(plays, perf)),
			VolumeCredits: volumeCreditsFor(perf, plays),
			Audience:      perf.Audience,
		}
		rates = append(rates, rate)
	}

	bill := Bill{
		Customer:           invoice.Customer,
		Rates:              rates,
		TotalAmount:        totalAmountFor(invoice.Performances, plays),
		TotalVolumeCredits: totalVolumeCreditsFor(invoice.Performances, plays),
	}


	return renderPlainText(bill)
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
		"hamlet":  NewPlay("Hamlet","tragedy"),
		"as-like": NewPlay("As You Like It", "comedy"),
		"othello": NewPlay("Othello","tragedy"),
	}

	bill := statement(inv, plays)
	fmt.Println(bill)
}



func (t Tragedy) CalculateAmount(perf Performance) float64{
	result := 40000.0
	if perf.Audience > 30 {
		result += 1000 * (float64(perf.Audience - 30))
	}
	return result
}


func (t Comedy) CalculateAmount(perf Performance) float64{
	result := 30000.0
	if perf.Audience > 20 {
		result += 10000 + 500*(float64(perf.Audience-20))
	}
	result += 300 * float64(perf.Audience)
	return result
}

