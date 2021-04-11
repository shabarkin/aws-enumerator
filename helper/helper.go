package helper

import (
	"flag"
	"fmt"

	"github.com/aws-enumerator/servicemaster"
	"github.com/aws-enumerator/servicestructs"
	"github.com/aws-enumerator/utils"
)

func changeSpeedForTime(speed string) (time int) {
	if speed == "fast" {
		time = 2000
	} else if speed == "slow" {
		time = 4000
	} else {
		time = 3000
	}
	return time
}

func SetEnumerationPipeline(services, speed *string) {
	// Load Global Variables from file
	utils.LoadEnv()

	if servicemaster.CheckAWSCredentials() {

		servicemaster.ServiceCall(
			servicestructs.GetServices(),
			utils.ProcessServiceArgument(*services),
			changeSpeedForTime(*speed),
		)

	}
}

func DumpInfo(services_dump *string, print_apicalls *bool, filter *string, errors_dump *bool) {
	if *services_dump == "all" {
		for _, service := range utils.ServiceNames() {
			utils.AnalyseService(service, *print_apicalls, *filter, *errors_dump)
		}
		fmt.Println()
	} else {
		for _, service := range utils.ProcessServiceArgument(*services_dump) {
			utils.AnalyseService(service, *print_apicalls, *filter, *errors_dump)
		}
		fmt.Println()
	}
}

var Cloudrider_help string = `Usage:
cloudrider [command]

Available Commands:
  cred         Saves AWS credentials to .env file in a propriate format for later authentication
  enum         Enumerates AWS Services with given AWS credentials, it can enumerate all services or target specific services
  dump         Displays gathered information about AWS Services available for a AWS account
  
Flags:
  -h		   Help for cloudride
`

var Cloudrider_cred_help string = `Usage:
cloudrider cred [command]
	  
Flags:
  -aws_region               Specify AWS region to which the account may have access to.
  -aws_access_key_id        Specify AWS Access Key for wanted account.
  -aws_secret_access_key    Specify AWS Secret Access Key for wanted account.
  -aws_session_token        Specify AWS Session Token for wanted account.

Example:
  ./cloudrider cred -aws_region us-west-2 -aws_access_key_id AKIA85CEHPO3GLIABKZD -aws_secret_access_key LW3bDF8xJvzGgArqMo0h4kuCYsnubU23kGICGp/p -aws_session_token LW3bDF8xJvzGgArqM.......
`

var Cloudrider_enum_help string = `Usage:
cloudrider enum [command]
	  
Flags:
  -services     Enumerate permissions specifying services divided by comma or 'all' for total enumeration
  -speed        Speed parameter has three defitions : fast or normal or slow (default is normal)

Example:
  ./cloudrider enum -services iam,sts,s3,ec2 -speed normal
  ./cloudrider enum -services all
`

var Cloudrider_dump_help string = `Usage:
cloudrider dump [command]
	  
Flags:
  -services     Enumerate permissions specifying services divided by comma or 'all' for total enumeration
  -filter       Retrieve only wanted API Call Names by filtering the name. Additional flag. (Filters by first specified chars)
  -print        Print data for a specified service API Calls Get the list of accessible apicalls. Additional flag.
  -errors       Analyse errors returned by AWS API for a specific service. Additional flag.

Example:
  #1 
  ./cloudrider dump -services iam
  ./cloudrider dump -services iam -filter List
  ./cloudrider dump -services iam -filter List -print
  ./cloudrider dump -services iam -filter List -print -errors

  #2 
  ./cloudrider dump -services all
  ./cloudrider dump -services all -filter Get
  ./cloudrider dump -services all -filter Get -print
  ./cloudrider dump -services all -filter Get -print -errors
`

// Command line flags:

var Cred *flag.FlagSet = flag.NewFlagSet("cred", flag.ExitOnError)
var AWS_region *string = Cred.String("aws_region", "", "")
var AWS_access_key_id *string = Cred.String("aws_access_key_id", "", "")
var AWS_secret_access_key *string = Cred.String("aws_secret_access_key", "", "")
var AWS_session_token *string = Cred.String("aws_session_token", "", "")

var Enum *flag.FlagSet = flag.NewFlagSet("enum", flag.ExitOnError)
var Services_enum *string = Enum.String("services", "all", "")
var Speed *string = Enum.String("speed", "normal", "")

var Dump *flag.FlagSet = flag.NewFlagSet("dump", flag.ExitOnError)
var Services_dump *string = Dump.String("services", "all", "")
var Errors_dump *bool = Dump.Bool("errors", false, "")
var Print *bool = Dump.Bool("print", false, "")
var Filter *string = Dump.String("filter", "", "")
