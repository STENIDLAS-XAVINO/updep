package main

import (
	"errors"

	packagemodel "updep/pkg/models/package"

	tea "github.com/charmbracelet/bubbletea"
)

type OutdatedPackagesMsg []packagemodel.Package

func getOutdatedPackages() tea.Msg {
	result, err := packagemodel.FetchOutdatedPackages()
	if err != nil {
		panic(err)
	}

	packages := []packagemodel.Package{}
	for packageName, value := range result {
		pkg, err := packagemodel.New(
			packageName,
			value.Wanted,
			value.Latest,
			value.Current,
		)
		if err != nil {
			_ = errors.New("invalid package versions")
			// TODO: handle error
			continue
		}

		packages = append(packages, *pkg)
	}

	return OutdatedPackagesMsg(packages)
}
