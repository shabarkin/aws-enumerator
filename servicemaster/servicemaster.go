package servicemaster

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/aws-enumerator/utils"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

type ServiceMaster struct {
	Svc     interface{}
	SvcName string

	ApiCalls           []map[string]interface{}
	json_result_struct map[string][]string
	json_error_struct  map[string][]string

	api_call_result_channel chan string
	api_call_error_channel  chan string

	result_counter int
	error_counter  int
}

func (svc *ServiceMaster) ServiceEnumerator() {
	// initialize counters, channels, result struct
	svc.initialize()

	// Starting go routines (handles 200 goroutines, then rate limit error of AWS, (recommended number of goroutines is less than 100)
	for i := 0; i < len(svc.ApiCalls); i++ {
		go svc.apicall_wrapper(i)
	}

	// Launch control manager for goroutines
	svc.control_node()

	// Save all gathered results to a json file
	svc.save_result_to_file()
	defer wg.Done()
}

func (svc *ServiceMaster) initialize() {
	// reset & init counters
	svc.error_counter = 0
	svc.result_counter = 0

	// reset & init result map
	delete(svc.json_error_struct, "errors")
	svc.json_error_struct = make(map[string][]string)

	// reset & init error map
	delete(svc.json_result_struct, svc.SvcName)
	svc.json_result_struct = make(map[string][]string)

	// reset & init channels
	svc.api_call_result_channel = make(chan string, len(svc.ApiCalls))
	svc.api_call_error_channel = make(chan string, len(svc.ApiCalls))
}

func (svc *ServiceMaster) control_node() {
	// Handling Goroutine results
	for {
		// If all of the api calls were done, close channels and break the loop
		if svc.result_counter >= len(svc.ApiCalls) {
			defer close(svc.api_call_result_channel)
			defer close(svc.api_call_error_channel)
			fmt.Println(utils.Green("Message: "), utils.Yellow("Successful"), utils.Yellow(strings.ToUpper(svc.SvcName))+utils.Yellow(":"), utils.Green(svc.result_counter-svc.error_counter), utils.Yellow("/"), utils.Red(svc.result_counter))
			break
		}

		select {
		// Handling results
		case msg := <-svc.api_call_result_channel:
			// Dumping to the map
			svc.json_result_struct[svc.SvcName] = append(svc.json_result_struct[svc.SvcName], msg)
			svc.result_counter++

		// Handling any kind of errors
		case err := <-svc.api_call_error_channel:
			svc.json_error_struct[svc.SvcName] = append(svc.json_error_struct[svc.SvcName], err)
			svc.result_counter++
			svc.error_counter++
		}
	}
}

func (svc *ServiceMaster) apicall_wrapper(it int) {

	apicall_name := reflect.ValueOf(svc.ApiCalls[it]["apicall"]).Interface().(string)

	s := reflect.ValueOf(svc.Svc).MethodByName(apicall_name).Call(
		[]reflect.Value{
			reflect.ValueOf(context.TODO()),
			reflect.ValueOf(svc.ApiCalls[it]["input_obj"]),
		},
	)
	response := s[0].Interface()
	err := s[1].Interface()

	if err != nil {
		svc.api_call_error_channel <- utils.PackResponse(map[string]string{apicall_name: err.(error).Error()})
	} else {
		svc.api_call_result_channel <- utils.PackResponse(map[string]interface{}{apicall_name: response})
	}
}

// 1. ADD THE FOLDER CREATION IF NOT EXISTED
func (svc *ServiceMaster) save_result_to_file() {

	if _, err := os.Stat(utils.FILEPATH); os.IsNotExist(err) {
		errDir := os.MkdirAll(utils.FILEPATH, 0755)
		if errDir != nil {
			log.Fatalln(utils.Red(err))
		}
	}
	// results
	file_results := utils.PackResponse(svc.json_result_struct)
	ioutil.WriteFile(utils.FILEPATH+svc.SvcName+".json", []byte(file_results), 0644)

	if _, err := os.Stat(utils.ERROR_FILEPATH); os.IsNotExist(err) {
		errDir := os.MkdirAll(utils.ERROR_FILEPATH, 0755)
		if errDir != nil {
			log.Fatalln(utils.Red(err))
		}
	}

	// save errors
	file_errors := utils.PackResponse(svc.json_error_struct)
	ioutil.WriteFile(utils.ERROR_FILEPATH+svc.SvcName+"_errors.json", []byte(file_errors), 0644)
}

func CheckAWSCredentials() bool {
	if utils.CheckEnvFileExistance() {
		cfg, err := config.LoadDefaultConfig(context.TODO())
		if err != nil {
			fmt.Println(utils.Red("Error:"), utils.Yellow("Unable to load SDK config,"))
			fmt.Println(utils.Green("Fix:"), utils.Yellow("The problem should be on our side, contact support please"))
			fmt.Println(utils.Red("Trace:"), utils.Yellow(err))
			os.Exit(1)
		}

		sts_svc := sts.NewFromConfig(cfg)
		_, aws_err := sts_svc.GetCallerIdentity(context.TODO(), &sts.GetCallerIdentityInput{})
		if aws_err != nil {
			fmt.Println(utils.Red("Error:"), utils.Yellow("AWS Credentials are not valid"))
			fmt.Println(utils.Green("Fix:"), utils.Yellow("Provide AWS Credentials, use `./aws-enumerator cred -h` command"))
			fmt.Println(utils.Red("Trace:"), utils.Yellow(aws_err))
			os.Exit(1)
		}
		return true
	} else {
		return false
	}
}

func sleep_delay(i, speed int) {
	if (i+1)%5 == 0 {
		time.Sleep(time.Duration(speed) * time.Millisecond)
	}
}

var wg sync.WaitGroup

func ServiceCall(AllAWSServices []ServiceMaster, wanted_services []string, speed int) {

	start := time.Now()
	if utils.Find(wanted_services, "all") {
		for i := range AllAWSServices {
			wg.Add(1)
			go AllAWSServices[i].ServiceEnumerator()
			sleep_delay(i, speed)
		}
		wg.Wait()

	} else {
		for aws_i := range AllAWSServices {
			for str_i := range wanted_services {

				if AllAWSServices[aws_i].SvcName == wanted_services[str_i] {
					wg.Add(1)
					go AllAWSServices[aws_i].ServiceEnumerator()
					break
				}
			}

		}
		wg.Wait()
	}
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println(utils.Green("Time:"), elapsed)
}
