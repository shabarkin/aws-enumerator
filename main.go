package main

import (
	"fmt"
	"os"

	"github.com/shabarkin/aws-enumerator/helper"
	"github.com/shabarkin/aws-enumerator/utils"
)

func main() {
	//var flagNoColor *bool

	helper.Cred.Usage = func() {
		fmt.Fprintf(os.Stderr, helper.Cloudrider_cred_help)
	}
	helper.Enum.Usage = func() {
		fmt.Fprintf(os.Stderr, helper.Cloudrider_enum_help)
	}
	helper.Dump.Usage = func() {
		fmt.Fprintf(os.Stderr, helper.Cloudrider_dump_help)
	}

	if len(os.Args) < 2 {
		fmt.Println(helper.Cloudrider_help)
		os.Exit(1)
	}
	switch os.Args[1] {

	case "cred":
		helper.Cred.Parse(os.Args[2:])
		utils.CreateAWScredentialsFile(helper.AWS_region, helper.AWS_access_key_id, helper.AWS_secret_access_key, helper.AWS_session_token)
		fmt.Println(utils.Green("Message: "), utils.Yellow("File"), utils.Red(".env"), utils.Yellow("with AWS credentials were created in current folder"))
	case "enum":
		helper.Enum.Parse(os.Args[2:])
		helper.SetEnumerationPipeline(helper.Services_enum, helper.Speed)
		fmt.Println(utils.Green("Message: "), utils.Yellow("Enumeration finished"))

	case "dump":
		helper.Dump.Parse(os.Args[2:])
		helper.DumpInfo(helper.Services_dump, helper.Print, helper.Filter, helper.Errors_dump)

	default:
		fmt.Println(helper.Cloudrider_help)
		os.Exit(1)
	}
}
