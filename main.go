package main

import (
	"fmt"
)

func main() {
	var a, b float64
	var op string

	fmt.Println("--- Calculatrice ---")

	// Boucle 'while true'
	for {
		fmt.Print("\nEntrez un calcul (Synthaxe : a, opération, b séparées par un ENTRER) : ")

		_, err := fmt.Scan(&a, &op, &b)
		if err != nil {
			fmt.Println("Erreur de lecture :", err)
			return
		}

		resultat, err := operer(a, b, op)
		if err != nil {
			fmt.Printf("Erreur : %v\n", err)
			continue
		}

		fmt.Printf("Résultat (via operer) : %.2f\n", resultat)
		fn := creerOperation(op)
		fmt.Printf("Résultat (via closure) : %.2f\n", fn(a, b))
	}
}

// Fonction operer
func operer(a, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("division par 0 impossible")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("opérateur inconnu '%s'", op)
	}
}

// Fonction creerOperation
func creerOperation(op string) func(float64, float64) float64 {
	switch op {
	case "+":
		return func(a, b float64) float64 { return a + b }
	case "-":
		return func(a, b float64) float64 { return a - b }
	case "*":
		return func(a, b float64) float64 { return a * b }
	case "/":
		return func(a, b float64) float64 {
			if b == 0 {
				panic("division par 0")
			}
			return a / b
		}
	default:
		panic("opérateur inconnu")
	}
}
