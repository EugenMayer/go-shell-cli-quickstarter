package cmd

import (
	"fmt"
	"github.com/eugenmayer/go-sshclient/sshwrapper"
	"github.com/spf13/cobra"
	"log"
)

var (
	localfile 	string
)

var myScpCmd = &cobra.Command{
	Use:   "myscp",
	Short: "connect to an ssh host and run hello world",
	Run:   myScpCommand,
}

func init() {
	myScpCmd.Flags().StringVar(&host, "host", "", "The first option 1")
	myScpCmd.Flags().StringVar(&user, "user", "root", "Optional, the ssh user, defaults to root")
	myScpCmd.Flags().IntVar(&port, "port", 22, "Optional, your ssh port, defaults to 22")
	myScpCmd.Flags().StringVar(&privkey, "key", "~/.ssh/id_rsa", "Optional, your private key, defaults to ~/.ssh/id_rsa")
	myScpCmd.Flags().StringVar(&localfile, "file", "test/dummytestfile", "Optional, the file to transfer to the server, defaults to test/dummytestfile")
	RootCmd.AddCommand(myScpCmd)
}

func myScpCommand(_ *cobra.Command, _ []string) {
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

	// copy the file to the remote host
	fmt.Printf("Copying file from your host to the remote server %s to remote:%s",  localfile, "/tmp/remotetestfile",)
	err = sshApi.CopyToRemote(localfile, "/tmp/remotetestfile")
	if err != nil {
		log.Fatal(err)
	}

	// and back
	fmt.Printf("Copying file from the remote server back to your host from remote:%s to %s", "/tmp/remotetestfile", "/tmp/fileisback")
	err = sshApi.CopyFromRemote("/tmp/remotetestfile", "/tmp/fileisback")
	if err != nil {
		log.Fatal(err)
	}
}
