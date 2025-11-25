package packagemodel

import (
	"encoding/json"
	"strings"
	"time"
)

type JSONPackage struct {
	Name    string
	Wanted  string `json:"wanted"`
	Latest  string `json:"latest"`
	Current string `json:"current"`
}

func FetchOutdatedPackages() (map[string]JSONPackage, error) {
	timer := time.NewTimer(time.Second * 1)
	<-timer.C
	// output, err := exec.Command("npm", "outdated", "--json").Output()
	// if err != nil {
	// 	fmt.Println("🪚 err:", err)
	// }
	var err error

	var outdated map[string]JSONPackage
	err = json.NewDecoder(strings.NewReader(string(output))).Decode(&outdated)
	if err != nil {
		return nil, err
	}

	return outdated, nil
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
