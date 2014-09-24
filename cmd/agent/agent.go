package main

import (
	"flag"
	"fmt"
	"github.com/fsouza/go-dockerclient"
)

var endpoint = flag.String("socket", "unix:///var/run/docker.sock", "the docker socket to use")

type eventListener chan *docker.APIEvents

type containerEvent string

const (
	evtStart   containerEvent = "start"
	evtStop                   = "stop"
	evtDie                    = "die"
	evtDestroy                = "destroy"
)

func main() {
	flag.Parse()
	client, _ := docker.NewClient(*endpoint)
	//client.AddEventListener(listener)
	imgs, _ := client.ListImages(true)
	for _, img := range imgs {
		fmt.Println("ID: ", img.ID)
		fmt.Println("RepoTags: ", img.RepoTags)
		fmt.Println("Created: ", img.Created)
		fmt.Println("Size: ", img.Size)
		fmt.Println("VirtualSize: ", img.VirtualSize)
		fmt.Println("ParentId: ", img.ParentId)
	}

	listener := make(eventListener)
	client.AddEventListener(listener)
	for e := range listener {
		evt := containerEvent(e.Status)
		if evt == evtStart {
			cnt, err := client.InspectContainer(e.ID)
			if err != nil {
				panic(err)
			}
			fmt.Printf("Started: %#v\n", *cnt)
		}
	}
}
