package filter


type FilteringPerson interface {
	GetAge() int
	GetName() string
	IsOlderThan(age int) bool
}


func GetOlderPerson(p1, p2 FilteringPerson) FilteringPerson {
	if p1.IsOlderThan(p2.GetAge()) {
		return p1
	}
	return p2
}
