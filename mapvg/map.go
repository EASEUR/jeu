package mapvg

import (
	"fmt"
	"jeu/CombatGobelin"
	"jeu/inventaire"
	"jeu/logo"
	"math/rand"
	"time"
)

func Lore3() {
	var buffer string
	logo.Logoontop()
	fmt.Println("-----------------------------------------------------")
	fmt.Println("Le silence retombe sur Xnov...")
	fmt.Println("Les échos de ton dernier combat résonnent encore,")
	fmt.Println("mais déjà l’atmosphère change.")
	fmt.Println("-----------------------------------------------------")
	fmt.Println("1- Continuer")
	fmt.Println("\n\n\n\n")
	fmt.Scanln(&buffer)

	logo.Logoontop()
	fmt.Println("-----------------------------------------------------")
	fmt.Println("Ton PC assemblé. Ton Mentor vaincu.")
	fmt.Println("Tu as franchi la frontière invisible :")
	fmt.Println("l’élève n’existe plus, le guide apparaît.")
	fmt.Println("-----------------------------------------------------")
	fmt.Println("1- Continuer")
	fmt.Println("\n\n\n\n")
	fmt.Scanln(&buffer)

	logo.Logoontop()
	fmt.Println("-----------------------------------------------------")
	fmt.Println("Ici, ton voyage s’arrête.")
	fmt.Println("Non pas parce que tu échoues...")
	fmt.Println("...mais parce que tu transmets.")
	fmt.Println("-----------------------------------------------------")
	fmt.Println("1- Continuer")
	fmt.Println("\n\n\n\n")
	fmt.Scanln(&buffer)

	logo.Logoontop()
	fmt.Println("-----------------------------------------------------")
	fmt.Println("Tu ne quitteras jamais ces murs.")
	fmt.Println("Tu resteras, dans l’ombre, en silence,")
	fmt.Println("jusqu’à ce qu’un nouvel élève arrive...")
	fmt.Println("Et alors, tu deviendras son épreuve.")
	fmt.Println("-----------------------------------------------------")
	fmt.Println("1- Continuer")
	fmt.Println("\n\n\n")
	fmt.Scanln(&buffer)

	logo.Logoontop()
	fmt.Println("-----------------------------------------------------")
	fmt.Println("Ainsi commence ta nouvelle vie :")
	fmt.Println("TU ES DEVENU MENTOR.")
	fmt.Println("-----------------------------------------------------")
	fmt.Println("1- Fin")
	fmt.Println("\n\n\n\n\n")
	fmt.Scanln(&buffer)
}
func indice() {
	var buff string

	logo.Logoontop()
	fmt.Println("\033[1mVous êtes tombé sur un indice....\033[0m")
	fmt.Println("\033[3mJ’étais ton obstacle, je fus ta victoire.\nC’est là que tu as prouvé ta mémoire.\nLe PC complet, ta quête s’éclaire,\nRetourne donc là où tu as vaincu ton adversaire. \033[0m")
	fmt.Println("1- Conrinuer")
	fmt.Println("\n\n\n\n")
	fmt.Scanln(&buff)

}
func creapc(pc bool) bool {
	var saisis int
	if pc != true {

		logo.Logoontop()
		fmt.Println("\033[1mVous êtes arrivé dans le labo, un homme vous parle :\033[0m")
		fmt.Println("\033[3mQue me voulez vous ?\033[0m")
		fmt.Println("1- Construire le PC")
		fmt.Println("2- Partir en courant")
		fmt.Println("\n\n\n\n\n\n\n\n")
		fmt.Scanln(&saisis)
		if saisis == 1 {
			logo.Logoontop()
			fmt.Println("\033[3mAvez vous récuperé les 3 pièces du PC à obtenir ?\033[0m")
			fmt.Println("1- Oui")
			fmt.Println("\n\n\n\n\n\n\n\n\n\n")
			fmt.Scanln(&saisis)
			logo.Logoontop()
			fmt.Println("\033[3mVoici le PC monté et prêt à être utilisé\nBonne chance pour la suite des épreuves !\033[0m")
			fmt.Println("1- Continuez")
			fmt.Println("\n\n\n\n\n\n\n\n\n\n")
			fmt.Scanln(&saisis)

			return true
		} else {
			return false
		}
	}
	return true
}

