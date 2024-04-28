package cmd

import (
	"fmt"
	"log"

	"github.com/connorwade/quickstart/internal"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(svelteCmd)
}

var svelteCmd = &cobra.Command{
	Use:   "svelte",
	Short: "Add initial Svelte setup to your project",
	Long:  `Share starter code for Svelte projects.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Creating Svelte setup...")
		err := internal.ConfigureSvelteProject()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("Svelte setup complete!")
	},
}
