package CombatGobelin

import (
	"fmt"
	"jeu/inventaire"
	"jeu/logo"
	"jeu/test"
)

var n = inventaire.Inv

// Structure du Gobelin
type Gobelin struct {
	Name          string
	CurrentHealth int
	AttackPoints  int
}

// Initialisation du Gobelin
func InitCombatGobelin() Gobelin {
	return Gobelin{
		Name:          "Erwann Kervoelen",
		CurrentHealth: 300,
		AttackPoints:  25,
	}
}

func CombatInteractif() {
	var buff string
	c := test.Player
	if c == nil {
		fmt.Println("Erreur : aucun personnage trouvé. Créez un personnage avant de lancer le combat.")
		return
	}

	gobelin := InitCombatGobelin()
	logo.Logoontop()
	fmt.Printf("\033[1m🔥 Combat contre %s lancé !\n\033[0m", gobelin.Name)
	fmt.Println("1- Continuez")
	fmt.Println("\n\n\n\n\n\n\n\n")
	fmt.Scanln(&buff)

	for c.HP > 0 && gobelin.CurrentHealth > 0 {
		logo.Logoontop()
		fmt.Printf("\n--- Tour du joueur ---\n")
		fmt.Printf("1. Attaquer\n2. Utiliser un objet\nChoix : ")
		fmt.Println("\n\n\n\n\n\n\n\n")
		var choix int
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			gobelin.CurrentHealth -= c.Power
			logo.Logoontop()
			fmt.Printf("%s attaque %s et inflige %d dégâts !\n", c.Name, gobelin.Name, c.Power)
			fmt.Println("1- Continuez")
			fmt.Println("\n\n\n\n\n\n\n\n")
			fmt.Scanln(&buff)

		case 2:
			if len(*n) > 0 { // On déréférence le pointeur pour avoir la map
				fmt.Println("Inventaire :")
				items := make([]string, 0, len(*n))
				i := 1
				for item := range *n {
					fmt.Printf("%d. %s\n", i, item)
					items = append(items, item)
					i++
				}

				fmt.Print("Choisissez un objet à utiliser : ")
				var itemChoice int
				fmt.Scanln(&itemChoice)

				if itemChoice >= 1 && itemChoice <= len(items) {
					item := items[itemChoice-1]
					fmt.Printf("%s utilise %s ! Effet : +20 PV\n", c.Name, item)
					c.HP += 20
					if c.HP > 120 {
						c.HP = 120
					}

					// Retirer l'item de l'inventaire global
					inventaire.RemoveInventory(*n, item, 1)
				} else {
					fmt.Println("Choix invalide.")
				}
			} else {
				fmt.Println("Inventaire vide.")
			}

		default:
			fmt.Println("Action inconnue.")
		}

		if gobelin.CurrentHealth <= 0 {
			logo.Logoontop()
			fmt.Printf("%s est vaincu ! 🎉\n", gobelin.Name)
			xpGain := 50
			c.Experience += xpGain
			fmt.Printf("%s gagne %d XP !\n", c.Name, xpGain)

			for c.Experience >= c.NextLevelExp {
				c.Level++
				c.Experience -= c.NextLevelExp
				c.NextLevelExp += 50
				c.HP += 20
				c.Power += 10
				fmt.Printf("⬆️ %s monte au niveau %d !\n", c.Name, c.Level)
			}
			fmt.Println("1- Continuez")
			fmt.Println("\n\n\n\n\n\n")
			fmt.Scanln(&buff)

			return
		}
		logo.Logoontop()
		fmt.Printf("\n--- Tour du Gobelin ---\n")
		damage := gobelin.AttackPoints
		if gobelin.CurrentHealth < 30 {
			damage *= 2
			fmt.Println("⚠️ Le Gobelin entre en RAGE !")
		}
		c.HP -= damage
		if c.HP < 0 {
			c.HP = 0
		}
		fmt.Printf("%s inflige %d dégâts à %s\n", gobelin.Name, damage, c.Name)
		fmt.Printf("PV restants : %d\n", c.HP)
		fmt.Println("1- Continuez")
		fmt.Println("\n\n\n\n\n\n\n")
		fmt.Scanln(&buff)

		if c.HP == 0 {
			logo.Logoontop()
			fmt.Printf("%s est vaincu... 💀\n", c.Name)
			fmt.Println("Bravo ! \n1- Continuez")
			fmt.Println("\n\n\n\n\n\n\n")
			fmt.Scanln(&buff)

			break
		}
	}
}
