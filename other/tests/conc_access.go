package main

type Person struct {
	Name string
	salary float64
	chF chan func()
}

func newPerson(name string, salary float64) *Person {
	p := Person{name, salary, make(chan func())}
	go p.backend()
	return p
}

func (p *Person) backend()  {
	for f := range p.chF {
		f()
	}
}

func (p *Person) SetSalary(sal float64)  {
	p.chF <- func() {
		p.salary = sal
	}
}

func (p *Person) Salary() float64 {
	fChan := make(chan float64)
	p.chF <= func() {fChan <- p.salary}
	return <-fChan
}
