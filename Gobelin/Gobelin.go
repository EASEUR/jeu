package Gobelin

import "fmt"

// Ceci définie la structure du Monstre
type Monster struct {
	Name          string
	MaxHealth     int
	CurrentHealth int
	AttackPoints  int
}

// Ceci définie la structure du Joueur
type Player struct {
	Name          string
	MaxHealth     int
	CurrentHealth int
}

// Cette func initialise le Gobelin
func InitGobelin() *Monster {
	return &Monster{
		Name:          "Erwann Kervoelen",
		MaxHealth:     300,
		CurrentHealth: 100,
		AttackPoints:  25,
	}
}

// Cette func affiche les infos du Gobelin
func Gobelin() {
	goblin := InitGobelin()
	fmt.Println("👹 Monstre créé :", goblin.Name)
	fmt.Printf("Points de vie : %d/%d\n", goblin.CurrentHealth, goblin.MaxHealth)
	fmt.Printf("Points d'attaque : %d\n", goblin.AttackPoints)
}

// Cette func met en place le pattern d'attaque du Gobelin
func goblinPattern(player *Player) {
	goblin := InitGobelin()
	for turn := 1; player.CurrentHealth > 0; turn++ {
		var damage int
		if turn%3 == 0 {
			damage = goblin.AttackPoints * 2
		} else {
			damage = goblin.AttackPoints
		}
		player.CurrentHealth -= damage
		if player.CurrentHealth < 0 {
			player.CurrentHealth = 0
		}
		fmt.Printf("%s inflige à %s %d de dégâts\n", goblin.Name, player.Name, damage)
		fmt.Printf("%s : %d/%d PV\n", player.Name, player.CurrentHealth, player.MaxHealth)
		if player.CurrentHealth == 0 {
			fmt.Printf("%s est vaincu !\n", player.Name)
			break
		}
	}
}
