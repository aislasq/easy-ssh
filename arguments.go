package main

import "os"

type Args struct {
	Command string
	Args    []string
	Length  int
}

func GetArgs() Args {
	var OsIn = Args{"", []string{}, 0}
	OsIn.Length = len(os.Args)

	if OsIn.Length > 1 {
		OsIn.Command = os.Args[1]
		OsIn.Args = os.Args[2:]
	}
	return OsIn
}
