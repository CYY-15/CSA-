package main

import (
	"fmt"
)

// 游戏人物
type Person struct {
	Name   string
	Level  int
	Exp    int
	HP     int
	Attack int
}

// 攻击
type Attacker interface {
	Attack(target *Person)
}

type Warrior struct {
	Person
}

func (w *Warrior) Attack(target *Person) {
	fmt.Printf("%s 攻击了 %s\n", w.Name, target.Name)
	target.HP -= w.Person.Attack
	fmt.Printf("%s 剩余血量: %d\n", target.Name, target.HP)
}

func main() {
	warrior1 := &Warrior{
		Person{
			Name:   "warrior1",
			Level:  10,
			Exp:    100,
			HP:     1000,
			Attack: 200,
		},
	}

	warrior2 := &Warrior{
		Person{
			Name:   "warrior2",
			Level:  8,
			Exp:    80,
			HP:     800,
			Attack: 150,
		},
	}

	//warrior1 攻击 warrior2
	warrior1.Attack(&warrior2.Person)
}
