package dictionary

import (
	"exemple1/dictionary"
	"fmt"
)

func main() {
	// Utilisation de la classe Dictionnaire
	dict := dictionary.NewDictionnaire()

	dict.Ajouter("estiam", "ecole")
	dict.Ajouter("HAROUN", "etudiant")

	fmt.Println("Get 'estiam':", dict.Get("estiam"))

	dict.Lister()

	dict.Supprimer("HAROUN")
	dict.Lister()
}
