package test

import (
	"fmt"
	"jeu/logo"
	"strings"
	"unicode"
)

// Structure qui représente le personne créé
type Character struct {
	Name         string
	Class        string
	Level        int
	HP           int
	Power        int
	CryptoCoins  int
	Inventory    []string
	Skills       []string
	Experience   int
	NextLevelExp int
}

// Cette func vérifie si les nom choisie contient suelement des lettres
func isValidName(name string) bool {
	for _, r := range name {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

// Cette func permet obligatoirement de mettre le nom crée avec la première lettre en majuscule et le reste en minuscule. ex : Ghost
func formatName(name string) string {
	name = strings.ToLower(name)
	return strings.Title(name)
}

// Initialise le personnage selon sa classe, ses compétences et son inventaire
func Init(name, class string) Character {
	var hp, power, cryptoCoins int
	var inventory, skills []string

	switch class {
	case "Hacker":
		hp = 80
		power = 60
		cryptoCoins = 100
	case "Mercenaire":
		hp = 120
		power = 80
		cryptoCoins = 50
	case "Cyberdoc":
		hp = 100
		power = 40
		cryptoCoins = 70
	case "Ghost":
		hp = 90
		power = 70
		cryptoCoins = 80
	default:
		hp = 100
		power = 50
		cryptoCoins = 20
		inventory = []string{}
		skills = []string{}
	}

	return Character{
		Name:         name,
		Class:        class,
		Level:        1,
		HP:           hp,
		Power:        power,
		CryptoCoins:  cryptoCoins,
		Inventory:    inventory,
		Skills:       skills,
		Experience:   0,
		NextLevelExp: 100,
	}
}

// Cette func permet de créer son personnage et de choisir sa classe pour l'Aventure
func CharacterCreation() Character {

	var buffer int
	var name string
	for {
		logo.Logoontop()
		fmt.Print("Entrez le nom de votre personnage : ")
		fmt.Print("\n\n\n\n\n\n\n\n\n")
		fmt.Scanln(&name)
		if isValidName(name) {
			break
		}
		fmt.Println("Erreur : Le nom ne doit contenir que des lettres.")
	}
	name = formatName(name)
	logo.Logoontop()
	fmt.Print("Choisissez une classe : Hacker (1), Mercenaire (2), Cyberdoc (3), Ghost (4) : ")
	fmt.Print("\n\n\n\n\n\n\n\n\n")
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
	logo.Logoontop()
	fmt.Println("Personnage créé avec succès !")
	fmt.Printf("\n--- Détails du personnage ---\n")
	fmt.Printf("Nom          : %s\n", character.Name)
	fmt.Printf("Classe       : %s\n", character.Class)
	fmt.Printf("Niveau       : %d\n", character.Level)
	fmt.Printf("HP           : %d\n", character.HP)
	fmt.Printf("Puissance    : %d\n", character.Power)
	fmt.Printf("CryptoCoins  : %d\n", character.CryptoCoins)
	fmt.Printf("Inventaire   : %v\n", character.Inventory)
	fmt.Printf("Compétences  : %v\n", character.Skills)
	fmt.Printf("XP           : %d/%d\n", character.Experience, character.NextLevelExp)
	fmt.Println("1- Continuer")
	fmt.Scanln(&buffer)
	return character
}

// Cette fonction permet d'afficher les compétences propre à chaque classe
func GetMarchandItems(class string) map[string]int {
	switch class {
	case "Hacker":
		return map[string]int{
			"Module de Brèche Quantique": 15,
			"Injecteur de Code Fantôme":  10,
			"Firewall Portatif":          8,
			"Clé USB Crypter-X":          5,
		}
	case "Mercenaire":
		return map[string]int{
			"Implant de Visée Thermique":    20,
			"Chargeur Plasma Haute Densité": 12,
			"Armure Tactique Légère":        18,
			"Grenade EMP":                   7,
		}
	case "Cyberdoc":
		return map[string]int{
			"Stimulant Neuro-Régénérant":  10,
			"Kit de Suture NanoTech":      6,
			"Implant Cortex Bio-Analyse":  22,
			"Spray Anti-Virus Cellulaire": 4,
		}
	case "Ghost":
		return map[string]int{
			"Camouflage Optique V3":       25,
			"Drone Silencieux Arachnide":  18,
			"Saboteur de Signal":          9,
			"Injecteur de Silence Neural": 6,
		}
	default:
		return map[string]int{}
	}
}

// Cette func permet d'intéragir avec le marchand et d'acheter les items
func MarchandInteraction(c *Character) {
	items := GetMarchandItems(c.Class)

	fmt.Println("\n--- Marchand Cyberpunk ---")
	fmt.Printf("Objets disponibles pour la classe %s :\n", c.Class)
	i := 1
	itemList := []string{}
	for item, price := range items {
		fmt.Printf("%d. %s (%d pièces d’or)\n", i, item, price)
		itemList = append(itemList, item)
		i++
	}

	fmt.Print("Entrez le numéro de l’objet à acheter (0 pour quitter) : ")
	var choice int
	fmt.Scanln(&choice)

	if choice == 0 {
		fmt.Println("Vous quittez le marchand.")
		return
	}

	if choice < 1 || choice > len(itemList) {
		fmt.Println("Choix invalide.")
		return
	}

	selectedItem := itemList[choice-1]
	cost := items[selectedItem]

	if c.CryptoCoins < cost {
		fmt.Printf("Pas assez de pièces d’or pour %s.\n", selectedItem)
		return
	}

	c.CryptoCoins -= cost
	c.Inventory = append(c.Inventory, selectedItem)
	fmt.Printf("Achat réussi : %s ajouté à l’inventaire.\n", selectedItem)
	fmt.Printf("CryptoCoins restants : %d\n", c.CryptoCoins)
}
