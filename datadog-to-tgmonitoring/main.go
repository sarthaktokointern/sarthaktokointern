package main

import (
	_ "context"
	"flag"
	_ "fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"

	"log"
	_ "os"
	_ "path/filepath"
	"strings"
	"text/template"
)

import (
	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

const IntMin = -2147483647

func main() {
	tgmPolicyName := flag.String("tgm-policy-name", "gql-alerts", "TG-Monitoring policy name")
	tgmPolicyRelativePath := flag.String("tgm-policy-relative-path", "../../../../policies/v0.35.13/gql-alerts", "TG-Monitoring policy relative path")
	list := []string{"accounts", "accountsapp", "ace", "acemisc", "Affiliate", "appsflyer", "bankingvcc", "Banner", "bms", "brandstore", "broadcasterreport", "byme", "campaign", "cartapp", "category", "categorytx", "cm", "contactus", "content", "creditapplication", "creditcard", "crosssell", "discussion", "donationcpg", "donationorder", "dynamicpdp", "egold", "event", "eventapp", "feeds", "filtron", "flights", "followership", "ftgenie", "fux", "galadriel", "gamenurturing", "gamification", "gamificationengine", "giftcard", "goal", "goldmerchant", "goods", "hades", "hoth", "hotlist", "ims", "inbox", "insight", "insuretech", "kai", "kero", "keroaddr", "kyc", "lending", "localservicescatalog", "localservicesmedia", "localservicestransaction", "mcl", "membership", "merchantvoucher", "mitraapp", "mojito", "mplogistic", "mutualfund", "myorders", "notifapp", "notifier", "notify", "oauth", "ocr", "ongkirappapi", "ongkirappenvoy", "openapi", "orderapp", "oshome", "osmicrosite", "osseller", "otp", "partnerapp", "partnerintapp", "payment", "pdp", "pdpongkirapp", "play", "productreview", "promo", "promosuggestion", "promotioncampaign", "purchaseprotect", "questengine", "r3", "recharge", "reputation", "resolution", "richie", "rolloutmanager", "salamexpquran", "salamreview", "saldo", "saldomitra", "saldoprioritas", "salesforce", "sauron", "sellerdashboard", "sellerinfo", "sellersearch", "seocms", "shipping", "shoppage", "tempo", "tokopoints", "tokoshop", "tome", "tomepdp", "topads", "topbot", "topchat", "travel", "travelcollective", "travelflightdiscovery", "travelhoteldiscovery", "travelhotelfulfillment", "travelhotelpostsales", "travelseo", "traveltraindiscovery", "trigger", "umrahcatalog", "umrahtx", "universesearch", "uploadpedia", "userapp", "vcc", "videosearch", "vision", "vod", "wallet", "walletoauth", "warehouse", "wishlist"}

	for _, val := range list {
		dirname := "/Users/sarthaksrivastava-mbp/go/src/github.com/tokopedia/tg-monitoring/new-relic/monitors/directorate/tech-graphql/conditions/v0.35.13"
		templ, _ := template.ParseFiles("tgMonitoringFileTemplate.tmpl")
		dirname = dirname + "/" + val + "/" + "RPS"
		log.Printf("INFO: creating directory %s", dirname)
		err := os.MkdirAll(dirname, 0755)

		if err != nil {
			log.Fatalf("ERROR: failed in creating directory")
		}
		name := "[GQL][Potential]" + "[" + val + "]" + "Upstream Traffic below threshold"

		func() {
			f, _ := os.Create(filepath.Join(dirname, "terragrunt.hcl"))

			defer f.Close()

			templ.Execute(f, map[string]interface{}{
				"TGMonitoringAlertPolicyName":         tgmPolicyName,
				"TGMonitoringAlertPolicyRelativePath": tgmPolicyRelativePath,
				"AlertCondition":                      name,
				"service":                             strings.ToLower(val),
			})

			log.Println("INFO: done creating", f.Name())

		}()

		log.Println("-------:)-------")

	}

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
