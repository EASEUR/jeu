package jeu

import (
	"fmt"
)

// Exemple de structure pour le personnage
type Personnage struct {
	Nom     string
	Classe  string
	Niveau  int
	Vie     int
	Force   int
	Agilite int
}

// Fonction pour afficher les informations du personnage
func displayInfo(p Personnage) {
	fmt.Println("Informations du personnage :")
	fmt.Printf("Nom      : %s\n", p.Nom)
	fmt.Printf("Classe   : %s\n", p.Classe)
	fmt.Printf("Niveau   : %d\n", p.Niveau)
	fmt.Printf("Vie      : %d\n", p.Vie)
	fmt.Printf("Force    : %d\n", p.Force)
	fmt.Printf("Agilité  : %d\n", p.Agilite)
}
