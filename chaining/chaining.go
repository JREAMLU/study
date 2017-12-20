package main

import "fmt"

func main() {
	a := a().b().c()
	fmt.Println("++++++++++++: ", a)
}

type tt struct {
	id   int
	name string
}

func a() tt {
	return tt{
		id:   1,
		name: "abc",
	}
}

func (t tt) b() tt {
	return tt{
		id:   t.id + 1,
		name: t.name,
	}
}

func (t tt) c() tt {
	return tt{
		id:   t.id + 2,
		name: t.name,
	}
}

func (t tt) String() string {
	return fmt.Sprintf("id: %d, name: %s", t.id, t.name)
}