func eprv1() bool {
	var r1 int
	logo.Logoontop()
	fmt.Println("\033[1mVous êtes arrivé à la première épreuve\033[0m")
	fmt.Println("\033[1mQuelle est la réponse à la vie ?\033[0m")
	fmt.Println("\n\n\n\n\n\n\n\n")
	fmt.Scanln(&r1)
	for r1 != 42 {
		logo.Logoontop()
		fmt.Println("\033[1mVous vous êtes trompé, veuillez réessayez\033[0m")
		fmt.Println("\033[1mQuelle est la réponse à la vie ?\033[0m")
		fmt.Println("\033[3mPetit indice :(-80538738812075974)^3 + (80435758145817515)^3 + (12602123297335631)^3\033[0m")
		fmt.Println("\n\n\n\n\n\n\n")
		fmt.Scanln(&r1)
	}
	return true
}

func eprv2() bool {
	var r2 int
	rand.Seed(time.Now().UnixNano())
	n1 := rand.Intn(1000) + 1
	n2 := rand.Intn(1000) + 1
	n3 := rand.Intn(1000) + 1

	resultat := (n1 * n2) / n3

	logo.Logoontop()
	fmt.Println("\033[1mVous êtes arrivé à la deuxième épreuve\033[0m")
	fmt.Println("Voici le petit calcul à résoudre :")
	fmt.Printf("(%d * %d) / %d = ? ", n1, n2, n3)
	fmt.Println("\n\n\n\n\n\n\n")
	fmt.Scanln(&r2)
	for r2 != resultat {
		logo.Logoontop()
		fmt.Println("\033[1mVous vous êtes trompé, veuillez réessayez\033[0m")
		fmt.Printf("(%d * %d) / %d = ? ", n1, n2, n3)
		fmt.Println("\033[3m(Ne mettez pas les nombres après la virgule)\033[0m")
		fmt.Println("\n\n\n\n\n\n\n")
		fmt.Scanln(&r2)
	}
	return true
}
func eprv3() bool {
	var buff string
	logo.Logoontop()
	fmt.Println("\033[1mVous êtes arrivé à la troisième épreuve\033[0m")
	fmt.Println("1- Continuez")
	fmt.Println("\n\n\n\n\n\n\n\n")
	fmt.Scanln(&buff)
	CombatGobelin.CombatInteractif()
	return true
}
func piece1() bool {
	var buff string
	eprv1()
	logo.Logoontop()
	fmt.Println("\033[1mBravo, vous avez obtenu la 1ere pièce\033[0m")
	fmt.Println("1- Continuez")
	fmt.Println("\n\n\n\n\n\n\n\n")
	fmt.Scanln(&buff)
	return true
}
func piece2() bool {
	var buff string
	eprv2()
	logo.Logoontop()
	fmt.Println("\033[1mBravo, vous avez obtenu la 2eme pièce\033[0m")
	fmt.Println("1- Continuez")
	fmt.Println("\n\n\n\n\n\n\n\n")
	fmt.Scanln(&buff)
	return true
}
func piece3() bool {
	var buff string
	eprv3()
	logo.Logoontop()
	fmt.Println("Vous avez trouvé la 3eme piece gg wp")
	fmt.Println("1- Continuez")
	fmt.Println("\n\n\n\n\n\n\n\n")
	fmt.Scanln(&buff)

	return true
}

