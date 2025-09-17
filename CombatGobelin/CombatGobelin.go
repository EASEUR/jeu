package CombatGobelin

import (
	"fmt"
	"jeu/test"
)

// Structure du Gobelin
type Gobelin struct {
	Name          string
	CurrentHealth int
	AttackPoints  int
}

// Initialisation du Gobelin
func InitCombatGobelin() Gobelin {
	return Gobelin{
		Name:          "Zero.exe",
		CurrentHealth: 100,
		AttackPoints:  15,
	}
}

// Func du combat entre le joueur et le monstre. Elle met en place un combat Tour par Tour avec des possbilités d'utiliser les compétences, les attaques basique etc
func CombatInteractif() {
	c := test.Player
	gobelin := InitCombatGobelin()
	fmt.Printf("\n🔥 Combat contre %s lancé !\n", gobelin.Name)

	for c.HP > 0 && gobelin.CurrentHealth > 0 {
		fmt.Printf("\n--- Tour du joueur ---\n")
		fmt.Printf("1. Attaquer\n2. Utiliser une compétence\n3. Utiliser un objet\nChoix : ")
		var choix int
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			gobelin.CurrentHealth -= c.Power
			fmt.Printf("%s attaque %s et inflige %d dégâts !\n", c.Name, gobelin.Name, c.Power)
		case 2:
			if len(c.Skills) > 0 {
				fmt.Println("Compétences disponibles :")
				for i, skill := range c.Skills {
					fmt.Printf("%d. %s\n", i+1, skill)
				}
				fmt.Print("Choisissez une compétence : ")
				var skillChoice int
				fmt.Scanln(&skillChoice)
				if skillChoice >= 1 && skillChoice <= len(c.Skills) {
					fmt.Printf("%s utilise %s ! Effet : +10 dégâts\n", c.Name, c.Skills[skillChoice-1])
					gobelin.CurrentHealth -= c.Power + 10
				} else {
					fmt.Println("Choix invalide.")
				}
			} else {
				fmt.Println("Aucune compétence disponible.")
			}
		case 3:
			if len(c.Inventory) > 0 {
				fmt.Println("Inventaire :")
				for i, item := range c.Inventory {
					fmt.Printf("%d. %s\n", i+1, item)
				}
				fmt.Print("Choisissez un objet à utiliser : ")
				var itemChoice int
				fmt.Scanln(&itemChoice)
				if itemChoice >= 1 && itemChoice <= len(c.Inventory) {
					item := c.Inventory[itemChoice-1]
					fmt.Printf("%s utilise %s ! Effet : +20 PV\n", c.Name, item)
					c.HP += 20
					if c.HP > 120 {
						c.HP = 120
					}
				} else {
					fmt.Println("Choix invalide.")
				}
			} else {
				fmt.Println("Inventaire vide.")
			}
		default:
			fmt.Println("Action inconnue.")
		}

		// Affiche le nombre d'xp gagnée
		if gobelin.CurrentHealth <= 0 {
			fmt.Printf("%s est vaincu ! 🎉\n", gobelin.Name)
			xpGain := 50
			c.Experience += xpGain
			fmt.Printf("%s gagne %d XP !\n", c.Name, xpGain)

			// ⬆️ Vérification de montée de niveau : affiche le niveau actuel et lorsque le joueur gagne ou monte de niveau, stats...
			for c.Experience >= c.NextLevelExp {
				c.Level++
				c.Experience -= c.NextLevelExp
				c.NextLevelExp += 50
				c.HP += 20
				c.Power += 10
				fmt.Printf("⬆️ %s monte au niveau %d !\n", c.Name, c.Level)
				fmt.Printf("💪 PV augmentés à %d, Puissance à %d\n", c.HP, c.Power)
			}
			break
		}

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

		if c.HP == 0 {
			fmt.Printf("%s est vaincu... 💀\n", c.Name)
			break
		}
	}
}
