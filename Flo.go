package jeu

import "fmt"

type Inventory map[string]int

// Affichage de l'inventaire et
func accessInventory(inv Inventory) {
	fmt.Println("=== Inventaire ===")
	if len(inv) == 0 {
		fmt.Println("(Vide)")
		return
	}
	for item, qty := range inv {
		fmt.Printf("- %s (x%d)\n", item, qty)
	}
}

// Ajouter un item dans l'inventaire
func addInventory(inv Inventory, item string, qty int) {
	inv[item] += qty
}

// Retirer un item de l'inventaire
func removeInventory(inv Inventory, item string, qty int) {
	if inv[item] > qty {
		inv[item] -= qty
	} else {
		delete(inv, item)
	}
}

// Interface du marchand
func merchant(inv Inventory) {
	fmt.Println("=== Marchand ===")
	fmt.Println("1. Acheter une Potion de vie (gratuit)")
	fmt.Println("2. Quitter le marchand")

	var choice int
	fmt.Print("Votre choix : ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		addInventory(inv, "Potion de vie", 1)
		fmt.Println("Vous avez obtenu : Potion de vie")
	case 2:
		fmt.Println("Le marchand vous salue.")
	default:
		fmt.Println("Choix invalide.")
	}
}

func main() {
	// Inventaire de départ
	inventory := Inventory{
		"Potion de poison": 1,
		"Clé USB":          1,
	}

	// Boucle du menu principal
	for {
		fmt.Println("\n=== Menu Principal ===")
		fmt.Println("1. Voir l’inventaire")
		fmt.Println("2. Parler au marchand")
		fmt.Println("3. Quitter")

		var choice int
		fmt.Print("Votre choix : ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			accessInventory(inventory)
		case 2:
			merchant(inventory)
		case 3:
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println("Choix invalide, réessayez.")
		}
	}
}
