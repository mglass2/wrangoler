package command

import (
    "fmt"
    "github.com/spf13/cobra"
)

type generateDatabaseSchemaCommand struct {
    use string
    short string
    long string
}

func (c generateDatabaseSchemaCommand) action(cmd *cobra.Command, args []string) {
    fmt.Println("generateDatabaseSchema")
}

func ProvideGenerateDatabaseSchemaCommand() generateDatabaseSchemaCommand {
    var command = generateDatabaseSchemaCommand{"db [db to connect]", "short desc", "long desc"}
    return command
}

func (c generateDatabaseSchemaCommand) MakeCobraCommand() *cobra.Command{
    var cobraConf = &cobra.Command{
        Use:   c.use,
        Short: c.short,
        Long: c.long,
        Args: cobra.MinimumNArgs(1),
        Run: c.action,
    }

    return cobraConf
}