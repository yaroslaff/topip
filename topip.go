package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/dlclark/regexp2"
)

func printTop(m map[string]int, top int) {

	type kv struct {
		Key   string
		Value int
	}

	var ss []kv
	for k, v := range m {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value < ss[j].Value
	})

	if top > len(ss) {
		top = len(ss)
	}

	for _, kv := range ss[len(ss)-top:] {
		fmt.Printf("%8d %s\n", kv.Value, kv.Key)
	}

}

func main() {

	var scanner *bufio.Scanner
	// regex := "\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}"
	//regex := `\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}(?!\.)`
	regex := `\b\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}\b(?!\.)`

	r, err := regexp2.Compile(regex, 0)
	if err != nil {
		log.Fatal("Regex compilation failed: ", err)
	}

	ipcount := make(map[string]int)

	grepmode := flag.Bool("g", false, "grep mode (full strings)")
	ipmode := flag.Bool("i", false, "IPv4 mode (only IP addresses)")
	top := flag.Int("t", 10, "print top N IPs")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(),
			"Usage: %s [filename]\n"+
				"  if no filename, then stdin", os.Args[0])
		flag.PrintDefaults()
	}

	/* Parse arguments */
	flag.Parse()
	topmode := !(*grepmode || *ipmode)

	infile := flag.Arg(0)

	if infile != "" {
		file, err := os.Open(infile)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		scanner = bufio.NewScanner(file)
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}

	for scanner.Scan() {
		if matched, _ := r.MatchString(scanner.Text()); matched {
			line := scanner.Text()
			ipaddr_match, _ := r.FindStringMatch(scanner.Text())
			ipaddr := ipaddr_match.String()
			if topmode {
				// top mode
				ipcount[ipaddr]++
			} else {
				// grep mode
				if *ipmode {
					fmt.Println(ipaddr)
				} else {
					fmt.Println(line)
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if *top > 0 {
		printTop(ipcount, *top)
	}

}
