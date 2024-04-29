package internal

func ConfigureDrizzleProject() error {
	err := CreateDrizzleFiles()
	if err != nil {
		return err
	}

	err = EditPackageJson(PackageEditConfig{
		DevDependencies: []string{
			"drizzle-kit",
			"tsx",
		},
		Dependencies: []string{
			"drizzle-orm",
			"better-sqlite3",
		},
		Scripts: [][2]string{
			{"db:clean", "tsx scripts/cleanDB.ts"},
			{"db:migrate", "tsx scripts/migrate.ts"},
			{"db:generate", "drizzle-kit generate:sqlite"},
			{"db:restart", "db:clean && db:generate"},
		}})

	if err != nil {
		return err
	}

	return nil
}
