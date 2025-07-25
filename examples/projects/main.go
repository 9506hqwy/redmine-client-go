package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	client "github.com/9506hqwy/redmine-client-go/pkg/redmine"
)

func basicAuth(ctx context.Context, req *http.Request) error {
	req.SetBasicAuth("admin", "admin")
	return nil
}

func main() {
	hc := http.Client{}

	c, err := client.NewClientWithResponses("http://127.0.0.1:3000", client.WithHTTPClient(&hc))
	if err != nil {
		log.Fatal(err)
	}

	params := client.ProjectsIndexParams{}
	resp, err := c.ProjectsIndexWithResponse(context.TODO(), &params, basicAuth)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode() != http.StatusOK {
		log.Fatalf("Expected HTTP 200 but received %d", resp.StatusCode())
	}

	for _, project := range *resp.JSON200.Projects {
		fmt.Printf("Project Identifier: %s, Name: %s\n", *project.Identifier, *project.Name)
	}
}
