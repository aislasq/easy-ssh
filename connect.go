package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func HandleConnect(args []string) {
	connectCommand := flag.NewFlagSet("connect", flag.ExitOnError)
	port, user, key := setConnectVars(connectCommand)
	connectCommand.Usage = func() {
		fmt.Println("Usage of connect:")
		fmt.Println("essh connect <host> [<args>]")
		connectCommand.PrintDefaults()
	}

	switch len(args) {
	case 0:
		connectCommand.Usage()
	case 1:

		switch args[0] {
		case "-h", "--help":
			connectCommand.Usage()
		default:
			host, hostIndex := GetHostByName(args[0])
			if hostIndex != -1 {
				keyByName, _ := GetKeyByName(host.Credentials.Key)
				connectToHost(host.Name, host.Credentials.User, host.Connection.Hostname, host.Connection.Port, keyByName.Path)
			} else {
				fmt.Printf("Could not find host by name: %s\n", args[0])
			}
		}

	default:
		host, hostIndex := GetHostByName(args[0])
		if hostIndex != -1 {
			_ = connectCommand.Parse(args[1:])
			if connectCommand.Parsed() {
				connectCommandExec(host, *port, *user, *key)
			}
		} else {
			fmt.Printf("Could not find host by name: %s\n", args[0])
		}
	}
}

func setConnectVars(connectCommand *flag.FlagSet) (*string, *string, *string) {
	port := connectCommand.String("p", "", "Override host default port.")
	user := connectCommand.String("u", "", "Override host default user.")
	key := connectCommand.String("k", "", "Override host default key.")
	_ = connectCommand.Bool("h", false, "Show help.")
	return port, user, key
}

func connectCommandExec(host Host, port string, user string, key string) {
	if user == "" {
		user = host.Credentials.User
	}

	if port == "" {
		port = host.Connection.Port
	}

	if key == "" {
		keyByName, keyIndex := GetKeyByName(host.Credentials.Key)
		if keyIndex != -1 {
			key = keyByName.Path
		}
	}

	connectToHost(host.Name, user, host.Connection.Hostname, port, key)
}

func connectToHost(name string, user string, host string, port string, key string) {
	fmt.Printf("Connecting to %s via %s@%s\n\n", name, user, host)
	args := []string{host, "-l", user}

	if key != "" {
		args = append(args, "-i", key)
	}

	if port != "" {
		args = append(args, "-p", port)
	}

	cmd := exec.Command("ssh", args...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("There was an error: %s\n", err)
	}
}