func Mapmobile() {
	var obtpc bool
	var choix string
	var buff string
	var piecen1 bool
	var piecen2 bool
	var piecen3 bool
	var lore bool
	var obtp1 bool
	var obtp2 bool
	var obtp3 bool

	i := 2 // hauter
	j := 0 // largeur
	hauteur_max := 9
	largeur_max := 9
	for {
		logo.Logoontop()
		fmt.Printf("Vous vous trouvez actuellement à la case '%d' (hauteur) ; '%d' (largeur)\nQuelle action voulez-vous faire :\n", i, j)
		if i == hauteur_max {
			fmt.Println("Vous ne pouvez pas aller en haut, vous êtes au max de la hauteur")
		}
		if j == largeur_max {
			fmt.Println("Vous ne pouvez pas aller à droite, vous êtes au max de la largeur")
		}
		if i == 0 {
			fmt.Println("Vous ne pouvez pas aller en bas, vous êtes au minimum de la hauteur")
		}
		if j == 0 {
			fmt.Println("Vous ne pouvez pas aller à gauche, vous êtes au minimum de la largeur")
		}
		if i > 0 {
			fmt.Println("\033[33mS- Aller en bas\033[0m")
		}
		if i < hauteur_max {
			fmt.Println("\033[33mZ- Aller en haut\033[0m")
		}
		if j > 0 {
			fmt.Println("\033[33mQ- Aller à gauche\033[0m")
		}
		if j < largeur_max {
			fmt.Println("\033[33mD- Aller à droite\033[0m")
		}
		fmt.Println("\033[33mI- Iventaire\033[0m")
		fmt.Println("\n\n\n\n")
		fmt.Scanln(&choix)
		switch choix {
		case "Z":
			if i == hauteur_max {
				logo.Logoontop()
				fmt.Println("Je ne sais pas si vous avez lu les lignes qui étaient là avant mais vous ne pouvez pas faire ça")
				fmt.Println("Etes-vous idiot ?")
				fmt.Println("1- Oui")
				fmt.Println("\n\n\n\n\n\n\n")
				fmt.Scanln(&buff)
			} else {
				i++
			}
		case "S":
			if i == 0 {
				logo.Logoontop()
				fmt.Println("Je ne sais pas si vous avez lu les lignes qui étaient là avant mais vous ne pouvez pas faire ça")
				fmt.Println("Etes-vous idiot ?")
				fmt.Println("1- Oui")
				fmt.Println("\n\n\n\n\n\n\n")
				fmt.Scanln(&buff)
			} else {
				i--
			}

		case "D":
			if j == largeur_max {
				logo.Logoontop()
				fmt.Println("Je ne sais pas si vous avez lu les lignes qui étaient là avant mais vous ne pouvez pas faire ça")
				fmt.Println("Etes-vous idiot ?")
				fmt.Println("1- Oui")
				fmt.Println("\n\n\n\n\n\n\n")
				fmt.Scanln(&buff)
			} else {
				j++
			}

		case "Q":
			if j == 0 {
				logo.Logoontop()
				fmt.Println("Je ne sais pas si vous avez lu les lignes qui étaient là avant mais vous ne pouvez pas faire ça")
				fmt.Println("Etes-vous idiot ?")
				fmt.Println("1- Oui")
				fmt.Println("\n\n\n\n\n\n\n")
				fmt.Scanln(&buff)
			} else {
				j--
			}

		case "I":
			logo.Logoontop()
			fmt.Println("\n=$=$=$= Menu Inventaire =$=$=$=")
			fmt.Println("1- Voir les items")
			fmt.Println("2- Parler au marchand")
			fmt.Println("3- Quitter l’inventaire")
			fmt.Println("\n\n\n\n\n\n")

			var choice int
			fmt.Scanln(&choice)

			switch choice {
			case 1:
				inventaire.ShowInventory(obtp1, obtp2, obtp3, obtpc)
			case 2:
				inventaire.TalkToMerchant()
			case 3:
			}
			// ouvrir inventaire
		case "DEV":
			piecen1 = true
			piecen2 = true
			piecen3 = true
			obtp1 = true
			obtp2 = true
			obtp3 = true
		}

		if i == 8 && j == 2 {
			indice()
		}
		if i == 0 && j == 0 && !piecen1 {
			piecen1 = piece1()
			obtp1 = true
		}
		if i == 3 && j == 7 && !piecen2 {
			piecen2 = piece2()
			obtp2 = true
		}
		if i == 9 && j == 9 {
			if obtpc {
				Lore3()
				return
			}
			if obtp3 != true {
				piecen3 = piece3()
				obtp3 = true
			}

		}
		if piecen1 && piecen2 && piecen3 {
			if !lore {
				logo.Logoontop()
				fmt.Println("Vous avez réussi à réunir les 3 pièces du PC")
				fmt.Println("Vous devez vous rendre au labo, il se situe à '1' '1'")
				fmt.Println("1 -Continuez")
				fmt.Println("\n\n\n\n\n\n\n")
				fmt.Scanln(&buff)
				lore = true
			}
			if i == 1 && j == 1 && lore {
				obtpc = creapc(obtpc)
			}
		}
		if obtpc {
			obtp3 = false
			obtp2 = false
			obtp1 = false
		}
	}
}
