package jeu

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func clear() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else { // Linux, macOS
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
func logoontop() {
	clear()
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
func menu() {
	for true {
		var choixmenu int
		logoontop()
		fmt.Println("Bienvenu sur Xtrack :") // deroulement du menu
		fmt.Println("Choisissez un des menus ci dessous :")
		fmt.Println("1- Jouons !")
		fmt.Println("2- Paramètres")
		fmt.Println("3- Ouvrir l'inventaire")
		fmt.Println("4- Quittez")
		fmt.Println("\n\n\n\n")
		fmt.Scanln(&choixmenu) // choix du menu
		var sousmenu int
		switch choixmenu {
		case 1:
			//commencer a jouer
			logoontop()
			fmt.Println("Etes-vous sûr de bien vouloir rentrer dans le Souk ?") // verification
			fmt.Println("1- Oui !")
			fmt.Println("2- Non...")
			fmt.Println("\n\n\n\n\n\n")
			fmt.Scanln(&sousmenu) // choix sous menu
			switch sousmenu {
			case 1:
				logoontop()
				fmt.Println("en vrai relou parce que c'est pas fini, reviens dans le menu")
				fmt.Println("1- Oui")
				fmt.Println("2- Oui")
				fmt.Println("\n\n\n\n\n\n")
				fmt.Scanln(&sousmenu)
				break // sortir
			case 2:
				break // sortir
			}
		case 2: // ouverture des parametres
			logoontop()
			fmt.Println("Paramètres")
			fmt.Println("1- Langues")
			fmt.Println("2- Quitter")
			fmt.Println("\n\n\n\n\n\n")
			fmt.Scanln(&sousmenu)
			// changer la langue mais y'en a qu'une
			logoontop()
			fmt.Println("Choisir la langue :")
			fmt.Println("1- Français")
			fmt.Println("2- Quitter")
			fmt.Println("\n\n\n\n\n\n")
			fmt.Scanln(&sousmenu)
			break
		case 3:
			logoontop() // ouvrir l'inventaire
			fmt.Println("Pourquoi tu veux voir l'inventaire, t'as même pas commencer le jeu, 0/20")
			fmt.Println("Et pour info y'a rien dans l'inventaire gros nullos")
			fmt.Println("1- Quitter")
			fmt.Println("\n\n\n\n\n\n")
			fmt.Scanln(&sousmenu)
			break

		case 4:
			return // fermer prg
		}
	}
}
