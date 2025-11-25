package packagemodel

import (
	"errors"

	"npmupdate/pkg/models/version"
)

type Package struct {
	Name    string
	Wanted  version.Version
	Latest  version.Version
	Current version.Version
}

func New(
	name string,
	wantedVersion string,
	latestVersion string,
	currentVersion string,
) (*Package, error) {
	wanted, err := version.New(wantedVersion)
	if err != nil {
		return nil, errors.New("invalid version")
	}
	latest, err := version.New(latestVersion)
	if err != nil {
		return nil, errors.New("invalid version")
	}
	current, err := version.New(currentVersion)
	if err != nil {
		return nil, errors.New("invalid version")
	}

	return &Package{
		Name:    name,
		Wanted:  *wanted,
		Latest:  *latest,
		Current: *current,
	}, nil
}
