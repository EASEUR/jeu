package jeu

import (
	"fmt"
	"strings"
	"unicode"
)

// Structure représentant un personnage du jeu
type Character struct {
	Name        string
	Class       string
	Level       int
	HP          int
	Power       int
	CryptoCoins int
	Inventory   []string
	Skills      []string
}

// Vérifie que le nom contient uniquement des lettres
func isValidName(name string) bool {
	for _, r := range name {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

// Formate le nom : première lettre en majuscule, le reste en minuscules
func formatName(name string) string {
	name = strings.ToLower(name)
	return strings.Title(name)
}

// Initialise le personnage selon sa classe
func Init(name, class string) Character {
	var hp, power, cryptoCoins int
	var inventory, skills []string

	switch class {
	case "Hacker":
		hp = 80
		power = 60
		cryptoCoins = 100
		inventory = []string{"Laptop", "VPN", "Exploit Kit"}
		skills = []string{"Piratage", "Analyse réseau", "Injection SQL"}
	case "Mercenaire":
		hp = 120
		power = 80
		cryptoCoins = 50
		inventory = []string{"Fusil", "Armure légère", "Grenade"}
		skills = []string{"Combat rapproché", "Tir de précision", "Intimidation"}
	case "Cyberdoc":
		hp = 100
		power = 40
		cryptoCoins = 70
		inventory = []string{"Kit médical", "Neuro-scanner", "Stimulant"}
		skills = []string{"Soins", "Implants", "Analyse biologique"}
	case "Ghost":
		hp = 90
		power = 70
		cryptoCoins = 80
		inventory = []string{"Camouflage", "Drone espion", "Silencieux"}
		skills = []string{"Infiltration", "Discrétion", "Sabotage"}
	default:
		hp = 100
		power = 50
		cryptoCoins = 20
		inventory = []string{}
		skills = []string{}
	}

	return Character{
		Name:        name,
		Class:       class,
		Level:       1,
		HP:          hp,
		Power:       power,
		CryptoCoins: cryptoCoins,
		Inventory:   inventory,
		Skills:      skills,
	}
}

// Création du personnage avec choix de classe
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
	fmt.Printf("\n--- Détails du personnage ---\n")
	fmt.Printf("Nom          : %s\n", character.Name)
	fmt.Printf("Classe       : %s\n", character.Class)
	fmt.Printf("Niveau       : %d\n", character.Level)
	fmt.Printf("HP           : %d\n", character.HP)
	fmt.Printf("Puissance    : %d\n", character.Power)
	fmt.Printf("CryptoCoins  : %d\n", character.CryptoCoins)
	fmt.Printf("Inventaire   : %v\n", character.Inventory)
	fmt.Printf("Compétences  : %v\n", character.Skills)

	return character
}

// Objets disponibles chez le marchand selon la classe
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

// Interaction avec le marchand : achat d’objets
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

// Point d’entrée du jeu
func main() {
	character := CharacterCreation()
	MarchandInteraction(&character)
}
