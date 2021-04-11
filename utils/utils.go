package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
	terminal "github.com/wayneashleyberry/terminal-dimensions"
)

const FILEPATH string = "./enum-results/"
const ERROR_FILEPATH string = "./enum-results/errors/"

// Colors
var Yellow = color.New(color.FgYellow).SprintFunc()
var Red = color.New(color.FgRed).SprintFunc()
var White = color.New(color.FgWhite).SprintFunc()
var Cyan = color.New(color.FgCyan).SprintFunc()
var Green = color.New(color.FgGreen).SprintFunc()
var Magenta = color.New(color.FgMagenta).SprintFunc()

func countPartWidth(terminal_width, service int) (n int) {
	if terminal_width%2 == 0 {
		if service%2 == 0 {
			n = (int(terminal_width) - service - 2) / 2
		} else {
			n = (int(terminal_width) - service - 2 - 1) / 2
		}
	} else {
		if service%2 == 0 {
			n = (int(terminal_width) - service - 2 - 1) / 2
		} else {
			n = (int(terminal_width) - service - 2) / 2
		}
	}
	return n
}

func PrintDividedLine(service string) {

	terminal_width, _ := terminal.Width()
	n := countPartWidth(int(terminal_width), len(service))

	fmt.Println()
	if service == "" {
		fmt.Println(White(strings.Repeat("-", int(terminal_width))))
	} else {
		fmt.Println(White(strings.Repeat("-", n)), Magenta(strings.ToUpper(service)), White(strings.Repeat("-", n)))
	}
}

func ProcessServiceArgument(str string) []string {

	s := strings.Split(str, ",")

	if len(s) > 10 {
		fmt.Println(Red("Error:"), Yellow("Provide less than 10 services or specify `all` for total enumeration"))
		os.Exit(1)
	} else if len(s) <= 0 {
		fmt.Println(Red("Error:"), Yellow("Provide at least 1 service or specify `all` for total enumeration"))
		os.Exit(1)
	}

	s = ToLowerServices(s)
	s = RemoveDuplicateValues(s)
	return s
}

func ServiceNames() []string {
	service_names := []string{"acm", "amplify", "apigateway", "appmesh", "appsync", "athena", "autoscaling", "backup", "batch", "chime", "cloud9", "clouddirectory", "cloudformation", "cloudfront", "cloudhsm", "cloudhsmv2", "cloudsearch", "cloudtrail", "codebuild", "codecommit", "codedeploy", "codepipeline", "codestar", "comprehend", "datapipeline", "datasync", "dax", "devicefarm", "directconnect", "dlm", "dynamodb", "ec2", "ecr", "ecs", "eks", "elasticache", "elasticbeanstalk", "elastictranscoder", "firehose", "fms", "fsx", "gamelift", "globalaccelerator", "glue", "greengrass", "guardduty", "health", "iam", "inspector", "iot", "iotanalytics", "kafka", "kinesis", "kinesisanalytics", "kinesisvideo", "kms", "lambda", "lightsail", "machinelearning", "macie", "mediaconnect", "mediaconvert", "medialive", "mediapackage", "mediastore", "mediatailor", "mobile", "mq", "opsworks", "organizations", "pinpoint", "polly", "pricing", "ram", "rds", "redshift", "rekognition", "robomaker", "route53", "route53domains", "route53resolver", "s3", "sagemaker", "secretsmanager", "securityhub", "servicecatalog", "shield", "signer", "sms", "snowball", "sns", "sqs", "ssm", "storagegateway", "sts", "support", "transcribe", "transfer", "translate", "waf", "workdocs", "worklink", "workmail", "workspaces", "xray"}
	return service_names
}

