package internal

func ConfigureSvelteProject() error {
	err := CreateTailwindConfig()
	if err != nil {
		return err
	}
	err = CreatePostCSSConfig()
	if err != nil {
		return err
	}
	err = CreateAppCSS()
	if err != nil {
		return err
	}
	err = EditPackageJson()
	if err != nil {
		return err
	}
	return nil
}
