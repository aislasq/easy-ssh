package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func HandleSftp(args []string) {
	sftpCommand := flag.NewFlagSet("sftp", flag.ExitOnError)
	port, user, key, dir := setSftpVars(sftpCommand)
	sftpCommand.Usage = func() {
		fmt.Println("Usage of sftp:")
		fmt.Println("  essh sftp <host> [<args>]")
		sftpCommand.PrintDefaults()
	}

	switch len(args) {
	case 0:
		sftpCommand.Usage()
	case 1:

		switch args[0] {
		case "-h", "--help":
			sftpCommand.Usage()
		default:
			host, hostIndex := GetHostByName(args[0])
			if hostIndex != -1 {
				keyByName, _ := GetKeyByName(host.Credentials.Key)
				sftpToHost(host.Name, host.Credentials.User, host.Connection.Hostname, host.Connection.Port, keyByName.Path, "")
			} else {
				fmt.Printf("Could not find host by name: %s\n", args[0])
			}
		}

	default:
		host, hostIndex := GetHostByName(args[0])
		if hostIndex != -1 {
			_ = sftpCommand.Parse(args[1:])

			if sftpCommand.Parsed() {
				sftpCommandExec(host, *port, *user, *key, *dir)
			}
		} else {
			fmt.Printf("Could not find host by name: %s\n", args[0])
		}
	}
}

func setSftpVars(sftpCommand *flag.FlagSet) (*string, *string, *string, *string) {
	port := sftpCommand.String("p", "", "Override host default port.")
	user := sftpCommand.String("u", "", "Override host default user.")
	key := sftpCommand.String("k", "", "Override host default key.")
	dir := sftpCommand.String("d", "", "Enter to a specific directory.")
	_ = sftpCommand.Bool("h", false, "Show help.")
	return port, user, key, dir
}

func sftpCommandExec(host Host, port string, user string, key string, dir string) {
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

	sftpToHost(host.Name, user, host.Connection.Hostname, port, key, dir)
}

func sftpToHost(name string, user string, host string, port string, key string, dir string) {
	host = fmt.Sprintf("%s@%s", user, host)

	if dir != "" {
		host = fmt.Sprintf("%s:%s", host, dir)
	}

	fmt.Printf("Connecting SFTP to %s via %s\n\n", name, host)
	var args []string

	if key != "" {
		args = append(args, "-i", key)
	}

	if port != "" {
		args = append(args, "-P", port)
	}

	args = append(args, host)

	cmd := exec.Command("sftp", args...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("There was an error: %s\n", err)
	}
}
