package main

import (
	"fmt"
)

type Player struct{
	name string
	health int
}

func ( p Player) TakeDamage(amount int) {
	p.health= p.health-amount
}


func ( p *Player) Heal(amount int) {
	p.health= p.health+amount
}

func IsDead(p Player) bool{
		if p.health==0 {
			return true
		}
		return false
	}

func main(){
	p1 := Player{ name: "Abel", health: 100}
	p1.TakeDamage(50)
	p1.Heal(30)
	fmt.Println(p1.name)
	fmt.Println(p1.health)
	fmt.Println(IsDead(p1))

}
