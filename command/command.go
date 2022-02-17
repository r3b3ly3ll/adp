package command

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/urfave/cli/v2"
	"opentext.com/axcelerate/adp/client"
	"opentext.com/axcelerate/adp/task"
)

var (
	// ListEntitiesCmd ...
	ListEntitiesCmd = &cli.Command{
		Name:    "listEntities",
		Usage:   `adp-cli -p * listEntities --Type dataSource`,
		Aliases: []string{"l"},
		Flags: []cli.Flag{
			ID,
			RelatedEntity,
			WhiteList,
			Type,
		},
		Action: listEntities,
	}

	Commands = []*cli.Command{
		ListEntitiesCmd,
	}
)

func ExecuteTask(c *cli.Context) error {
	var err error

	if c.Bool("debug") {
		client.ADP.TaskReq.TaskConfiguration.EnableAdpLogging()
		client.ADP.TaskReq.TaskConfiguration.EnableAdpExecutionPersistent()
	}

	if err = client.ADP.Run(); err != nil {
		return fmt.Errorf("executeTask: %w", err)
	}

	if client.ADP.TaskResp.IsSuccess() {
		return fmt.Errorf("%s", "executeTask: status does not match success")
	}

	output := client.ADP.TaskResp.ExecutionMetaData.Output()

	if c.Bool("pretty") {
		buf := new(bytes.Buffer)
		json.Indent(buf, []byte(output), "", "  ")
		fmt.Println(buf)
	} else {
		fmt.Println(output)
	}

	return nil
}

func listEntities(c *cli.Context) error {
	var err error

	client.ADP.TaskReq = task.NewListEntitiesTaskRequest(
		task.WithListEntitiesID(c.String("ID")),
		task.WithListEntitiesRelatedEntity(c.String("RelatedEntity")),
		task.WithListEntitiesType(c.String("Type")),
		task.WithListEntitiesWhiteList(c.String("WhiteList")),
	)

	// initialize ListEntitiesTaskResponse with specific ExecutionMetaData struct.
	client.ADP.TaskResp = task.NewListEntitiesTaskResponse()

	if err = ExecuteTask(c); err != nil {
		return fmt.Errorf("task ListEntities: %w", err)
	}

	return nil
}
