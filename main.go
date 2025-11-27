package main

import (
	"github.com/3sarojbhattarai/govault/cmd"
	"github.com/3sarojbhattarai/govault/internal/commands"
)

func main() {
	// Register all commands
	cmd.RegisterCommands(
		commands.AddCmd,
		commands.GetCmd,
		commands.ListCmd,
		commands.DeleteCmd,
		commands.GenerateCmd,
	)

	// Execute CLI
	cmd.Execute()
}
