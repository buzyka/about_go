package updater

import (
	"time"
)

type UpdatingPerson interface {
	GetBirthday() time.Time
	HaveBirthday()
}

type Updater struct {
	Persons []UpdatingPerson
}

func (u *Updater) UpdateBirthdaysForDate(date time.Time) {
	for _, person := range u.Persons {
		tmpBirthday := person.GetBirthday()
		if tmpBirthday.Month() == date.Month() && tmpBirthday.Day() == date.Day() {
			person.HaveBirthday()
		}
	}
}
