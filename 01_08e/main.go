package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
)

const path = "friends.json"

// Friend represents a friend and their connections.
type Friend struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Friends []string `json:"friends"`
}

// hearGossip indicates that the friend has heard the gossip.
func (f *Friend) hearGossip() {
	log.Printf("%s has heard the gossip!\n", f.Name)
}

// Friends represents the map of friends and connections
type Friends struct {
	fmap map[string]Friend
}

// getFriend fetches the friend given an id.
func (f *Friends) getFriend(id string) Friend {
	return f.fmap[id]
}

// getRandomFriend returns an random friend.
func (f *Friends) getRandomFriend() Friend {
	rand.Seed(time.Now().Unix())
	id := (rand.Intn(len(f.fmap)-1) + 1) * 100
	return f.getFriend(fmt.Sprint(id))
}

// spreadGossip ensures that all the friends in the map have heard the story
func spreadGossip(root Friend, friends Friends, visited map[string]interface{}) {
	for _, id := range root.Friends {
		if _, isVisited := visited[id];!isVisited {
			f := friends.getFriend(id)
			f.hearGossip()
			visited[id] = nil
			spreadGossip(f, friends, visited)
		}
	}
}

func main() {
	friends := importData()
	root := friends.getRandomFriend()
	root.hearGossip()
	visited := make(map[string]interface{})
	visited[root.ID] = nil
	spreadGossip(root, friends, visited)
}

// importData reads the input data from file and creates the friends map.
func importData() Friends {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var data []Friend
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}

	fm := make(map[string]Friend, len(data))
	for _, d := range data {
		fm[d.ID] = d
	}

	return Friends{
		fmap: fm,
	}
}