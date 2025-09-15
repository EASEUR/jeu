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
func menu() {
	var choixmenu int
	clear()
	fmt.Println("\n\n\n\n\n")
	fmt.Println(` /$$   /$$  /$$                                  /$$      
| $$  / $$ | $$                                 | $$      
|  $$/ $$//$$$$$$    /$$$$$$  /$$$$$$   /$$$$$$$| $$   /$$
 \  $$$$/|_  $$_/   /$$__  $$|____  $$ /$$_____/| $$  /$$/
  >$$  $$  | $$    | $$  \__/ /$$$$$$$| $$      | $$$$$$/ 
 /$$/\  $$ | $$ /$$| $$      /$$__  $$| $$      | $$_  $$ 
| $$  \ $$ |  $$$$/| $$     |  $$$$$$$|  $$$$$$$| $$ \  $$ 
|__/  |__/  \___/  |__/      \_______/ \_______/|__/  \__/`)
	fmt.Println("\n\n\n")
	fmt.Println("Bienvenu sur Xtrack :") // deroulement du menu
	fmt.Println("Choisissez un des menus ci dessous :")
	fmt.Println("1- Jouons !")
	fmt.Println("2- Paramètres")
	fmt.Println("3- Ouvrir l'inventaire")
	fmt.Println("4- Quittez")
	fmt.Println("\n\n\n\n")
	fmt.Scanln(choixmenu) // choix du menu

	switch choixmenu {
	case 1:
		//commencer a jouer
	case 2:
		// ouverture des parametres
	case 3:
		// ouvrir l'inventaire
	case 4:
		// quitte le menu ?
	}
}
