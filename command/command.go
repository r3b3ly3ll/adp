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

	// TaxonomyStatisticCmd ...
	TaxonomyStatisticCmd = &cli.Command{
		Name:    "taxonomyStatistic",
		Usage:   `adp-cli -p * taxonomyStatistic --EngineTaxonomies "csv_guts_datatype=docs" --TargetTaxonomy rm_loadbatch --ApplicationIdentifier documentHold.G00000`,
		Aliases: []string{"t"},
		Flags: []cli.Flag{
			ApplicationIdentifier,
			EngineTaxonomies,
			EngineName,
			TargetTaxonomy,
			ListCategoryProperties,
			ComputeCounts,
			EngineUserName,
			EngineUserPassword,
		},
		Action: taxonomyStatistic,
	}

	// PingProjectCmd ...
	PingProjectCmd = &cli.Command{
		Name:    "pingProject",
		Usage:   `adp-cli -p * pingProject --Identifiers documentHold.G00000`,
		Aliases: []string{"p"},
		Flags: []cli.Flag{
			Identifiers,
		},
		Action: pingProject,
	}

	Commands = []*cli.Command{
		ListEntitiesCmd,
		TaxonomyStatisticCmd,
		PingProjectCmd,
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

func taxonomyStatistic(c *cli.Context) error {
	var err error

	client.ADP.TaskReq = task.NewTaxonomyStatisticTaskRequest(
		task.WithTaxonomyStatisticEngineTaxonomies(c.String("EngineTaxonomies")),
		task.WithTaxonomyStatisticEngineName(c.String("EngineName")),
		task.WithTaxonomyStatisticOutputTaxonomies(c.String("TargetTaxonomy")),
		task.WithTaxonomyStatisticComputeCounts(c.String("ComputeCounts")),
		task.WithTaxonomyStatisticListCategoryProperties(c.String("ListCategoryProperties")),
		task.WithTaxonomyStatisticApplicationIdentifier(c.String("ApplicationIdentifier")),
		task.WithTaxonomyStatisticEngineUserName(c.String("EngineUserName")),
		task.WithTaxonomyStatisticEngineUserPassword(c.String("EngineUserPassword")),
	)

	if err = ExecuteTask(c); err != nil {
		return fmt.Errorf("task TaxonomyStatistic: %w", err)
	}

	return nil
}

func pingProject(c *cli.Context) error {
	var err error

	client.ADP.TaskReq = task.NewPingProjectTaskRequest(
		task.WithPingProjectIdentifiers(c.String("Identifiers")),
	)

	if err = ExecuteTask(c); err != nil {
		return fmt.Errorf("task TaxonomyStatistic: %w", err)
	}
	return nil
}
