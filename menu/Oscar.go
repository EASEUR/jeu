package menu

import (
	"fmt"
	"jeu/test"
	"os"
	"os/exec"
	"runtime"
)


// variable
var sousmenu int
var choixmenu int

func Lore() {
	Logoontop()
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
	fmt.Scanln(&sousmenu)
	Logoontop()
	fmt.Println("-----------------------------------------------------")
	fmt.Println("Les professeurs, appelés Mentores, ne sont pas de")
	fmt.Println("simples guides : ce sont des gardiens d’un secret")
	fmt.Println("ancestral. Ils détiennent le savoir, mais aussi le")
	fmt.Println("pouvoir de l’empêcher de se transmettre.")
	fmt.Println("Leur rôle ? Tester les élèves, parfois jusqu’à l’échec.")
	fmt.Println("-----------------------------------------------------")
	fmt.Println("1- Continuez")
	fmt.Println("\n\n")
	fmt.Scanln(&sousmenu)
	Logoontop()
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
	fmt.Scanln(&sousmenu)
	Logoontop()
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
	fmt.Scanln(&sousmenu)
}
func Clear() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else { // Linux, macOS
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
func Logoontop() {
	Clear()
	fmt.Println("\n\n\n\n\n")
	fmt.Println(`▀████    ▐████▀     ███        ▄████████    ▄████████  ▄████████    ▄█   ▄█▄ 
  ███▌   ████▀  ▀█████████▄   ███    ███   ███    ███ ███    ███   ███ ▄███▀ 
   ███  ▐███       ▀███▀▀██   ███    ███   ███    ███ ███    █▀    ███▐██▀   
   ▀███▄███▀        ███   ▀  ▄███▄▄▄▄██▀   ███    ███ ███         ▄█████▀    
   ████▀██▄         ███     ▀▀███▀▀▀▀▀   ▀███████████ ███        ▀▀█████▄    
  ▐███  ▀███        ███     ▀███████████   ███    ███ ███    █▄    ███▐██▄   
 ▄███     ███▄      ███       ███    ███   ███    ███ ███    ███   ███ ▀███▄ 
████       ███▄    ▄████▀     ███    ███   ███    █▀  ████████▀    ███   ▀█▀ 
                              ███    ███                           ▀         `)
	fmt.Println("\n\n\n")
}
func Menu() {
	for true {
		Logoontop()
		fmt.Println("Bienvenue sur Xtrack :") // deroulement du menu
		fmt.Println("Choisissez un des menus ci dessous :")
		fmt.Println("1- Jouons !")
		fmt.Println("2- Paramètres")
		fmt.Println("3- Ouvrir l'inventaire")
		fmt.Println("4- Crédits")
		fmt.Println("5- Quittez")
		fmt.Println("\n\n\n")
		fmt.Scanln(&choixmenu) // choix du menu
		switch choixmenu {
		case 1:
			//commencer a jouer
			Logoontop()
			fmt.Println("Etes-vous sûr de bien vouloir rentrer dans le Souk ?") // verification
			fmt.Println("1- Oui !")
			fmt.Println("2- Non...")
			fmt.Println("\n\n\n\n\n\n")
			fmt.Scanln(&sousmenu) // choix sous menu
			switch sousmenu {
			case 1:
				Lore()
				test.CharacterCreation()
				break // sortir
			case 2:
				break // sortir
			}
		case 2: // ouverture des parametres
			Logoontop()
			fmt.Println("Paramètres")
			fmt.Println("1- Langues")
			fmt.Println("2- Quitter")
			fmt.Println("\n\n\n\n\n\n")
			fmt.Scanln(&sousmenu)
			if sousmenu == 2 {
				break
			}
			// changer la langue mais y'en a qu'une
			Logoontop()
			fmt.Println("Choisir la langue :")
			fmt.Println("1- Français")
			fmt.Println("2- Quitter")
			fmt.Println("\n\n\n\n\n\n")
			fmt.Scanln(&sousmenu)
			break
		case 3:
			Logoontop() // ouvrir l'inventaire
			fmt.Println("Pourquoi tu veux voir l'inventaire, t'as même pas commencer le jeu, 0/20")
			fmt.Println("Et pour info y'a rien dans l'inventaire gros nullos")
			fmt.Println("1- Quitter")
			fmt.Println("\n\n\n\n\n\n")
			fmt.Scanln(&sousmenu)
			break

		case 4:
			Logoontop()
			fmt.Println("Fait par :")
			fmt.Println("1- Oscar V")
			fmt.Println("2- Florent F")
			fmt.Println("3- Enzo D-R")
			fmt.Println("Appuyer sur une touche pour continuer")
			fmt.Println("\n\n\n\n")
			fmt.Scanln(&sousmenu)
		case 5:
			Clear()
			return // fermer prg
		}
	}
}