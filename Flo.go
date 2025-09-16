package jeu

import "fmt"

// Type pour représenter un inventaire
type Inventory map[string]int

// Fonction pour afficher l’inventaire
func showInventory(inv Inventory) {
	fmt.Println("=== Inventaire ===")
	if len(inv) == 0 {
		fmt.Println("(Vide)")
		return
	}
	for item, qty := range inv {
		fmt.Printf("- %s (x%d)\n", item, qty)
	}
}

// Fonction pour ajouter un item
func addInventory(inv Inventory, item string, qty int) {
	inv[item] += qty
}

// Fonction pour retirer un item
func removeInventory(inv Inventory, item string, qty int) {
	if qty <= 0 {
		return
	}
	if inv[item] > qty {
		inv[item] -= qty
	} else {
		delete(inv, item)
	}
}

// Parler au Marchand
func talkToMerchant(inv Inventory) {
	for {
		fmt.Println("\n=$=$=$= Marchand =$=$=$=")
		fmt.Println("1- Acheter une Potion de vie (gratuit)")
		fmt.Println("2- Vendre (retirer) un item")
		fmt.Println("3- Quitter le marchand")
		fmt.Println("\n\n\n\n\n\n")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addInventory(inv, "Potion de vie", 1)
			fmt.Println("Vous avez obtenu : Potion de vie")
		case 2:
			// Retirer un item (avec inventaire vide)
			if len(inv) == 0 {
				fmt.Println("Votre inventaire est vide, rien à vendre.")
				continue
			}

			// Retirer un item
			fmt.Println("Quel item voulez-vous retirer ?")
			items := make([]string, 0, len(inv))
			i := 1
			for item := range inv {
				fmt.Printf("%d. %s (x%d)\n", i, item, inv[item])
				items = append(items, item)
				i++
			}

			var idx int
			fmt.Scanln(&idx)

			if idx > 0 && idx <= len(items) {
				removeInventory(inv, items[idx-1], 1)
				fmt.Printf("Vous avez retiré : %s\n", items[idx-1])
			} else {
				fmt.Println("Choix invalide.")
			}
		case 3:
			fmt.Println("À bientôt !")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

// Menu Iventaire
func AccessInventory(inv Inventory) {
	for {
		fmt.Println("\n=$=$=$= Menu Inventaire =$=$=$=")
		fmt.Println("1- Voir les items")
		fmt.Println("2- Parler au marchand")
		fmt.Println("3- Quitter l’inventaire")
		fmt.Println("\n\n\n\n\n\n")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			showInventory(inv)
		case 2:
			talkToMerchant(inv)
		case 3:
			fmt.Println("Vous quittez l’inventaire.")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
