package main

import "fmt"

func main() {
	OsIn := GetArgs()

	if OsIn.Length == 1 {
		printUsage()
		return
	}

	switch OsIn.Command {
	case "connect":
		HandleConnect(OsIn.Args)
	case "run":
		HandleRun(OsIn.Args)
	case "sftp":
		HandleSftp(OsIn.Args)
	case "view":
		HandleView()
	case "help":
		printUsage()
	default:
		fmt.Printf("%q is not valid command.\n", OsIn.Command)
	}
}

func printUsage() {
	fmt.Println("usage: essh <command> <host> [<args>]")
	fmt.Println("The commands are: ")
	fmt.Println("  -c, --connect    Connect to host with ssh")
	fmt.Println("  -f, --sftp       Connect to host with sftp")
	fmt.Println("  -r, --run        Run command on host")
	fmt.Println("  -v, --view       View hosts and keys")
	fmt.Println("  -h, --help       View this page")
}
