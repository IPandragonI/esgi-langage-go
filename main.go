package main

import (
	"fmt"
)

func main() {
	const prenom = "Mathys"
	var poids float64
	var taille float64

	const (
		IMCMaigreur = 18.5
		IMCNormal   = 25.0
		IMCSurpoids = 30.0
	)

	// On demande à l'utilisateur de renseigner son poids et sa taille
	fmt.Print("Entrez votre poids (en kg, ex: 70.5) : ")
	_, err := fmt.Scanln(&poids)
	if err != nil {
		return
	}

	fmt.Print("Entrez votre taille (en m, ex: 1.75) : ")
	_, err = fmt.Scanln(&taille)
	if err != nil {
		return
	}

	imc := calculImc(poids, taille)

	fmt.Printf("\nBonjour %s, votre IMC est de : %.2f\n", prenom, imc)

	// On affiche la catégorie de poids en fonction de l'imc calculé
	switch {
	case imc < IMCMaigreur:
		fmt.Println("Catégorie : Maigreur")
	case imc < IMCNormal:
		fmt.Println("Catégorie : Normal")
	case imc < IMCSurpoids:
		fmt.Println("Catégorie : Surpoids")
	default:
		fmt.Println("Catégorie : Obésité")
	}
}

func calculImc(poids float64, taille float64) float64 {
	if taille == 0 {
		return 0
	}
	return poids / (taille * taille)
}
