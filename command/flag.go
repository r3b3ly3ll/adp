package command

import "github.com/urfave/cli/v2"

var (

	/*
		ListEntities
	*/

	// ID ...
	ID = &cli.StringFlag{
		Name:  "ID",
		Usage: "Entity ID e.g., singleMindServer.G00000",
	}

	// RelatedEntity ...
	RelatedEntity = &cli.StringFlag{
		Name:  "RelatedEntity",
		Usage: "Related Entity ID e.g., documentHold.G00000",
	}

	// WhiteList ...
	WhiteList = &cli.StringFlag{
		Name:  "WhiteList",
		Usage: "Whitelisted Field Names",
		Value: "id,displayName,processStatus,hostName",
	}

	// Type ...
	Type = &cli.StringFlag{
		Name:  "Type",
		Usage: "Entity Type e.g., singleMindServer, mergingMeta",
	}

	/*
		TaxonomoyStatistic
	*/

	// ApplicationIdentifier ...
	ApplicationIdentifier = &cli.StringFlag{
		Name:  "ApplicationIdentifier",
		Usage: "Application Identifier e.g., documentHold.G00000",
	}

	// TargetTaxonomy ...
	TargetTaxonomy = &cli.StringFlag{
		Name:  "TargetTaxonomy",
		Usage: "Target Taxonomy e.g., rm_loadbatch",
	}

	// ListCategoryProperties ...
	ListCategoryProperties = &cli.StringFlag{
		Name:  "ListCategoryProperties",
		Usage: "List Category Properties",
		Value: "false",
	}

	// ComputeCounts ...
	ComputeCounts = &cli.StringFlag{
		Name:  "ComputeCounts",
		Usage: "Compute Counts",
		Value: "false",
	}

	// EngineName ...
	EngineName = &cli.StringFlag{
		Name:  "EngineName",
		Usage: "Engine Identifier e.g., singleMindServer.G00000",
	}

	// EngineTaxonomies ...
	EngineTaxonomies = &cli.StringFlag{
		Name:  "EngineTaxonomies",
		Usage: "Engine Taxonomies e.g., rm_loadbatch=Google;csv_guts_datatype!=docs",
	}

	/*
		PingProject
	*/

	// Identifiers ...
	Identifiers = &cli.StringFlag{
		Name:  "Identifiers",
		Usage: "List of Identifier e.g., documentHold.G00000,singleMindServer.G00001",
	}

	/*
		StopProcesses
	*/

	// ProcessIdentifiers ...
	ProcessIdentifiers = &cli.StringFlag{
		Name:  "ProcessIdentifiers",
		Usage: "List of Identifier e.g., documentHold.G00000,singleMindServer.G00001",
	}

	/*
	 Shared
	*/

	// EngineUserName ...
	EngineUserName = &cli.StringFlag{
		Name:  "EngineUserName",
		Usage: "Engine User Name, e.g., docs-pull-scripts",
	}

	// EngineUserPassword ...
	EngineUserPassword = &cli.StringFlag{
		Name:  "EngineUserPassword",
		Usage: "Engine User Password",
	}
)
