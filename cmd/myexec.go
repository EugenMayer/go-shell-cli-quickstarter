package cmd

import (
	"fmt"
	"github.com/eugenmayer/go-exec/exec"
	"github.com/eugenmayer/go-exec/utils/dialog"
	"github.com/spf13/cobra"
	"log"
)

var (
	shellCmd 	string
)

var myExecCmd = &cobra.Command{
	Use:   "myexec",
	Short: "connect to an ssh host and run hello world",
	Run:   myExecCommand,
}

func init() {
	myExecCmd.Flags().StringVar(&shellCmd, "command", "echo hello", "Optional, the command to exec locally defaults to 'echo hello'")
	RootCmd.AddCommand(myExecCmd)
}

func myExecCommand(_ *cobra.Command, _ []string) {
	stdout,stderr, err := exec.Run(shellCmd, verbose)
	if err != nil {
		log.Print(stdout)
		log.Print(stderr)
		log.Fatal(err)
	}
	
	// else
	log.Print(fmt.Sprintf("your command returned:\n %s",stdout))
	
	// question based execution, get confirmaion
	if utils.ConfirmQuestion("Should i show you the curren time?") {
		stdout,stderr, err := exec.Run("date", verbose)
		if err != nil {
			log.Print(stdout)
			log.Print(stderr)
			log.Fatal(err)
		} 
		// else
		log.Print(fmt.Sprintf("date:\n %s",stdout))
	}
}
