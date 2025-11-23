package entities

import (
	"encoding/json"
	"strings"
)

type Package struct {
	Name     string
	Current  string `json:"current"`
	Wanted   string `json:"wanted"`
	Latest   string `json:"latest"`
	selected string
}

func ParseJSON() []Package {
	// output, err := exec.Command("npm", "outdated", "--json").Output()
	// if err != nil {
	// 	fmt.Println("🪚 err:", err)
	// }

	var outdated map[string]Package
	err := json.NewDecoder(strings.NewReader(string(output))).Decode(&outdated)
	if err != nil {
		panic(err)
	}

	packages := make([]Package, 0, len(outdated))
	for packageName, value := range outdated {
		packages = append(packages, Package{
			Name:    packageName,
			Current: value.Current,
			Wanted:  value.Wanted,
			Latest:  value.Latest,
		})
	}

	return packages
}

var output string = `
	{
	  "react-native-reanimated": {
    "current": "4.1.0",
    "wanted": "4.1.3",
    "latest": "4.1.3"
  },
  "react-native-screens": {
    "current": "4.16.0",
    "wanted": "4.17.1",
    "latest": "4.17.1"
  },
  "react-native-svg": {
    "current": "15.13.0",
    "wanted": "15.14.0",
    "latest": "15.14.0"
  },
  "react-native-worklets": {
    "current": "0.5.1",
    "wanted": "0.5.1",
    "latest": "0.6.1"
  }
}`
