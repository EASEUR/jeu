package inventaire

import (
	"fmt"
	"jeu/logo"
)

var Inv *Inventory

const MaxInventory = 10

type Inventory map[string]int

var inv Inventory = make(map[string]int)
var Inv *Inventory = &inv // <- ça initialise le pointeur

// Fonction pour compter le nombre d'items dans l'inventaire (10max)
func CountItems(inv Inventory) int {
	total := 0
	for _, qty := range inv {
		total += qty
	}
	return total
}

// Fonction pour afficher l’inventaire
func ShowInventory(p1, p2, p3, ppc bool) {
	var buff string
	var total int
	logo.Logoontop()
	fmt.Println("\033[1m=$=$=$= Inventaire =$=$=$=\033[0m")

	for item, qty := range inv {
		fmt.Printf("\033[31m- %s (x%d)\033[0m\n", item, qty)
	}
	if p1 {
		fmt.Println("\033[34mCarte mère\033[0m")
		total++
	}
	if p2 {
		fmt.Println("\033[34mProcesseur\033[0m")
		total++
	}
	if p3 {
		fmt.Println("\033[34mCarte Graphique\033[0m")
		total++
	}
	if ppc {
		fmt.Println("\033[34mPC assemblé\033[0m")
		total++
	}
	if p1 != true {
		fmt.Println("")
	}
	if p2 != true {
		fmt.Println("")
	}
	if p3 != true {
		fmt.Println("")
	}
	if ppc != true {
		fmt.Println("")
	}
	total += CountItems(inv)
	fmt.Printf("Total : %d/%d items\n", total, MaxInventory)
	fmt.Println("1 - Retour au jeu")
	fmt.Println("\n\n\n\n")

	fmt.Scanln(&buff)

}

// Fonction pour ajouter un item (avec limite)
func addInventory(inv Inventory, item string, qty int) bool {
	if CountItems(inv)+qty > MaxInventory {
		return false
	}
	inv[item] += qty
	return true
}

// Fonction pour retirer un item
func RemoveInventory(inv Inventory, item string, qty int) {
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
func TalkToMerchant() {
	var buff string
	for {
		logo.Logoontop()
		fmt.Println("\033[1m\n=$=$=$= Marchand =$=$=$=\033[0m")
		fmt.Println("1- Acheter une Potion")
		fmt.Println("2- Retirer un item")
		fmt.Println("3- Quitter le marchand")
		fmt.Println("\n\n\n\n\n\n")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			logo.Logoontop()
			// Ajouter un item + rappel place dispo inventaire
			if addInventory(inv, "Potion de vie", 1) {
				fmt.Println("Vous avez obtenu : \033[31mPotion de vie\033[0m")
			} else {
				fmt.Println("Inventaire plein ! Vous ne pouvez rien acheter.")
			}
			fmt.Printf("Place de l’inventaire : %d/%d\n", CountItems(inv), MaxInventory)
			fmt.Println("1- Continuez")
			fmt.Println("\n\n\n\n\n\n\n")
			fmt.Scanln(&buff)

		case 2:
			// Retirer un item (avec inventaire vide)
			if len(inv) == 0 {
				logo.Logoontop()
				fmt.Println("Votre inventaire est vide, rien à vendre.")
				fmt.Println("\n\n\n\n\n\n\n\n\n")
				fmt.Scanln(&buff)
				continue
			}

			// Retirer un item + rappel place dispo inventaire
			logo.Logoontop()
			fmt.Println("Quel item voulez-vous retirer ?")
			items := make([]string, 0, len(inv))
			i := 1
			for item := range inv {
				fmt.Printf("%d. %s (x%d)\n", i, item, inv[item])
				items = append(items, item)
				i++
			}
			fmt.Println("\n\n\n\n\n\n\n\n\n")

			var idx int
			fmt.Scanln(&idx)
			logo.Logoontop()

			if idx > 0 && idx <= len(items) {
				RemoveInventory(inv, items[idx-1], 1)
				fmt.Printf("Vous avez retiré : %s\n", items[idx-1])
			} else {
				fmt.Println("Choix invalide.")
			}
			fmt.Printf("Place de l’inventaire : %d/%d\n", CountItems(inv), MaxInventory)
			fmt.Println("1- Continuez")
			fmt.Println("\n\n\n\n\n\n\n")
			fmt.Scanln(&buff)
		case 3:
			logo.Logoontop()
			fmt.Println("À bientôt !")
			fmt.Println("1- Continuez")
			fmt.Println("\n\n\n\n\n\n\n\n")
			fmt.Scanln(&buff)
			return
		}
	}
}

// Menu Iventaire
// func AccessInventory(inv Inventory) {
// 	for {
// 		fmt.Println("\n=$=$=$= Menu Inventaire =$=$=$=")
// 		fmt.Println("1- Voir les items")
// 		fmt.Println("2- Parler au marchand")
// 		fmt.Println("3- Quitter l’inventaire")
// 		fmt.Println("\n\n\n\n\n\n")

// 		var choice int
// 		fmt.Scanln(&choice)

// 		switch choice {
// 		case 1:
// 			ShowInventory(inv)
// 		case 2:
// 			talkToMerchant(inv)
// 		case 3:
// 			fmt.Println("Vous quittez l’inventaire.")
// 			return
// 		default:
// 			fmt.Println("Choix invalide.")
