package main

import (
//     "fmt"
//     "strings"
    "github.com/spf13/cobra"
    "command"
)

func main() {

  var cmdEcho = &cobra.Command{
    Use:   "echo [string to echo]",
    Short: "Echo anything to the screen",
    Long: `echo is for echoing anything back.
Echo works a lot like print, except it has a child command.`,
    Args: cobra.MinimumNArgs(1),
    Run: command.Echo,
  }

  var rootCmd = &cobra.Command{Use: "app"}
  rootCmd.AddCommand(cmdEcho)
  rootCmd.Execute()
}