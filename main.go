package main

import (
	"fmt"
)

type Personne struct {
	Prenom string
	Nom    string
	Age    int
	Email  string
}

func (p Personne) NomComplet() string {
	return p.Prenom + " " + p.Nom
}

func (p Personne) Presentation() string {
	return fmt.Sprintf("%s %s (%d ans) - %s", p.Prenom, p.Nom, p.Age, p.Email)
}

type Adresse struct {
	Rue        string
	Ville      string
	CodePostal string
}

func (a Adresse) Format() string {
	return fmt.Sprintf("%s, %s %s", a.Rue, a.CodePostal, a.Ville)
}

type Employe struct {
	Personne
	Adresse
	Poste   string
	Salaire float64
}

func (e *Employe) FicheEmploye() string {
	return fmt.Sprintf("=== FICHE EMPLOYÉ ===\nNom: %s\nAge: %d\nEmail: %s\nAdresse: %s\nPoste: %s\nSalaire: %.2f €\n",
		e.NomComplet(),
		e.Age,
		e.Email,
		e.Format(),
		e.Poste,
		e.Salaire,
	)
}

func (e *Employe) AugmenterSalaire(pourcentage float64) {
	e.Salaire += e.Salaire * pourcentage / 100
}

type Etudiant struct {
	Personne
	Promo   string
	Moyenne float64
}

func (e Etudiant) FicheEtudiant() string {
	return fmt.Sprintf("=== FICHE ÉTUDIANT ===\nNom: %s\nAge: %d\nEmail: %s\nPromo: %s\nMoyenne: %.1f/20\n",
		e.NomComplet(),
		e.Age,
		e.Email,
		e.Promo,
		e.Moyenne,
	)
}

func (e Etudiant) MentionObtenue() string {
	switch {
	case e.Moyenne >= 16:
		return "Très Bien"
	case e.Moyenne >= 14:
		return "Bien"
	case e.Moyenne >= 12:
		return "Assez Bien"
	case e.Moyenne >= 10:
		return "Passable"
	default:
		return "Ajourné"
	}
}

func main() {
	employes := []Employe{
		{
			Personne: Personne{Prenom: "Novak", Nom: "Djokovic", Age: 36, Email: "nDjoko@mail.com"},
			Adresse:  Adresse{Rue: "Rue du Service", Ville: "Monaco", CodePostal: "98000"},
			Poste:    "Directeur Technique de l'Open D'australie",
			Salaire:  8500.00,
		},
		{
			Personne: Personne{Prenom: "Rafael", Nom: "Nadal", Age: 37, Email: "rNadal@mail.com"},
			Adresse:  Adresse{Rue: "Via Manacor", Ville: "Mallorca", CodePostal: "07500"},
			Poste:    "Directeur Opérationnel de Roland Garros",
			Salaire:  7200.00,
		},
	}

	etudiants := []Etudiant{
		{
			Personne: Personne{Prenom: "Roger", Nom: "Federer", Age: 42, Email: "rFederer@mail.com"},
			Promo:    "Master DevOps 2026",
			Moyenne:  17.5,
		},
		{
			Personne: Personne{Prenom: "Andy", Nom: "Murray", Age: 39, Email: "aMurray@mail.com"},
			Promo:    "Bachelor Cloud 2026",
			Moyenne:  11.5,
		},
	}

	// Test de l'augmentation de salaire
	fmt.Printf("\nSalaire de Novak avant augmentation : %.2f €\n", employes[0].Salaire)
	fmt.Println("(AugmenterSalaire(10))")
	employes[0].AugmenterSalaire(10)
	fmt.Printf("Salaire de Novak après augmentation : %.2f €\n\n", employes[0].Salaire)

	fmt.Print("--------------------------\n")

	//On affiche pour chaque employé sa fiche
	for _, emp := range employes {
		fmt.Println(emp.FicheEmploye() + "\n")
	}

	fmt.Print("--------------------------\n")

	//idem pour les étudiants
	for _, etu := range etudiants {
		fmt.Println(etu.FicheEtudiant())
		fmt.Printf("Mention obtenue : %s \n\n", etu.MentionObtenue())
	}
}
