package jeu

import (
	"fmt"
	"strings"
	"unicode"
)

type Character struct {
	Name      string
	Class     string
	Level     int
	HP        int
	Power     int
	Crypto    int
	Inventory []string
	Skills    []string
}

// Valide que le nom contient uniquement des lettres
func isValidName(name string) bool {
	for _, r := range name {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

// Formate le nom : Majuscule + minuscules
func formatName(name string) string {
	name = strings.ToLower(name)
	return strings.Title(name)
}

// Initialise le personnage selon sa classe
func Init(name, class string) Character {
	var hp, power, crypto int
	var inventory, skills []string

	switch class {
	case "Hacker":
		hp = 80
		power = 60
		crypto = 100
		inventory = []string{"Laptop", "VPN", "Exploit Kit"}
		skills = []string{"Piratage", "Analyse réseau", "Injection SQL"}
	case "Mercenaire":
		hp = 120
		power = 80
		crypto = 50
		inventory = []string{"Fusil", "Armure légère", "Grenade"}
		skills = []string{"Combat rapproché", "Tir de précision", "Intimidation"}
	case "Cyberdoc":
		hp = 100
		power = 40
		crypto = 70
		inventory = []string{"Kit médical", "Neuro-scanner", "Stimulant"}
		skills = []string{"Soins", "Implants", "Analyse biologique"}
	case "Ghost":
		hp = 90
		power = 70
		crypto = 80
		inventory = []string{"Camouflage", "Drone espion", "Silencieux"}
		skills = []string{"Infiltration", "Discrétion", "Sabotage"}
	default:
		hp = 100
		power = 50
		crypto = 20
		inventory = []string{}
		skills = []string{}
	}

	return Character{
		Name:      name,
		Class:     class,
		Level:     1,
		HP:        hp,
		Power:     power,
		Crypto:    crypto,
		Inventory: inventory,
		Skills:    skills,
	}
}

// Fonction principale de création
func CharacterCreation() Character {
	fmt.Println("Bienvenue dans Xtrack")

	var name string
	for {
		fmt.Print("Entrez le nom de votre personnage : ")
		fmt.Scanln(&name)
		if isValidName(name) {
			break
		}
		fmt.Println("Erreur : Le nom ne doit contenir que des lettres.")
	}
	name = formatName(name)

	fmt.Print("Choisissez une classe : Hacker (1), Mercenaire (2), Cyberdoc (3), Ghost (4) : ")
	var classChoice int
	fmt.Scanln(&classChoice)

	var className string
	switch classChoice {
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

	character := Init(name, className)
	fmt.Println("Personnage créé avec succès !")
	fmt.Printf("Nom : %s\nClasse : %s\nHP : %d\nPower : %d\nCrypto : %d\nInventaire : %v\nCompétences : %v\n",
		character.Name, character.Class, character.HP, character.Power, character.Crypto, character.Inventory, character.Skills)

	return character
}
