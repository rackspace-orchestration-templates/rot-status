package main

import (
	"io"
	"fmt"
	"net/http"
	"github.com/google/go-github/github"
)

func statusPage(w http.ResponseWriter, req *http.Request) {
	// Connect to Github and list public repos for rot
	client := github.NewClient(nil)
	opt := &github.RepositoryListByOrgOptions{
		Type: "public",
		ListOptions: github.ListOptions{PerPage: 50},
	}
	repos, _, err := client.Repositories.ListByOrg(
		"rackspace-orchestration-templates",
		opt,
	)

	if err != nil {
		// Write out status unknown
		io.WriteString(w, err.Error())
	} else {
		for _, repo := range repos {
			io.WriteString(w, *repo.Name)
		}
        }
}

func main() {
	// Serve up the index page
	http.HandleFunc("/", statusPage)
	fmt.Println("Starting HTTP Server ...")
	http.ListenAndServe(":8080", nil)
}
