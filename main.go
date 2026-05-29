package main

import (
	"fmt"
	"slices"
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
	slices.IndexFunc(c.Produits, func(p Produit) bool {
		return p.Id == id
	})
	return nil, nil
}

func (c *Catalogue) TrouverParCategorie(cat string) ([]Produit, error) {
	slices.IndexFunc(c.Produits, func(p Produit) bool {
		return p.Categorie == cat
	})
	return nil, nil
}

func (c *Catalogue) AppliquerReduction(categorie string, pct float64) int {
	count := 0
	for i := range c.Produits {
		if c.Produits[i].Categorie == categorie {
			c.Produits[i].Prix = c.Produits[i].Prix * (1 - pct/100)
			count++
		}
	}
	return count
}

func (c *Catalogue) Vendre(id int, qte int) error {
	p, err := c.TrouverParID(id)
	if err != nil {
		fmt.Println("Produit non trouvé")
		return err
	}
	if p.Stock < qte {
		fmt.Println("Stock insuffisant")
		return nil
	}
	p.Stock -= qte
	return nil
}

func (c *Catalogue) Rapport() string {
	nbProduits := len(c.Produits)
	valeurStock := 0

	for i := 0; i < nbProduits; i++ {
		valeurStock += c.Produits[i].Stock * int(c.Produits[i].Prix)
	}

	return "Catalogue: " + fmt.Sprintf("%d produits, valeur totale du stock: %.2f", nbProduits, float64(valeurStock))
}

func main() {
	//test
}
