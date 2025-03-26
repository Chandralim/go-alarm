package main

import (
	"fmt"
	"io"
	"mygo/konstanta"
	"os"
	"os/exec"
	"runtime"

	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/spf13/cobra"
)

// _ "github.com/go-sql-driver/mysql"

var act string
var tableName string

var dbURL string
var tablePath = "./migration/tables"

var rootCmd = &cobra.Command{Use: "mycli"}
var cmd1 = ""
var cmd2 = ""

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "create migrations",
	Run: func(cmd *cobra.Command, args []string) {

		cmdStr := fmt.Sprintf(`migrate create -ext sql -dir %s -seq %s`, tablePath, tableName)
		fmt.Println("cmdStr:", cmdStr)

		var cmdTerminal *exec.Cmd

		cmdTerminal = exec.Command(cmd1, cmd2, cmdStr)

		output, err := cmdTerminal.CombinedOutput()

		if err != nil {
			fmt.Println("Error:", err)
			fmt.Println("Details:", string(output))
		} else {
			fmt.Println("Output:", string(output))
		}

	},
}

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Apply database migrations",
	Run: func(cmd *cobra.Command, args []string) {

		// cmdStr := fmt.Sprintf(`migrate -database "%s" -path %s %s`, dbURL, tablePath, act)
		// // cmdStr := fmt.Sprintf(`migrate -help`)
		// fmt.Println("cmdStr:", cmdStr)

		// var cmdTerminal *exec.Cmd

		// cmdTerminal = exec.Command(cmd1, cmd2, cmdStr)

		// output, err := cmdTerminal.CombinedOutput()

		cmdPath := "C:\\Users\\chand\\go\\bin\\migrate.exe" // Use the full path
		cmdArgs := []string{"-verbose", "-database", dbURL, "-path", tablePath, act}

		// Debugging: Print the command being executed
		fmt.Println("Executing command:", cmdPath, cmdArgs)

		// Create the command with explicit environment variables
		cmdTerminal := exec.Command(cmdPath, cmdArgs...)
		cmdTerminal.Env = os.Environ() // Ensure it inherits all environment variables

		// Get the stdin pipe
		stdin, err := cmdTerminal.StdinPipe()
		if err != nil {
			fmt.Println("Error getting stdin pipe:", err)
			return
		}

		// Start the command
		// if err := cmdTerminal.Start(); err != nil {
		// 	fmt.Println("Error starting command:", err)
		// 	return
		// }

		// Write "y\n" to stdin to simulate user confirmation
		_, err = io.WriteString(stdin, "y\n")
		if err != nil {
			fmt.Println("Error writing to stdin:", err)
			return
		}
		stdin.Close() // Close stdin to signal input is complete
		output, err := cmdTerminal.CombinedOutput()
		if err != nil {
			fmt.Println("Error:", err)
			fmt.Println("Details:", string(output))

		} else {
			fmt.Println("Output:", string(output))
		}
	},
}

func init() {
	doCmd.Flags().StringVarP(&act, "act", "a", "version", "Check Version")

	migrateCmd.Flags().StringVarP(&tableName, "table_name", "t", "x", "Create Table")

	dbURL = "mysql://" + *konstanta.Connection

	if runtime.GOOS == "windows" {
		cmd1 = "cmd"
		cmd2 = "/C"
	} else {
		cmd1 = "sh"
		cmd2 = "-c"
	}

	// migrationPath = fmt.Sprintf("")

	rootCmd.AddCommand(doCmd)
	rootCmd.AddCommand(migrateCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
