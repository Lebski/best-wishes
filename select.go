package main

import (
	"time"
)

func selectGreeting(greetings *[]Greeting) []Greeting {
	selected := []Greeting{}
	time := time.Now()
	// convert to Maps later for quicker runtime
	for _, g := range *greetings {
		if g.AllTime {
			selected = append(selected, g)
			continue
		}

		// Morning
		if time.Hour() < 11 && g.Morning {
			selected = append(selected, g)
			continue
		}

		// Noon
		if time.Hour() > 11 && time.Hour() < 16 && g.Noon {
			selected = append(selected, g)
			continue
		}

		// Evening
		if time.Hour() > 16 && g.Noon {
			selected = append(selected, g)
			continue
		}

		// if Sunday, Friday or Saturday
		if (time.Weekday() == 0 || time.Weekday() > 5) && g.Weekend {
			selected = append(selected, g)
			continue
		}

		if time.Weekday() >= 1 && time.Weekday() <= 2 && g.StartOfTheWeek{
			selected = append(selected, g)
			continue
		}

		if time.YearDay() >= 360 && g.Season == "NY-Pre" {
			selected = append(selected, g)
			continue
		}
		if time.YearDay() >= 1 &&  time.YearDay() <= 5 && g.Season == "NY" {
			selected = append(selected, g)
			continue
		}

		if time.YearDay() <= 358 && time.YearDay() > 346 && g.Season == "Christmas" {
			selected = append(selected, g)
			continue
		}
	}
	//fmt.Println(selected)
	return selected
}

