package jeu

import "fmt"

// Structure représentant un personnage du jeu
type Character struct {
	Name      string   // Nom du personnage
	Class     string   // Classe choisie
	Level     int      // Niveau du personnage
	HP        int      // Points de vie actuels
	Power     int      // Puissance d'attaque ou compétence
	Gold      int      // Quantité d'or possédée
	Inventory []string // Inventaire contenant les objets
}

// Fonction pour créer un nouveau personnage via saisie utilisateur
func CreateCharacter() Character {
	fmt.Println("Bienvenue dans Xtrack")

	fmt.Print("Entrez le nom de votre personnage :")
	var name string
	fmt.Scanln(&name)

	fmt.Print("Choisissez une classe : Hacker (1), Mercenaire (2), Cyberdoc (3), Ghost (4) : ")
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

// Nouvelle fonction pour initialiser un personnage avec des valeurs précises
func InitCharacter(name string, class string, level int, hp int, inventory []string) Character {
	return Character{
		Name:      name,
		Class:     class,
		Level:     level,
		HP:        hp,
		Power:     50,
		Gold:      20,
		Inventory: inventory,
	}
}
