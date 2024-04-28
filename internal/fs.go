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
