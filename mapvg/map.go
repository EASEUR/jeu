package mapvg

import (
	"fmt"
	"jeu/logo"
)

func mapvg() {
	var colonne [5]int // création de la map
	var ligne [5]int
	i := 0
	j := 2
	for {
		logo.Logoontop()
		fmt.Println("Vous vous trouvez actuellement à la case '%d' ; '%d'", colonne[i], ligne[j])
		fmt.Println("Quelle action voulez-vous faire :")
		fmt.Println("1- Aller à droite") //rajouter le if pour les bordures de map
		fmt.Println("1- Aller à gauche") //rajouter le if pour les bordures de map
		fmt.Println("1- Aller en haut")  //rajouter le if pour les bordures de map
		fmt.Println("1- Aller en bas")   //rajouter le if pour les bordures de map
	}
}
