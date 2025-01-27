//go:build !test

package main

import (
	"github.com/spf13/cobra"
	"log"
	"short-link/internal/adapter/storage/postgres"
)

var migrateCmd = &cobra.Command{Use: "migrate"}

func main() {
	if err := migrateCmd.Execute(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

const defaultMigrationStep = -1

// upCmd represents the migrate command
var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Run database migrations",
	Long:  `Run database migrations to update the database schema as per defined migration files.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := postgres.RunMigrations()
		if err != nil {
			return
		}
	},
}

// downCmd represents the migrate command
var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Revert the last database migration",
	Long:  `Revert the last database migration.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := postgres.RunDownMigration(step)
		if err != nil {
			return
		}
	},
}

var step int

func init() {
	downCmd.Flags().IntVarP(&step, "step", "s", defaultMigrationStep, "Number of migrations to revert")
	migrateCmd.AddCommand(upCmd, downCmd)
}
