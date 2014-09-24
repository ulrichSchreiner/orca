package main

import (
	"flag"
	"fmt"
	"github.com/fsouza/go-dockerclient"
)

var endpoint = flag.String("socket", "unix:///var/run/docker.sock", "the docker socket to use")

func main() {
	flag.Parse()
	client, _ := docker.NewClient(*endpoint)
	imgs, _ := client.ListImages(true)
	for _, img := range imgs {
		fmt.Println("ID: ", img.ID)
		fmt.Println("RepoTags: ", img.RepoTags)
		fmt.Println("Created: ", img.Created)
		fmt.Println("Size: ", img.Size)
		fmt.Println("VirtualSize: ", img.VirtualSize)
		fmt.Println("ParentId: ", img.ParentId)
	}
}
