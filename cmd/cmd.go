package cmd

import (
	"fmt"
	"os"

	"family-tree/database"
	"family-tree/models"
	"family-tree/operationsAndqueries"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "family-tree",
	Short: "A command-line tool for managing family trees.",
}

func init() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(connectCmd)
	rootCmd.AddCommand(countCmd)
	rootCmd.AddCommand(fatherCmd)
}

var addCmd = &cobra.Command{
	Use:   "add <entity> <name> [<name_surname>]",
	Short: "Add a new person or relationship to the family tree.",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		entity := args[0]
		name := args[1]

		db, err := database.InitDB()
		if err != nil {
			fmt.Printf("Failed to connect to the database: %v\n", err)
			os.Exit(1)
		}
		defer db.Close()

		switch entity {
		case "person":
			var fullName string
			if len(args) > 2 {
				fullName = name + " " + args[2]
			} else {
				fullName = name
			}

			// Check if the person already exists
			if operationsAndqueries.PersonExists(db, fullName) {
				fmt.Printf("Person '%s' already exists\n", fullName)
				os.Exit(1)
			}

			person := models.Person{Username: fullName}
			err := operationsAndqueries.AddPerson(db, &person)
			if err != nil {
				fmt.Printf("Failed to add person: %v\n", err)
				os.Exit(1)
			}
			fmt.Println("Person added successfully!")
		case "relationship":
			// Check if the relationship already exists
			if operationsAndqueries.RelationshipExists(db, name) {
				fmt.Printf("Relationship '%s' already exists\n", name)
				os.Exit(1)
			}

			relationship := models.Relationship{RelationshipType: name}
			err := operationsAndqueries.AddRelationship(db, &relationship)
			if err != nil {
				fmt.Printf("Failed to add relationship: %v\n", err)
				os.Exit(1)
			}
			fmt.Println("Relationship added successfully!")
		default:
			fmt.Println("Unknown entity:", entity)
			os.Exit(1)
		}
	},
}

var connectCmd = &cobra.Command{
	Use:   "connect <name1> <name1_surname> as <relationship> of <name2> <name2_surname>",
	Short: "Connect two persons with a specified relationship",
	Args:  cobra.MinimumNArgs(6),
	Run: func(cmd *cobra.Command, args []string) {
		name1 := args[0] + " " + args[1]
		relationship := args[3]
		name2 := args[5] + " " + args[6]

		db, err := database.InitDB()
		if err != nil {
			fmt.Printf("Failed to connect to the database: %v\n", err)
			os.Exit(1)
		}
		defer db.Close()

		err = operationsAndqueries.ConnectPersons(db, name1, relationship, name2)
		if err != nil {
			fmt.Printf("Failed to connect persons: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Persons connected successfully!")
	},
}

var countCmd = &cobra.Command{
	Use:   "count <entity> of <name> <name_surname>",
	Short: "Count the number of sons for a person.",
	Args:  cobra.ExactArgs(4), // Ensure exactly 4 arguments are provided
	Run: func(cmd *cobra.Command, args []string) {
		entity := args[0]
		name := args[2] + " " + args[3]

		db, err := database.InitDB()
		if err != nil {
			fmt.Printf("Failed to connect to the database: %v\n", err)
			os.Exit(1)
		}
		defer db.Close()

		var count int
		switch entity {
		case "sons":
			count, err = operationsAndqueries.CountSons(db, name)
		case "daughters":
			count, err = operationsAndqueries.CountDaughters(db, name)
		case "wives":
			count, err = operationsAndqueries.CountWives(db, name)
		// Add cases for other entities as needed
		default:
			fmt.Printf("Invalid entity: %s\n", entity)
			os.Exit(1)
		}

		if err != nil {
			fmt.Printf("Failed to count %s of %s: %v\n", entity, name, err)
			os.Exit(1)
		}
		fmt.Printf("Number of %s of %s: %d\n", entity, name, count)
	},
}

var fatherCmd = &cobra.Command{
	Use:   "father of <name> <name_surname>",
	Short: "Get the name of the father for a person.",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[1] + " " + args[2]

		db, err := database.InitDB()
		if err != nil {
			fmt.Printf("Failed to connect to the database: %v\n", err)
			os.Exit(1)
		}
		defer db.Close()

		fatherName, err := operationsAndqueries.FatherOf(db, name)
		if err != nil {
			fmt.Printf("Failed to get father of %s: %v\n", name, err)
			os.Exit(1)
		}
		fmt.Printf("Father of %s: %s\n", name, fatherName)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
