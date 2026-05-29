package main

import "fmt"

type Severity int

// Utilisation de iota
const (
	StatusInfo Severity = iota
	StatusWarning
	StatusCritical
)

func main() {
	// Contexte exercice : Gestion d'alertes

	// Utilisation de slice
	alertesRecues := []Severity{StatusInfo, StatusCritical, StatusWarning, StatusCritical}

	// Utilisation de for
	for index, severite := range alertesRecues {
		fmt.Printf("\n[Alerte #%d - Code ID %d] Actions requises : ", index+1, severite)

		// Utilisation de switch avec fallthrough
		switch severite {
		case StatusCritical:
			fmt.Print("CALL_EMERGENCY_SERVICES -> ")
			fallthrough

		case StatusWarning:
			fmt.Print("SEND_MAIL -> ")
			fallthrough

		case StatusInfo:
			fmt.Print("WRITE_TO_LOGFILE")

		default:
			fmt.Print("IGNORE")
		}

		fmt.Println()
	}
}
