package cmd

import (
	"fmt"
	"log"

	"github.com/connorwade/quickstart/internal"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(drizzleCmd)
}

var drizzleCmd = &cobra.Command{
	Use:   "drizzle",
	Short: "Add initial drizzle setup to your project",
	Long:  `Share starter code for Svelte + Drizzle projects.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Creating Drizzle setup...")
		err := internal.ConfigureDrizzleProject()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("Drizzle setup complete!")
	},
}
