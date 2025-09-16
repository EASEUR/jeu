package menu

import (
	"fmt"
	"jeu/logo"
	"jeu/mapvg"
	"jeu/test"
)

// variable
var sousmenu int
var choixmenu int
var buff1 string

func Lore() {
	logo.Logoontop()
	fmt.Println("-----------------------------------------------------")
	fmt.Println("Dans une école mystérieuse, réputée pour former")
	fmt.Println("les plus grands esprits du numérique, chaque élève")
	fmt.Println("n’est pas simplement là pour apprendre.")
	fmt.Println()
	fmt.Println("Ici, l’enseignement est une épreuve, un jeu dangereux")
	fmt.Println("où seuls les plus débrouillards atteignent le sommet.")
	fmt.Println("-----------------------------------------------------")
	fmt.Println("1- Continuez")
	fmt.Println("\n")
	fmt.Scanln(&buff1)
	logo.Logoontop()
	fmt.Println("-----------------------------------------------------")
	fmt.Println("Les professeurs, appelés Mentores, ne sont pas de")
	fmt.Println("simples guides : ce sont des gardiens d’un secret")
	fmt.Println("ancestral. Ils détiennent le savoir, mais aussi le")
	fmt.Println("pouvoir de l’empêcher de se transmettre.")
	fmt.Println("Leur rôle ? Tester les élèves, parfois jusqu’à l’échec.")
	fmt.Println("-----------------------------------------------------")
	fmt.Println("1- Continuez")
	fmt.Println("\n\n")
	fmt.Scanln(&buff1)
	logo.Logoontop()
	fmt.Println("-----------------------------------------------------")
	fmt.Println("Au cœur de l’établissement, caché dans un couloir oublié,")
	fmt.Println("se trouve le Labo, une salle interdite dont l’entrée")
	fmt.Println("n’apparaît qu’à ceux qui en sont dignes.")
	fmt.Println()
	fmt.Println("On raconte que quiconque parvient à y assembler")
	fmt.Println("un Ordinateur parfait peut défier les Mentors.")
	fmt.Println("-----------------------------------------------------")
	fmt.Println("1- Continuez")
	fmt.Println("\n")
	fmt.Scanln(&buff1)
	logo.Logoontop()
	fmt.Println("-----------------------------------------------------")
	fmt.Println("Avant de pénétrer dans les couloirs de l'École Mystérieuse,")
	fmt.Println("tu dois révéler qui tu es vraiment.")
	fmt.Println("Chaque élève possède une histoire, une force et une faiblesse.")
	fmt.Println("Ta destinée dépendra de l'identité que tu choisis.")
	fmt.Println()
	fmt.Println("Prépare-toi à forger ton personnage...")
	fmt.Println("-----------------------------------------------------")
	fmt.Println("1- Continuez")
	fmt.Println("\n")
	fmt.Scanln(&buff1)
}
func Lore2() {
	logo.Logoontop()
	fmt.Println("-----------------------------------------------------")
	fmt.Println("Ton esprit s'éveille dans une salle plongée dans l'obscurité.")
	fmt.Println("Un grincement résonne... puis une porte s'ouvre lentement.")
	fmt.Println("Tu viens d’être libéré dans la cour principale de l’École Mystérieuse.")
	fmt.Println("Devant toi s’étend une immense carte labyrinthique :")
	fmt.Println("salles de classe abandonnées, bibliothèque silencieuse, gymnase poussiéreux,")
	fmt.Println("et des couloirs qui semblent n’avoir ni début ni fin.")
	fmt.Println("-----------------------------------------------------")
	fmt.Println("1- Continuez")
	fmt.Println("\n")
	fmt.Scanln(&buff1)
	logo.Logoontop()
	fmt.Println("-----------------------------------------------------")
	fmt.Println("Ton objectif est clair :")
	fmt.Println("   → Retrouver les 3 pièces essentielles pour assembler ton premier PC.")
	fmt.Println("   Ces fragments sont dispersés et jalousement protégés :")
	fmt.Println("      1. La Carte Mère, symbole du contrôle.")
	fmt.Println("      2. Le Processeur, cœur de l’intelligence.")
	fmt.Println("      3. La Carte Graphique, clé de la vision.")
	fmt.Println("-----------------------------------------------------")
	fmt.Println("1- Continuez")
	fmt.Println("\n")
	fmt.Scanln(&buff1)
	logo.Logoontop()
	fmt.Println("-----------------------------------------------------")
	fmt.Println("Une fois réunies, il te faudra atteindre le Labo,")
	fmt.Println("le seul lieu capable d’assembler ces fragments en une machine parfaite.")
	fmt.Println("Prépare-toi, élève.")
	fmt.Println("Chaque pas que tu feras dans cette école te rapprochera du pouvoir...")
	fmt.Println("ou de ta perte.")
	fmt.Println("-----------------------------------------------------")
	fmt.Println("1 - Continuer dans la map")
	fmt.Println("\n\n")
	fmt.Scanln(&buff1)
}
func Menu() {
	for true {
		logo.Logoontop()
		fmt.Println("\033[1mBienvenue sur Xtrack :\033[0m") // deroulement du menu
		fmt.Println("\033[3mPetit conseil, restez en CapsLock durant le jeu :3\033[0m")
		fmt.Println("\033[1mChoisissez un des menus ci dessous :\033[0m")
		fmt.Println("1- Jouons !")
		fmt.Println("2- Paramètres")
		fmt.Println("3- Ouvrir l'inventaire")
		fmt.Println("4- Crédits")
		fmt.Println("5- Quittez")
		fmt.Println("\n\n")
		fmt.Scanln(&choixmenu) // choix du menu
		switch choixmenu {
		case 1:
			//commencer a jouer
			logo.Logoontop()
			fmt.Println("\033[1mEtes-vous sûr de bien vouloir rentrer dans le Souk ?\033[0m") // verification
			fmt.Println("1- Oui !")
			fmt.Println("2- Non...")
			fmt.Println("\n\n\n\n\n\n\n")
			fmt.Scanln(&sousmenu) // choix sous menu
			switch sousmenu {
			case 1:
				Lore()
				test.CharacterCreation()
				Lore2()
				mapvg.Mapmobile()
				break // sortir
			case 2:
				break // sortir
			}
		case 2: // ouverture des parametres
			logo.Logoontop()
			fmt.Println("\033[1mEParametres :\033[0m")
			fmt.Println("1- Langues")
			fmt.Println("2- Quitter")
			fmt.Println("\n\n\n\n\n\n")
			fmt.Scanln(&sousmenu)
			if sousmenu == 2 {
				break
			}
			// changer la langue mais y'en a qu'une
			logo.Logoontop()
			fmt.Println("\033[1mChoisir la langue :\033[0m")
			fmt.Println("1- Français")
			fmt.Println("2- Quitter")
			fmt.Println("\n\n\n\n\n\n")
			fmt.Scanln(&buff1)
			break
		case 3:
			logo.Logoontop() // ouvrir l'inventaire
			fmt.Println("Pourquoi tu veux voir l'inventaire, t'as même pas commencer le jeu, 0/20")
			fmt.Println("Et pour info y'a rien dans l'inventaire gros nullos")
			fmt.Println("1- Quitter")
			fmt.Println("\n\n\n\n\n\n")
			fmt.Scanln(&buff1)
			break

		case 4:
			logo.Logoontop()
			fmt.Println("\033[1mFait par :\033[0m")
			fmt.Println("- Oscar V")
			fmt.Println("- Florent F")
			fmt.Println("- Enzo D-R")
			fmt.Println("1- Quitter")
			fmt.Println("\n\n\n\n")
			fmt.Scanln(&buff1)
		case 5:
			logo.Clear()
			return // fermer prg
		}
	}
}

/*
fmt.Println("=== Styles de texte ===")
fmt.Println("\033[0mTexte normal\033[0m")
fmt.Println("\033[3mTexte en italique\033[0m")
fmt.Println("\033[1mTexte en gras\033[0m")
fmt.Println("\033[4mTexte souligné\033[0m")

fmt.Println("\n=== Couleurs ===")
fmt.Println("\033[31mTexte rouge\033[0m")
fmt.Println("\033[32mTexte vert\033[0m")
fmt.Println("\033[33mTexte jaune\033[0m")
fmt.Println("\033[34mTexte bleu\033[0m")
fmt.Println("\033[35mTexte magenta\033[0m")
fmt.Println("\033[36mTexte cyan\033[0m")

fmt.Println("\n=== Combinaisons ===")
fmt.Println("\033[1;31mTexte gras + rouge\033[0m")
fmt.Println("\033[3;34mTexte italique + bleu\033[0m")
fmt.Println("\033[4;32mTexte souligné + vert\033[0m")
*/
