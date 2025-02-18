package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
)

type Expense struct {
	ID         string
	Name       string
	Amount     float64
	CategoryID string
}

type ExpenseCategory struct {
	ID         string
	Name       string
	ExpenseIDs []string
}

var (
	expenses   []Expense
	categories []ExpenseCategory
)

// TODO: implementar mejoras esto lo hizo la ia xd
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Comando requerido: add, list, delete, summary")
		return
	}

	switch os.Args[1] {
	case "add":
		addCmd := flag.NewFlagSet("add", flag.ExitOnError)
		addType := addCmd.String("type", "", "Tipo a añadir (expense/category)")
		name := addCmd.String("name", "", "Nombre del elemento")
		amount := addCmd.Float64("amount", 0, "Monto del gasto")
		categoryID := addCmd.String("category", "", "ID de categoría")

		addCmd.Parse(os.Args[2:])
		handleAdd(*addType, *name, *amount, *categoryID)

	case "list":
		listCmd := flag.NewFlagSet("list", flag.ExitOnError)
		listType := listCmd.String("type", "", "Tipo a listar (expenses/categories)")
		listCmd.Parse(os.Args[2:])
		handleList(*listType)

	case "delete":
		deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
		id := deleteCmd.String("id", "", "ID a eliminar")
		deleteCmd.Parse(os.Args[2:])
		handleDelete(*id)

	case "summary":
		handleSummary()

	default:
		fmt.Println("Comando no reconocido")
	}
}

func handleAdd(addType, name string, amount float64, categoryID string) {
	switch addType {
	case "expense":
		if name == "" || amount <= 0 || categoryID == "" {
			log.Fatal("Faltan parámetros para añadir gasto")
		}

		newExpense := Expense{
			ID:         uuid.New().String(),
			Name:       name,
			Amount:     amount,
			CategoryID: categoryID,
		}
		expenses = append(expenses, newExpense)
		fmt.Printf("Gasto añadido: %+v\n", newExpense)

	case "category":
		if name == "" {
			log.Fatal("Nombre de categoría requerido")
		}

		newCategory := ExpenseCategory{
			ID:   uuid.New().String(),
			Name: name,
		}
		categories = append(categories, newCategory)
		fmt.Printf("Categoría creada: %+v\n", newCategory)

	default:
		log.Fatal("Tipo no válido para añadir")
	}
}

func handleList(listType string) {
	switch listType {
	case "expenses":
		fmt.Println("Lista de gastos:")
		for _, e := range expenses {
			fmt.Printf("ID: %s, Nombre: %s, Monto: %.2f, Categoría: %s\n",
				e.ID, e.Name, e.Amount, e.CategoryID)
		}

	case "categories":
		fmt.Println("Lista de categorías:")
		for _, c := range categories {
			fmt.Printf("ID: %s, Nombre: %s\n", c.ID, c.Name)
		}

	default:
		log.Fatal("Tipo de lista no válido")
	}
}

func handleDelete(id string) {
	// Implementar lógica para eliminar gasto/categoría
	// y relaciones ONE2MANY/MANY2ONE
	fmt.Println("Eliminando elemento con ID:", id)
}

func handleSummary() {
	total := 0.0
	for _, e := range expenses {
		total += e.Amount
	}
	fmt.Printf("Resumen total: %.2f\n", total)
}
