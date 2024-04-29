package internal

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

type PackageJSONSpec struct {
	Name            string            `json:"name"`
	Version         string            `json:"version"`
	Private         bool              `json:"private"`
	Type            string            `json:"type"`
	Dependencies    map[string]string `json:"dependencies"`
	DevDependencies map[string]string `json:"devDependencies"`
	Scripts         map[string]string `json:"scripts"`
}

type PackageData struct {
	DistTags struct {
		Latest string `json:"latest"`
	} `json:"dist-tags"`
}

type PackageEditConfig struct {
	DevDependencies []string
	Dependencies    []string
	Scripts         [][2]string
}

func getLatestVersion(packageName string) (string, error) {
	resp, err := http.Get("https://registry.npmjs.org/" + packageName)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var data PackageData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", err
	}

	return data.DistTags.Latest, nil
}

func readPackageJson() ([]byte, error) {
	file, err := os.ReadFile("package.json")
	if err != nil {
		return nil, err
	}

	return file, nil
}

func EditPackageJson(c PackageEditConfig) error {
	file, err := readPackageJson()
	if err != nil {
		return err
	}

	var p PackageJSONSpec
	json.Unmarshal(file, &p)

	for _, dep := range c.DevDependencies {
		latestVersion, err := getLatestVersion(dep)
		if err != nil {
			return err
		}
		p.DevDependencies[dep] = "^" + latestVersion
	}

	for _, scriptArray := range c.Scripts {
		p.Scripts[scriptArray[0]] = scriptArray[1]
	}

	for _, dep := range c.Dependencies {
		latestVersion, err := getLatestVersion(dep)
		if err != nil {
			return err
		}
		p.Dependencies[dep] = "^" + latestVersion
	}

	if p.Dependencies == nil {
		p.Dependencies = make(map[string]string)
	}

	if p.DevDependencies == nil {
		p.DevDependencies = make(map[string]string)
	}

	if p.Scripts == nil {
		p.Scripts = make(map[string]string)
	}

	newFile, err := json.MarshalIndent(p, "", "    ")

	if err != nil {
		return err
	}

	err = os.WriteFile("package.json", newFile, 0644)
	if err != nil {
		return err
	}

	return nil
}
