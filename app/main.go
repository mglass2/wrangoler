package main

import (
//     "fmt"
//     "strings"
    "github.com/spf13/cobra"
    "command"
)

func main() {

  var rootCmd = &cobra.Command{Use: "app"}
  rootCmd.AddCommand(command.Echo)
  rootCmd.AddCommand(command.ProvideGenerateDatabaseSchemaCommand().MakeCobraCommand())
  rootCmd.Execute()
}