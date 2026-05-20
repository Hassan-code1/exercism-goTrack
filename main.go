package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type Profile struct {
	User struct {
		Handle string `json:"handle"`
	} `json:"user"`

	Reputation int `json:"reputation"`
}

func main() {

	username := "Hassan-code1"

	url := fmt.Sprintf(
		"https://exercism.org/api/v2/profiles/%s",
		username,
	)

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var profile Profile

	err = json.Unmarshal(body, &profile)
	if err != nil {
		panic(err)
	}

	stats := fmt.Sprintf(`
- 👤 Username: %s
- 🏆 Reputation: %d
- 📅 Last Updated: Auto Updated
`,
		profile.User.Handle,
		profile.Reputation,
	)

	readme, err := os.ReadFile("README.md")
	if err != nil {
		panic(err)
	}

	content := string(readme)

	start := "<!-- EXERCISM_STATS_START -->"
	end := "<!-- EXERCISM_STATS_END -->"

	newSection := start + "\n" + stats + "\n" + end

	before := strings.Split(content, start)[0]
	after := strings.Split(content, end)[1]

	updated := before + newSection + after

	err = os.WriteFile(
		"README.md",
		[]byte(updated),
		0644,
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("README updated successfully")
}
