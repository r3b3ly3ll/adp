package main

import (
	"fmt"

	"opentext.com/axcelerate/adp/client"
	"opentext.com/axcelerate/adp/task"
)

func main() {
	req := task.NewListEntitiesTaskRequest(
		task.WithListEntitiesID("documentHold.Test"),
	)
	client := client.NewClient()

	resp, _ := client.Do(req)
	if resp.ExecutionStatus == "success" {
		fmt.Println("hello")
	}

}
