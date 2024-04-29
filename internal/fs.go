package internal

import (
	"os"

	"github.com/connorwade/quickstart/files"
)

func CreateTailwindConfig() error {
	err := os.WriteFile("tailwind.config.js", files.TailwindConfig, 0644)
	if err != nil {
		return err
	}

	return nil
}

func CreatePostCSSConfig() error {
	err := os.WriteFile("postcss.config.js", files.PostCSSConfig, 0644)
	if err != nil {
		return err
	}

	return nil
}

func CreateAppCSS() error {
	err := os.WriteFile("src/app.css", files.AppCSS, 0644)
	if err != nil {
		return err
	}

	return nil
}

func CreateDrizzleFiles() error {
	err := os.WriteFile("drizzle.config.ts", files.DrizzleConfig, 0644)
	if err != nil {
		return err
	}

	err = os.WriteFile("scripts/migrate.ts", files.Migrate, 0644)
	if err != nil {
		return err
	}

	err = os.WriteFile("src/lib/server/schema.ts", files.Schema, 0644)
	if err != nil {
		return err
	}

	err = os.WriteFile("scripts/cleanDB.ts", files.CleanDB, 0644)
	if err != nil {
		return err
	}

	err = os.WriteFile("src/lib/server/db.ts", files.DB, 0644)
	if err != nil {
		return err
	}

	return nil
}
