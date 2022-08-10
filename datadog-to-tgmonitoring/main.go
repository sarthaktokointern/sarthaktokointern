package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"strconv"

	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

import (
	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

const IntMin = -2147483647

func main() {

	tgmPolicyName := flag.String("tgm-policy-name", "", "TG-Monitoring policy name")
	tgmPolicyRelativePath := flag.String("tgm-policy-relative-path", "", "TG-Monitoring policy relative path")
	//monitor id of datadog monitor
	MonitorId := flag.Int("monitor-id", IntMin, "Monitor Id of datadog monitor")
	isEnable := flag.Bool("is-enable", false, "State of the datadog monitor (Muted/Not Muted)")
	/*set the nrql query, example of nrql query: from metric select count(test_key.summary) where env = 'production'
	for more info on nrql, go through the following doc: https://docs.newrelic.com/docs/query-your-data/nrql-new-relic-query-language/get-started/nrql-syntax-clauses-functions/
	*/
	query := flag.String("new-relic-query", "", "new relic alert query")
	WarningThreshold := flag.Int("warning-threshold", IntMin, "warning threshold of the alert")

	flag.Parse()

	if *tgmPolicyName == "" {
		log.Fatalln("ERROR: tgm-policy-name is not set or empty")
	}
	if *tgmPolicyRelativePath == "" {
		log.Fatalln("ERROR: tgm-policy-relative-path is not set or empty")
	}
	if *MonitorId == IntMin {
		log.Fatalln("ERROR: monitor id is not set")
	}

	if *query == "" {
		log.Fatalln("new relic alert query not set or empty")
	}

	if *WarningThreshold == IntMin {
		log.Fatalf("warning threshold not set")
	}

	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)

	resp, _, err := apiClient.MonitorsApi.GetMonitor(ctx, int64(*MonitorId), *datadog.NewGetMonitorOptionalParameters())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.GetMonitor`: %v\n", err)

	}

	dirname := *resp.Name

	AlertDuration := getAlertDuration(resp.Query)

	if AlertDuration == 0 {
		log.Fatalf("could not gather duration info from the data dog query")
	}

	AlertOp := getAlertOp(resp.Query)

	if AlertOp == "" {
		log.Fatalf("could not gather alert operation from the data dog query")
	}

	AlertThreshold := getAlertThreshold(resp.Query, AlertOp)

	templ, _ := template.ParseFiles("tgMonitoringFileTemplate.tmpl")

	log.Printf("INFO: creating directory %s", dirname)
	err = os.MkdirAll(dirname, 0755)

	if err != nil {
		log.Fatalf("ERROR: failed in creating directory")
	}

	println("directory created!!")
	func() {
		f, err := os.Create(filepath.Join(dirname, "terragrunt.hcl"))
		if err != nil {
			log.Fatalf("ERROR: failed in creating tg monitoring file under %s for monitor id=%d name=%s: %v", dirname, *resp.Id, *resp.Name, err)
		}
		defer f.Close()
		alertschema := map[string]string{
			"Op":             AlertOp,
			"threshCritical": strconv.Itoa(AlertThreshold),
			"dur":            strconv.Itoa(AlertDuration),
			"isEnable":       strconv.FormatBool(*isEnable),
			"query":          *query,
			"threshWarning":  strconv.Itoa(*WarningThreshold),
		}
		err = renderTGMonitoringFile(*tgmPolicyName, *tgmPolicyRelativePath, &resp, alertschema, templ, f)

		if err != nil {
			log.Fatalln("ERROR: unable to render tg monitoring file:", err)
		}
		log.Println("INFO: done creating", f.Name())

	}()

	log.Printf("INFO: tg-monitoring file for condition belonging to monitorID=%d has been created in directory %s", *resp.Id, dirname)
	log.Println("-------:)-------")

}

func getAlertThreshold(query string, AlertOp string) int {
	threshold := ""
	switch AlertOp {
	case "below or equals":
		st := strings.Index(query, "<=")
		st += 3
		for i := st; i < len(query); i++ {
			threshold = threshold + string(query[i])
		}
		break
	case "below":
		st := strings.Index(query, "<")
		st += 2
		for i := st; i < len(query); i++ {
			threshold = threshold + string(query[i])
		}
		break
	case "above or equals":
		st := strings.Index(query, ">=")
		st += 3
		for i := st; i < len(query); i++ {
			threshold = threshold + string(query[i])
		}
		break
	case "above":
		st := strings.Index(query, ">")
		st += 2
		for i := st; i < len(query); i++ {
			threshold = threshold + string(query[i])
		}
		break
	case "not equals":
		st := strings.Index(query, "!=")
		st += 3
		for i := st; i < len(query); i++ {
			threshold = threshold + string(query[i])
		}
		break
	case "equals":
		st := strings.Index(query, "=")
		st += 2
		for i := st; i < len(query); i++ {
			threshold = threshold + string(query[i])
		}
	}

	thresh := 0
	thresh, _ = strconv.Atoi(threshold)

	return thresh

}

func getAlertOp(query string) string {

	if strings.Contains(query, "<=") {
		return "below or equals"
	} else if strings.Contains(query, "<") {
		return "below"
	} else if strings.Contains(query, ">=") {
		return "above or equals"
	} else if strings.Contains(query, ">") {
		return "above"
	} else if strings.Contains(query, "!=") {
		return "not equals"
	} else if strings.Contains(query, "=") {
		return "not equals"
	}

	return ""

}

func getAlertDuration(query string) int {
	st := strings.Index(query, "last_")
	dur := 0

	if st == -1 {
		log.Fatalf("could not get the alert duration")

	}
	st += 5
	for i := st; i < len(query); i++ {
		if query[i]-'0' <= 9 {
			dur = dur*10 + int(query[i]-'0')
		} else {
			st = i
			break
		}
	}

	switch query[st] {
	case 'd':
		dur = dur * 86400
		break
	case 'm':
		dur = dur * 60
		break
	case 'h':
		dur = dur * 3600
		break

	}
	return dur

}

func renderTGMonitoringFile(tgPolicyName, tgPolicyRelativePath string, resp *datadog.Monitor, alertschema map[string]string, templ *template.Template, w io.Writer) error {

	threshcritical, _ := strconv.Atoi(alertschema["threshCritical"])
	dur, _ := strconv.Atoi(alertschema["dur"])
	isEnable, _ := strconv.ParseBool(alertschema["isEnable"])
	threshWarning, _ := strconv.Atoi(alertschema["threshWarning"])

	return templ.Execute(w, map[string]interface{}{
		"TGMonitoringAlertPolicyName":         tgPolicyName,
		"TGMonitoringAlertPolicyRelativePath": tgPolicyRelativePath,
		"AlertCondition":                      resp,
		"Op":                                  alertschema["Op"],
		"threshCritical":                      threshcritical,
		"dur":                                 dur,
		"isEnable":                            isEnable,
		"query":                               alertschema["query"],
		"threshWarning":                       threshWarning,
	})
}
