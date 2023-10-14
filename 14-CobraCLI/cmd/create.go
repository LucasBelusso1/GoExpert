/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/LucasBelusso1/14-CobraCli/internal/database"
	"github.com/spf13/cobra"
)

func NewCreateCmd(categoryDb database.Category) *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "Create a new category",
		Long:  `Create a new category based on the flags "name" (-n or --name) and "description" (-d or --description)`,
		RunE:  RunCreate(categoryDb),
	}
}

func RunCreate(categoryDb database.Category) RunEFunc {
	return func(cmd *cobra.Command, args []string) error {
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")

		_, err := categoryDb.Create(name, description)

		if err != nil {
			return err
		}

		return nil
	}
}

func init() {
	createCmd := NewCreateCmd(GetCategoryDb(GetDb()))
	categoryCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("name", "n", "", "Name of the category")
	createCmd.Flags().StringP("description", "d", "", "Description of the category")
}
