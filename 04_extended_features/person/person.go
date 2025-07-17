package person

import (
	"fmt"
	"time"
)

type Person struct {
	name string
	birthday time.Time
	age  int
}

func NewPerson(name string, birthday time.Time, age int) *Person {
	return &Person{
		name: name, 
		birthday: birthday,
		age: age,
	}
}

func (p *Person) GetName() string {
	return p.name
}

func (p *Person) SetName(name string) {
	p.name = name
}

func (p *Person) GetAge() int {
	return p.age
}

func (p *Person) SetAge(age int) {
	p.age = age
}

func (p *Person) GetBirthday() time.Time {
	return p.birthday
}

func (p *Person) SetBirthday(birthday time.Time) {
	p.birthday = birthday
}

func (p *Person) Greet() string {	
	return fmt.Sprintf("Hello, my name is %s and I am %d years old.", p.name, p.age)
}

func (p *Person) HaveBirthday() {
	p.age++
}

func (p *Person) IsOlderThan(age int) bool {
	return p.age > age
}



