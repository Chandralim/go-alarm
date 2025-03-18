package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"mygo/konstanta"

	"github.com/spf13/cobra"
)

var act string
var tableName string

var dbURL *string
var migrationPath = ""

var rootCmd = &cobra.Command{Use: "mycli"}
var cmd1 = ""
var cmd2 = ""

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "create migrations",
	Run: func(cmd *cobra.Command, args []string) {

		cmdStr := fmt.Sprintf(`migrate create -ext sql -dir migrations -seq %s`, tableName)

		var cmdTerminal *exec.Cmd

		cmdTerminal = exec.Command(cmd1, cmd2, cmdStr)

		output, err := cmdTerminal.CombinedOutput()

		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Output:", string(output))
		}

	},
}

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Apply database migrations",
	Run: func(cmd *cobra.Command, args []string) {

		cmdStr := fmt.Sprintf(`migrate -database "%s" -path migrations %s`, dbURL, act)

		var cmdTerminal *exec.Cmd

		cmdTerminal = exec.Command(cmd1, cmd2, cmdStr)

		output, err := cmdTerminal.CombinedOutput()

		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Output:", string(output))
		}
	},
}

func init() {
	doCmd.Flags().StringVarP(&act, "act", "a", "version", "Check Version")

	migrateCmd.Flags().StringVarP(&tableName, "table_name", "tn", "x", "Create Table")

	dbURL = konstanta.Connection

	if runtime.GOOS == "windows" {
		cmd1 = "cmd"
		cmd2 = "/C"
	} else {
		cmd1 = "sh"
		cmd2 = "-c"
	}

	migrationPath = fmt.Sprintf("")

	rootCmd.AddCommand(doCmd)
	rootCmd.AddCommand(migrateCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
