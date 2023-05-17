package main

import (
	naabuRunner "github.com/jellyHero/vscan/pkg/naabu/v2/pkg/runner"
	"github.com/projectdiscovery/gologger"
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
//	"github.com/projectdiscovery/goflags"
//	"log"
//
//	naabuRunner "github.com/jellyHero/vscan/pkg/naabu/v2/pkg/runner"
//)
//
//func main() {
//	host, _ := goflags.ToNormalizedStringSlice("www.baidu.com")
//	options := naabuRunner.Options{
//		Host:              host,
//		Ports:             "80",
//		Output:            "",
//		ExcludeCDN:        false,
//		Threads:           25,
//		Rate:              1000,
//		JSON:              true,
//		CSV:               false,
//		NoPOC:             false,
//		ScanAllIPS:        false,
//		ScanType:          "s",
//		InterfacesList:    false,
//		Nmap:              false,
//		Resume:            false,
//		Stream:            false,
//		Passive:           false,
//		Retries:           3,
//		Timeout:           1000,
//		WarmUpTime:        2,
//		Ping:              false,
//		Verify:            false,
//		Debug:             false,
//		Verbose:           false,
//		NoColor:           false,
//		Silent:            false,
//		Version:           false,
//		EnableProgressBar: false,
//		StatsInterval:     5,
//		Stdin:             false,
//	}
//
//	runner, err := naabuRunner.NewRunner(naabuRunner.SetOptions(&options))
//	if err != nil {
//		log.Println(err)
//	}
//
//	runner.RunEnumeration()
//	defer runner.Close()
//	runner.Httpxrun()
//}
