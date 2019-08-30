package main

import (
	"fmt"
	"math"
	"strings"
)

func HandleView() {
	hosts()
	keys()
}

func hosts() {
	var maxName, maxConn, maxKey float64
	var hosts []map[string]string
	for _, h := range config.Hosts {
		host := map[string]string{}
		host["name"] = h.Name

		host["conn"] = fmt.Sprintf("%s@%s", h.Credentials.User, h.Connection.Hostname)
		if h.Connection.Port != "" {
			host["conn"] += ":" + h.Connection.Port
		}

		if h.Credentials.Key != "" {
			host["key"] = h.Credentials.Key
		}

		hosts = append(hosts, host)

		maxName = math.Max(float64(len(host["name"])), maxName)
		maxConn = math.Max(float64(len(host["conn"])), maxConn)
		maxKey = math.Max(float64(len(host["key"])), maxKey)
	}
	length := int(maxName + maxConn + maxKey)

	lineFormat := fmt.Sprintf("%%-%ds%%-%ds%%-%ds\n", int(maxName)+2, int(maxConn)+2, int(maxKey)+1)
	sep := strings.Join([]string{strings.Repeat("-", int(maxName)+1), strings.Repeat("-", int(maxConn)+1), strings.Repeat("-", int(maxKey)+1)}, "+")
	head := strings.Repeat("*", length/2+length%2) + "HOSTS" + strings.Repeat("*", length/2)

	fmt.Println("\n" + head)
	fmt.Printf(lineFormat, "Name", "User@Host", "Key")
	fmt.Println(sep)

	for _, h := range hosts {
		fmt.Printf(lineFormat, h["name"], h["conn"], h["key"])
	}

	fmt.Print("\n")
}

func keys() {
	var maxName, maxPath float64
	var keys []map[string]string
	for _, k := range config.Keys {
		key := map[string]string{}

		key["name"] = k.Name
		key["path"] = k.Path

		keys = append(keys, key)

		maxName = math.Max(float64(len(key["name"])), maxName)
		maxPath = math.Max(float64(len(key["path"])), maxPath)

	}
	length := int(maxName+maxPath) - 1
	lineFormat := fmt.Sprintf("%%-%ds%%-%ds\n", int(maxName)+2, int(maxPath)+1)
	sep := strings.Join([]string{strings.Repeat("-", int(maxName)+1), strings.Repeat("-", int(maxPath)+1)}, "+")
	head := strings.Repeat("*", length/2+length%2) + "KEYS" + strings.Repeat("*", length/2)

	fmt.Println("\n" + head)
	fmt.Printf(lineFormat, "Name", "Path")
	fmt.Println(sep)

	for _, k := range keys {
		fmt.Printf(lineFormat, k["name"], k["path"])
	}

	fmt.Print("\n")
}
