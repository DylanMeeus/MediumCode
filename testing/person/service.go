package person

type repository interface {
	create(Person)
	lookup(string) Person
}

type inmemory struct {
	// connection information
	people []Person
}

func (i inmemory) create(p Person) {
	i.people = append(i.people, p)
}

func (i inmemory) lookup(n string) Person {
	for _, p := range i.people {
		if p.Name == n {
			return p
		}
	}
	return Person{}
}

type Service struct {
	repo repository
}

func (s Service) createPerson(p Person) {
	s.repo.create(p)
}

func (s Service) lookup(n string) Person {
	return s.repo.lookup(n)
}
