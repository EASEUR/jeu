package jeu

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
