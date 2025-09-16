package Gobelin

import "fmt"

// Définition de la structure Monster
type Monster struct {
	Name          string
	MaxHealth     int
	CurrentHealth int
	AttackPoints  int
}

func Gobelin() {
	// Création du Gobelin
	goblin := Monster{
		Name:          "Zero.exe",
		MaxHealth:     100,
		CurrentHealth: 100,
		AttackPoints:  15,
	}

	// Affichage des infos du Gobelin
	fmt.Println("👹 Monstre créé :", goblin.Name)
	fmt.Printf("Points de vie : %d/%d\n", goblin.CurrentHealth, goblin.MaxHealth)
	fmt.Printf("Points d'attaque : %d\n", goblin.AttackPoints)
}
