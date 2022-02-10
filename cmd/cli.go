package main

import (
	"opentext.com/axcelerate/adp/client"
	"opentext.com/axcelerate/adp/task"
)

func main() {
	req := task.NewListEntitiesTaskRequest(
		task.WithListEntitiesID("documentHold.Test"),
	)
	client := client.NewClient()

	resp, _ := client.Do(req)
	resp.Output()

}
