package cmd

import (
	"fmt"
	"github.com/eugenmayer/go-sshclient/sshwrapper"
	"github.com/spf13/cobra"
	"log"
)

var (
	host 	string
	user 	string
	port    int
	privkey		string
	cmd 	string
)

var mySshCmd = &cobra.Command{
	Use:   "myssh",
	Short: "connect to an ssh host and run hello world",
	Run:   mySshCommand,
}

func init() {
	mySshCmd.Flags().StringVar(&host, "host", "", "The first option 1")
	mySshCmd.Flags().StringVar(&user, "user", "root", "Optional, the ssh user, defaults to root")
	mySshCmd.Flags().IntVar(&port, "port", 22, "Optional, your ssh port, defaults to 22")
	mySshCmd.Flags().StringVar(&privkey, "key", "~/.ssh/id_rsa", "Optional, your private key, defaults to ~/.ssh/id_rsa")
	mySshCmd.Flags().StringVar(&cmd, "command", "echo hello", "Optional, command to run, defaults to 'echo hello'")
	RootCmd.AddCommand(mySshCmd)
}

func mySshCommand(_ *cobra.Command, _ []string) {
	if host == "" {
		log.Fatal("please set the --host option")
	}

	sshApi := sshwrapper.SshApi{Host:host,Port:port, User:user}
	// this could be also
	// for ssh agent setup: sshApi.DefaultSshAgentSetup()
	// for password based connection: sshApi.DefaultSshPasswordSetup()
	err := sshApi.DefaultSshPrivkeySetup(privkey)
	if err != nil {
		log.Fatal(err)
	}

	stdout, stderr, err := sshApi.Run(cmd)
	if err != nil {
		log.Print(stdout)
		log.Print(stderr)
		log.Fatal(err)
	}
	// else
	log.Print(fmt.Sprintf("your ssh host '%s' returned:\n %s",sshApi.Host, stdout))
}
