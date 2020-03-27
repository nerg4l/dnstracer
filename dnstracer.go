package main

import (
	"flag"
	"fmt"

	"github.com/Redislabs-Solution-Architects/dnstracer/collection"
	"github.com/Redislabs-Solution-Architects/dnstracer/rules"
	"github.com/Redislabs-Solution-Architects/dnstracer/suggestions"
)

func main() {
	cluster := flag.String("cluster-fqdn", "", "The name of the redis cluster eg: redis-10000.foo.example.com")
	dbg := flag.Bool("debug", false, "Show debug information")
	suggest := flag.Bool("suggest", false, "Suggest possible fixes")
	flag.Parse()

	coll := collection.Collect(*cluster)
	results := rules.Check(coll, *dbg)

	//fmt.Println("FML")
	//coll.LocalNS = []string{}
	//coll.LocalMatchNS = false
	//results.ResultNS = false

	if *dbg {
		fmt.Printf("--------------------------------\n")
		fmt.Printf("%+v\n", coll)
		fmt.Printf("%+v\n", results)
	}

	if results.ResultA && results.ResultGlue && results.ResultNS && results.ResultAccess {
		fmt.Println("OK")
	} else if *suggest {
		suggestions.Suggest(coll, results, cluster)
	} else {
		fmt.Println("Error - run with -debug for more information")

	}
}
