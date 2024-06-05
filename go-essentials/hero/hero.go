package hero

import (
	"fmt"
	"strings"
)

type Soldier interface {
	Match()
}

type Hero interface {
	Soldier
	Attack()
}

type Character struct {
	name     string
	hitPoint int64
	stamina  int64
	speed    int64
}

type Knight struct {
	Character
}

func (k *Knight) Attack() {
	fmt.Println("Charging")
}

func (k *Knight) Match() {
	k.stamina -= 10
	fmt.Println("Matching")
}

type Archer struct {
	Character
}

func (k *Archer) Attack() {
	fmt.Println("Firing")
}

func (k *Archer) Match() {
	k.stamina -= 10
	fmt.Println("Matching")
}

type SpearMan struct {
	Character
}

func (k *SpearMan) Attack() {
	fmt.Println("Defend")
}

func (k *SpearMan) Match() {
	k.stamina -= 10
	fmt.Println("Matching")
}

type SwordMan struct {
	Character
}

func (k *SwordMan) Attack() {
	fmt.Println("Attack")
}

func (k *SwordMan) Match() {
	k.stamina -= 10
	fmt.Println("Matching")
}

func (c *Character) GetInformation() string {
	output := ""
	output += strings.Repeat("*", 20) + "\n"
	output += fmt.Sprintf("Name: %v\n", c.name)
	output += fmt.Sprintf("Hit Point: %v\n", c.hitPoint)
	output += fmt.Sprintf("Stamina: %v\n", c.stamina)
	output += fmt.Sprintf("Speed: %v\n", c.speed)
	output += strings.Repeat("*", 20) + "\n"
	return output
}

func NewKnight() *Knight {
	return &Knight{
		Character{
			name:     "Knight",
			hitPoint: 200,
			stamina:  100,
			speed:    120,
		},
	}
}

func NewArcher() *Archer {
	return &Archer{
		Character{
			name:     "Archer",
			hitPoint: 80,
			stamina:  100,
			speed:    80,
		},
	}
}

func NewSpearMan() *SpearMan {
	return &SpearMan{
		Character{
			name:     "Spearman",
			hitPoint: 120,
			stamina:  100,
			speed:    60,
		},
	}
}

func NewSwordMan() *SwordMan {
	return &SwordMan{
		Character{
			name:     "Swordman",
			hitPoint: 100,
			stamina:  100,
			speed:    80,
		},
	}
}

func HeroAttack(hero Hero) {
	switch hero.(type) {
	case *SwordMan:
		fmt.Println("SwordMan")
	case *Knight:
		fmt.Println("Knight")
	}

	archer, isArcher := hero.(*Archer)
	if isArcher {
		fmt.Println(archer.name)
	}

	spearMan, isSpearMan := hero.(*SpearMan)
	if isSpearMan {
		fmt.Println(spearMan.name)
	}

	hero.Match()
	hero.Attack()
}
