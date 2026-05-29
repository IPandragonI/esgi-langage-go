package main

import (
	"fmt"
)

func main() {
	const prenomConst = "Mathys"
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

	imc := poids / (taille * taille)
	fmt.Printf("\nBonjour %s, votre IMC est de : %.2f\n", prenomConst, imc)

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
