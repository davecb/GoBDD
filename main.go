package main

import (
	"github.com/google/go-github/github"
	"context"
	"fmt"
	"GoBDD/services"
)

// Model
type Package struct {
	FullName      string
	Description   string
	StarsCount    int
	ForksCount    int
	LastUpdatedBy string
}

func main() {
	ctx := context.Background()
	client := github.NewClient(nil)

	// Step 1 - create github api client
	githubAPI := services.NewGithub(ctx, client.Repositories)

	// Step 1 - Get Repository Package Information
	pack, err := githubAPI.GetPackageRepoInfo("Golang-Coach", "Lessons")
	fmt.Printf("%+v\n", pack)
	fmt.Printf("%+v\n", err)
}