func Find(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func RemoveDuplicateValues(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func ToLowerServices(a []string) (services_lowercase []string) {
	for _, n := range a {
		services_lowercase = append(services_lowercase, strings.ToLower(n))
	}
	return services_lowercase
}

func CheckEnvFileExistance() bool {
	if _, err := os.Stat(".env"); err != nil {
		if os.IsNotExist(err) {
			fmt.Println(Red("Error:"), Yellow("File .env does not exist"))
			fmt.Println(Green("Fix:"), Yellow("use `./cloudrider cred -h` command"))
			//fmt.Println(Red("Trace:"), Yellow(err))
			os.Exit(1)
			return false
		}
	}
	return true
}

func CheckFileExistance(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			fmt.Println(Red("Error:"), Yellow("File"), Yellow(path), Yellow("does not exist"))
			fmt.Println(Green("Fix:"), Yellow("DB does not exist, skip this service"))
			//fmt.Println(Red("Trace:"), Yellow(err))
			return false
		}
	}
	return true
}

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(Red("Error:"), Yellow("loading `.env` file"))
		os.Exit(1)
	}
}

func PackResponse(response interface{}) string {
	obj, err := json.Marshal(response)
	if err != nil {
		fmt.Println(Red("Error:"), Yellow("Marshal error, can't create a map:"))
		fmt.Println(Green("Fix:"), Yellow("The problem should be on our side, contact support please"))
		fmt.Println(Red("Trace:"), Yellow(err))
		os.Exit(1)
	}
	return string(obj)
}

func PrintPrettifiedJson(to_prettify string) (string, error) {
	var prettified bytes.Buffer
	err := json.Indent(&prettified, []byte(to_prettify), "", "\t")

	if err != nil {
		return "", err
	}

	return prettified.String(), err
}

func GetJsonKey(m string) (key interface{}) {
	var l map[string]interface{}
	json.Unmarshal([]byte(m), &l)

	p := reflect.ValueOf(l).MapRange()

	for p.Next() {

		key = reflect.ValueOf(p.Key()).Interface()

	}
	return key
}

func GetJsonFromFile(filepath string) (result map[string]interface{}) {
	jsonFile, _ := os.Open(filepath)
	byteValue, _ := ioutil.ReadAll(jsonFile)

	err := json.Unmarshal(byteValue, &result)
	if err != nil {
		fmt.Println(Red("Error:"), Yellow("UnMarshal error, can't recreate an object"))
		fmt.Println(Green("Fix:"), Yellow("The problem should be on our side, contact support please"))
		fmt.Println(Red("Trace:"), Yellow(err))
	}

	return result
}

func AnalyseService(service string, print bool, filter string, errors_dump bool) {

	// Error or result analyse
	var filepath string
	if !errors_dump {
		filepath = FILEPATH + service + ".json"
	} else {
		filepath = ERROR_FILEPATH + service + "_errors.json"
	}

	// Display
	PrintDividedLine(strings.ToUpper(service))
	fmt.Println("")

	// Check the file existance
	if CheckFileExistance(filepath) {

		// dump service db_file
		result := GetJsonFromFile(filepath)

		// check whether any info was stored
		if len(result) == 0 {
			fmt.Println(Red("Error:"), Yellow("No entries in provided service"))
		} else {
			s := reflect.ValueOf(result[service])

			// go through apicalls
			for i := 0; i < s.Len(); i++ {

				m := s.Index(i).Interface().(string)
				key := fmt.Sprintf("%v", GetJsonKey(m))
				if strings.Contains(key, filter) {

					// Display key
					fmt.Println(Cyan(key))

					if print {

						json, err := PrintPrettifiedJson(m)
						if err != nil {
							fmt.Println(err)
						}

						// Display
						fmt.Println(White(json))
						PrintDividedLine("")
						fmt.Println("")
					}
				}
			}
		}

	}
}

func CreateAWScredentialsFile(aws_region, aws_access_key_id, aws_secret_access_key, aws_session_token *string) {

	awsCredentials := "AWS_REGION=" + *aws_region + "\n"
	awsCredentials += "AWS_ACCESS_KEY_ID=" + *aws_access_key_id + "\n"
	awsCredentials += "AWS_SECRET_ACCESS_KEY=" + *aws_secret_access_key + "\n"
	awsCredentials += "AWS_SESSION_TOKEN=" + *aws_session_token + "\n"

	ioutil.WriteFile(".env", []byte(awsCredentials), 0644)
}

// func LoadAWSCreds() {
// 	ioutil.ReadFile()
// }
