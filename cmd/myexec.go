package cmd

import (
	"fmt"
	"github.com/eugenmayer/go-exec/exec"
	"github.com/spf13/cobra"
	"log"
)

var myExecCmd = &cobra.Command{
	Use:   "myexec",
	Short: "connect to an ssh host and run hello world",
	Run:   myExecCommand,
}

func init() {
	myExecCmd.Flags().StringVar(&cmd, "command", "echo hello", "Optional, the command to exec locally defaults to 'echo hello'")
	RootCmd.AddCommand(myExecCmd)
}

func myExecCommand(_ *cobra.Command, _ []string) {
	stdout,stderr, err := exec.Run(cmd, verbose)
	if err != nil {
		log.Print(stdout)
		log.Print(stderr)
		log.Fatal(err)
	}
	// else
	log.Print(fmt.Sprintf("your cmmand returned:\n %s",stdout))
}
