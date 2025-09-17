package mapvg

import (
	"fmt"
	"jeu/logo"
)

func piece1() bool {
	var buff string
	logo.Logoontop()
	fmt.Println("Vous avez trouver la 1ere piece gg wp")
	fmt.Println("1- Continuez")
	fmt.Println("\n\n\n\n\n\n\n\n")
	fmt.Scanln(&buff)
	return true
}
func piece2() bool {
	var buff string
	logo.Logoontop()
	fmt.Println("Vous avez trouver la 2eme piece gg wp")
	fmt.Println("1- Continuez")
	fmt.Println("\n\n\n\n\n\n\n\n")
	fmt.Scanln(&buff)
	return true
}
func piece3() bool {
	var buff string
	logo.Logoontop()
	fmt.Println("Vous avez trouver la 3eme piece gg wp")
	fmt.Println("1- Continuez")
	fmt.Println("\n\n\n\n\n\n\n\n")
	fmt.Scanln(&buff)
	return true
}

func Mapmobile() {
	var choix string
	var buff string
	var piecen1 bool
	var piecen2 bool
	var piecen3 bool

	// var colonne [5]int // création de la map
	// var ligne [5]int
	i := 2 // hauter
	j := 0 // largeur
	for {
		logo.Logoontop()
		fmt.Printf("Vous vous trouvez actuellement à la case '%d' (hauteur) ; '%d' (largeur)\nQuelle action voulez-vous faire :\n", i, j)
		if i == 5 {
			fmt.Println("Vous ne pouvez pas aller en haut, vous êtes au max de la hauteur")
		}
		if j == 5 {
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
		if i < 5 {
			fmt.Println("\033[33mZ- Aller en haut\033[0m")
		}
		if j > 0 {
			fmt.Println("\033[33mQ- Aller à gauche\033[0m")
		}
		if j < 5 {
			fmt.Println("\033[33mD- Aller à droite\033[0m")
		}
		fmt.Println("\033[33mI- Iventaire\033[0m")
		fmt.Println("\n\n\n\n")
		fmt.Scanln(&choix)
		switch choix {
		case "Z":
			if i == 5 {
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
			if j == 5 {
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
			// ouvrir inventaire
		}
		if i == 0 && j == 0 && !piecen1 {
			piecen1 = piece1()
		}
		if i == 2 && j == 2 && !piecen2 {
			piecen2 = piece2()
		}
		if i == 5 && j == 5 && !piecen3 {
			piecen3 = piece3()
		}

	}
}
