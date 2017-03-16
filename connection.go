package main

import (
	"fmt"
	"github.com/g8os/go-client"
)

func GetConnection(id string) client.Client {
	return client.NewClient(fmt.Sprintf("%s:6379", id), "")
}
