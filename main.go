package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Produit struct {
	Id        int
	Nom       string
	Marque    string
	Prix      float64
	Stock     int
	Categorie string
	Actif     bool
}

type Catalogue struct {
	Produits []Produit
}

func (c *Catalogue) AjouterProduit(p Produit) error {
	c.Produits = append(c.Produits, p)
	return nil
}

func (c *Catalogue) TrouverParID(id int) (*Produit, error) {
	produitIndex := slices.IndexFunc(c.Produits, func(prod Produit) bool { return prod.Id == id })
	if produitIndex == -1 {
		return nil, fmt.Errorf("produit avec l'ID %d introuvable", id)
	}
	return &c.Produits[produitIndex], nil
}

func (c *Catalogue) TrouverParCategorie(cat string) []Produit {
	var _ [][]Produit
	var produitsTrouves []Produit

	for _, p := range c.Produits {
		if strings.EqualFold(p.Categorie, cat) {
			produitsTrouves = append(produitsTrouves, p)
		}
	}
	return produitsTrouves
}

func (c *Catalogue) AppliquerReduction(categorie string, pct float64) int {
	count := 0
	for i := range c.Produits {
		if strings.EqualFold(c.Produits[i].Categorie, categorie) {
			c.Produits[i].Prix = c.Produits[i].Prix * (1 - pct/100)
			count++
		}
	}
	return count
}

func (c *Catalogue) Vendre(id int, qte int) error {
	p, err := c.TrouverParID(id)
	if err != nil {
		return err
	}
	if p.Stock < qte {
		return fmt.Errorf("stock insuffisant pour le produit %s (dispo: %d, demandé: %d)", p.Nom, p.Stock, qte)
	}
	p.Stock -= qte
	return nil
}

func (c *Catalogue) Rapport() string {
	nbProduits := len(c.Produits)
	var valeurTotalStock float64

	for _, p := range c.Produits {
		valeurTotalStock += float64(p.Stock) * p.Prix
	}

	return fmt.Sprintf("=== RAPPORT ===\n"+
		"Nombre de produits : %d\n"+
		"Valeur totale du stock : %.2f €",
		nbProduits, valeurTotalStock)
}

func main() {
	cat := Catalogue{}

	_ = cat.AjouterProduit(Produit{Id: 2, Nom: "iPhone 17", Marque: "Apple", Prix: 1229.00, Stock: 25, Categorie: "Smartphone", Actif: true})
	_ = cat.AjouterProduit(Produit{Id: 3, Nom: "Galaxy S26", Marque: "Samsung", Prix: 1449.00, Stock: 15, Categorie: "Smartphone", Actif: true})
	_ = cat.AjouterProduit(Produit{Id: 1, Nom: "MacBook Pro M3", Marque: "Apple", Prix: 1999.99, Stock: 10, Categorie: "Ordinateur", Actif: true})
	_ = cat.AjouterProduit(Produit{Id: 4, Nom: "AirPods Max", Marque: "Apple", Prix: 579.99, Stock: 8, Categorie: "Accessoire", Actif: true})
	_ = cat.AjouterProduit(Produit{Id: 5, Nom: "Galaxy Watch", Marque: "Samsung", Prix: 279.0, Stock: 40, Categorie: "Accessoire", Actif: true})

	// bufio permet de regarder les entrées de l'utilisateur même si elles contiennent des espaces
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n--- Catalogue TechShop ---")
		fmt.Println("[1] Ajouter un produit")
		fmt.Println("[2] Chercher un produit (ID ou Catégorie)")
		fmt.Println("[3] Appliquer une réduction")
		fmt.Println("[4] Vendre un produit")
		fmt.Println("[5] Rapport global")
		fmt.Println("[0] Quitter")
		fmt.Print("Votre choix : ")

		var choix int
		var err error
		_, err = fmt.Scanln(&choix)
		if err != nil {
			fmt.Println("Veuillez entrer un nombre valide.")
			scanner.Scan()
			continue
		}

		switch choix {
		case 1:
			var p Produit
			//id auto généré à la main
			p.Id = len(cat.Produits) + 1
			fmt.Print("Nom : ")
			scanner.Scan()
			p.Nom = scanner.Text()
			fmt.Print("Marque : ")
			scanner.Scan()
			p.Marque = scanner.Text()
			fmt.Print("Prix : ")
			_, err = fmt.Scanln(&p.Prix)
			if err != nil {
				return
			}
			fmt.Print("Stock initial : ")
			_, err = fmt.Scanln(&p.Stock)
			if err != nil {
				return
			}
			fmt.Print("Catégorie : ")
			scanner.Scan()
			p.Categorie = scanner.Text()
			p.Actif = true

			err = cat.AjouterProduit(p)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Produit ajouté avec succès !")
			}
		case 2:
			fmt.Println("Chercher par : [1] ID | [2] Catégorie")
			var sousChoix int
			_, err := fmt.Scanln(&sousChoix)
			if err != nil {
				return
			}

			if sousChoix == 1 {
				var id int
				fmt.Print("Entrez l'ID : ")
				_, err := fmt.Scanln(&id)
				if err != nil {
					return
				}
				p, err := cat.TrouverParID(id)
				if err != nil {
					fmt.Println("Error:", err)
				} else {
					fmt.Printf("Trouvé : [%d] %s - %s | Prix: %.2f € | Stock: %d\n", p.Id, p.Marque, p.Nom, p.Prix, p.Stock)
				}
			} else if sousChoix == 2 {
				fmt.Print("Entrez la catégorie : ")
				scanner.Scan()
				rechercheCat := scanner.Text()
				produits := cat.TrouverParCategorie(rechercheCat)
				if len(produits) == 0 {
					fmt.Println("Aucun produit trouvé dans cette catégorie.")
				} else {
					fmt.Printf("%d produits trouvés :\n", len(produits))
					for _, p := range produits {
						fmt.Printf(" - [%d] %s %s (%.2f €) | Stock: %d\n", p.Id, p.Marque, p.Nom, p.Prix, p.Stock)
					}
				}
			}
		case 3:
			var catNom string
			var pct float64
			fmt.Print("Catégorie : ")
			scanner.Scan()
			catNom = scanner.Text()
			fmt.Print("Pourcentage de réduction (ex: 15) : ")
			_, err := fmt.Scanln(&pct)
			if err != nil {
				return
			}

			modifies := cat.AppliquerReduction(catNom, pct)
			fmt.Printf("Réduction appliquée ! %d produits modifiés.\n", modifies)
		case 4:
			var id, qte int
			fmt.Print("ID du produit vendu : ")
			_, err := fmt.Scanln(&id)
			if err != nil {
				return
			}
			fmt.Print("Quantité : ")
			_, err = fmt.Scanln(&qte)
			if err != nil {
				return
			}

			err = cat.Vendre(id, qte)
			if err != nil {
				fmt.Println("Erreur de vente :", err)
			} else {
				fmt.Println("Vente enregistrée avec succès !")
			}
		case 5:
			fmt.Println(cat.Rapport())
		case 0:
			fmt.Println("Fermeture de l'application.")
			return
		default:
			fmt.Println("Option invalide, veuillez réessayer.")
		}
	}
}
