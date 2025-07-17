package main

import (
	"fmt"
	"time"

	"github.com/buzyka/about_go/04_extended_features/filter"
	"github.com/buzyka/about_go/04_extended_features/person"
	"github.com/buzyka/about_go/04_extended_features/updater"
)

func exFilter() {
	p1 := person.NewPerson(
		"Alice", 
		time.Date(2000, 5, 1, 0, 0, 0, 0, time.UTC), 
		30,
	)
	p2 := person.NewPerson(
		"Bob", 
		time.Date(2001, 6, 1, 0, 0, 0, 0, time.UTC),
		31,
	)
	older := filter.GetOlderPerson(p1, p2)
	fmt.Println(older.GetName(), "is older.")
}

func exUpdater(updateDate time.Time) {
	p1 := person.NewPerson(
		"Charlie", 
		time.Date(1995, 5, 5, 0, 0, 0, 0, time.UTC), 
		28,
	)
	p2 := person.NewPerson(
		"Dave", 
		time.Date(1995, 10, 10, 0, 0, 0, 0, time.UTC), 
		29,
	)

	updater := &updater.Updater{
		Persons: []updater.UpdatingPerson{p1, p2},
	}

	updater.UpdateBirthdaysForDate(updateDate)
	fmt.Println(p1.Greet())
	fmt.Println(p2.Greet())
}


func main() {
	// Used filter package with FilteringPerson interface
	exFilter()

	// Used updater package with UpdatingPerson interface
	exUpdater(time.Date(2025, 5, 5, 0, 0, 0, 0, time.UTC))
}
