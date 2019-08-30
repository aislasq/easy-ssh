package main

import (
	"flag"
	"fmt"
	"os/exec"
)

func HandleRun(args []string) {
	runCommand := flag.NewFlagSet("run", flag.ExitOnError)
	run, user, port, key := setRunVars(runCommand)
	runCommand.Usage = func() {
		fmt.Println("Usage of run:")
		fmt.Println("  essh run <host> -r '<command>' [<args>]")
		runCommand.PrintDefaults()
	}

	switch len(args) {
	case 0:
		runCommand.Usage()
	case 1:
		switch args[0] {
		case "-h", "--help":
			runCommand.Usage()
		default:
			fmt.Println("Both host and -r '<args>' are needed")
		}
	default:
		host, hostIndex := GetHostByName(args[0])
		if hostIndex != -1 {
			_ = runCommand.Parse(args[1:])
			if runCommand.Parsed() {
				runCommandExec(host, *run, *user, *port, *key)
			}
		} else {
			fmt.Printf("Could not find host by name: %s\n", args[0])
		}
	}

}

func setRunVars(runCommand *flag.FlagSet) (*string, *string, *string, *string) {
	key := runCommand.String("k", "", "Override host default key.")
	port := runCommand.String("p", "", "Override host default port.")
	user := runCommand.String("u", "", "Override host default user.")
	run := runCommand.String("r", "", "Set Command to Run: ex. -r 'ls /var/www' (Required)")
	_ = runCommand.Bool("h", false, "Show help.")
	return run, user, port, key
}

func runCommandExec(host Host, run string, user string, port string, key string) {
	if run == "" {
		fmt.Println("Please supply the command to run using -r option.")
		return
	}

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

	runInHost(host.Name, user, host.Connection.Hostname, port, key, run)
}

func runInHost(name string, user string, host string, port string, key string, run string) {
	fmt.Printf("Running %s in %s via %s@%s\n", run, name, user, host)
	args := []string{host, "-l", user}

	if key != "" {
		args = append(args, "-i", key)
	}

	if port != "" {
		args = append(args, "-p", port)
	}

	args = append(args, run)
	out, err := exec.Command("ssh", args...).Output()

	if err != nil {
		fmt.Printf("There was an error: %s\n", err)
		return
	}

	fmt.Printf("%s\n", out)
}
