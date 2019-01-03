package main

import (
    "fmt"
    "github.com/xanzy/go-gitlab"
)

func main() {
    //client, err := gitlab.NewBasicAuthClient(nil, "https://0.0.0.0:80/", "root", "adminadmin")
    client := gitlab.NewClient(nil, "HqHfTWxNqKH7s3xoGf3M")
    //if err != nil {
    //    fmt.Printf("create client error: %v \n", err)
    //}
    // List all projects
    if err := client.SetBaseURL("http://0.0.0.0:80"); err != nil {
        fmt.Printf("create client error: %v \n", err)
    }
    fmt.Println("hello")
    projects, _, err := client.Projects.ListProjects(nil)
    if err != nil {
        fmt.Printf("list project error: %v \n", err)
    }

    fmt.Printf("Found %d projects \n", len(projects))
}
