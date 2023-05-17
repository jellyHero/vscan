package main

import (
	"github.com/projectdiscovery/gologger"
	naabuRunner "github.com/veo/vscan/pkg/naabu/v2/pkg/runner"
	"runtime"
)

func main() {
	options := naabuRunner.ParseOptions()

	if runtime.GOOS == "windows" {
		options.NoColor = true
	}
	runner, err := naabuRunner.NewRunner(options)

	if err != nil {
		gologger.Fatal().Msgf("Could not create runner: %s\n", err)
	}
	defer runner.Close()
	err = runner.RunEnumeration()
	if err != nil {
		gologger.Fatal().Msgf("Could not run enumeration: %s\n", err)
	}
	gologger.Info().Msg("Port scan over,web scan starting")
	err = runner.Httpxrun()
	if err != nil {
		gologger.Fatal().Msgf("Could not run httpRunner: %s\n", err)
	}
}

//import (
//	"log"
//
//	naabuRunner "github.com/veo/vscan/pkg/naabu/v2/pkg/runner"
//)
//func main() {
//	options := naabuRunner.SetOptions("www.baidu.com","80","baidu.txt")
//
//
//	//options := naabuRunner.Options{
//	//	Host:      slice,
//	//	ScanType: "s",
//	//	OnResult: func(host string, ip string,port []int) {
//	//		log.Println(ip)
//	//	},
//	//	Ports: "80",
//	//	Rate: 50,
//	//}
//
//	runner, err := naabuRunner.NewRunner(options)
//	if err != nil {
//		log.Println(err)
//	}
//
//	runner.RunEnumeration()
//	defer runner.Close()
//	runner.Httpxrun()
//}
