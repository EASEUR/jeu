package jeu

import "fmt"

type Character struct {
	Name      string
	Class     string
	Level     int
	HP        int
	Power     int
	Gold      int
	Inventory []string
}

func CreateCharacter() Character {
	fmt.Println("Bienvenue dans Xtrack")
	fmt.Print("Entrez le nom de votre personnage :")
	var name string
	fmt.Scanln(&name)

	fmt.Print("Choisissez une classe : Hacker, Mercenaire, Cyberdoc, Ghost")
	var class int
	fmt.Scanln(&class)

	fmt.Println("Personnage créé avec succès !")
	var className string
	switch class {
	case 1:
		className = "Hacker"
	case 2:
		className = "Mercenaire"
	case 3:
		className = "Cyberdoc"
	case 4:
		className = "Ghost"
	default:
		className = "Inconnu"
	}
	char := Character{
		Name:      name,
		Class:     className,
		Level:     1,
		HP:        100,
		Power:     50,
		Gold:      20,
		Inventory: []string{},
	}
	return char
}
