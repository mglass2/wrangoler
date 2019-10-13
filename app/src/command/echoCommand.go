package command

import (
    "fmt"
    "strings"
    "github.com/spf13/cobra"
)

func Echo(cmd *cobra.Command, args []string) {
    fmt.Println("Print: " + strings.Join(args, " "))
}
