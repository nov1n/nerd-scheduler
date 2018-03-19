package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var listOnly bool

type NodeList struct {
	Items []Node `json:"items"`
}

type Node struct {
	Metadata Metadata `json:"metadata"`
}

type Metadata struct {
	Name        string            `json:"name,omitempty"`
	Annotations map[string]string `json:"annotations"`
}

func main() {
	flag.BoolVar(&listOnly, "l", false, "List current annotations and exist")
	flag.Parse()

	temps := []string{"30", "54", "77", "11", "41", "13"}
	resp, err := http.Get("http://127.0.0.1:8001/api/v1/nodes")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if resp.StatusCode != 200 {
		fmt.Println("Invalid status code", resp.Status)
		os.Exit(1)
	}

	var nodes NodeList
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&nodes)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if listOnly {
		for _, node := range nodes.Items {
			temp := node.Metadata.Annotations["nerdalize/temp"]
			fmt.Printf("%s %s\n", node.Metadata.Name, temp)
		}
		os.Exit(0)
	}

	rand.Seed(time.Now().Unix())
	for _, node := range nodes.Items {
		temp := temps[rand.Intn(len(temps))]
		annotations := map[string]string{
			"nerdalize/temp": temp,
		}
		patch := Node{
			Metadata{
				Annotations: annotations,
			},
		}

		var b []byte
		body := bytes.NewBuffer(b)
		err := json.NewEncoder(body).Encode(patch)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		url := "http://127.0.0.1:8001/api/v1/nodes/" + node.Metadata.Name
		request, err := http.NewRequest("PATCH", url, body)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		request.Header.Set("Content-Type", "application/strategic-merge-patch+json")
		request.Header.Set("Accept", "application/json, */*")

		resp, err := http.DefaultClient.Do(request)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if resp.StatusCode != 200 {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("%s %s\n", node.Metadata.Name, temp)
	}
}
