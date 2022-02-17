package command

import "github.com/urfave/cli/v2"

var (
	// ID ...
	ID = &cli.StringFlag{
		Name:  "ID",
		Usage: "Entity ID e.g., singleMindServer.G00000, documentHold.G00000",
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
)
