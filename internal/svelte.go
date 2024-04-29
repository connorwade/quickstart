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
	err = EditPackageJson(PackageEditConfig{
		DevDependencies: []string{
			"daisyui",
			"autoprefixer",
			"tailwindcss",
			"postcss",
		}})
	if err != nil {
		return err
	}
	return nil
}
