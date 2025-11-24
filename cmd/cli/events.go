package main

import (
	"errors"

	"npmupdate/pkg/entities"

	tea "github.com/charmbracelet/bubbletea"
)

type OutdatedPackagesMsg []entities.Package

func getOutdatedPackages() tea.Msg {
	result, err := entities.FetchOutdatedPackages()
	if err != nil {
		panic(err)
	}

	packages := []entities.Package{}
	for packageName, value := range result {
		pkg, err := entities.NewPackage(
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
