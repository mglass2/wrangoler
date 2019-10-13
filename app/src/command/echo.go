package command

import (
    "fmt"
    "strings"
    "github.com/spf13/cobra"
)
var source string

var Echo = &cobra.Command{
    Use:   "echo [string to echo]",
    Short: "Echo anything to the screen",
    Args: cobra.MinimumNArgs(1),
    Run: action,
}

func init() {
    Echo.Flags().StringVarP(&source, "source", "s", "", "Source directory to read from")
}

func action(cmd *cobra.Command, args []string) {
    fmt.Println("Print: " + strings.Join(args, " "))
}

