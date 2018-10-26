package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var (
	opt1 	string
	flag1   bool
)

var mySubCmd = &cobra.Command{
	Use:   "mysub",
	Short: "First sub command to do mysub",
	Run:   mySubCommandRun,
}

func init() {
	mySubCmd.Flags().StringVar(&opt1, "opt1", "", "The first option 1")
	mySubCmd.Flags().BoolVar(&flag1, "flag1", false, "One option to turn of an on, defaults to false")


	RootCmd.AddCommand(mySubCmd)
}

func mySubCommandRun(_ *cobra.Command, _ []string) {
	if opt1 == "" {
		log.Fatal("please set the --opt1 option")
	}

	log.Printf("You have set opt1 to: %s", opt1)

	if flag1 {
		log.Print("You enabled flag1")
	} else {
		log.Print("You disabled flag1")
	}

	if verbose {
		log.Print("You enabled the global verbose flag")
	} else {
		log.Print("You disabled the global verbose flag")
	}
}
