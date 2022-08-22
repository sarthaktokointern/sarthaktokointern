package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	files, err := ioutil.ReadDir("/Users/sarthaksrivastava-mbp/go/src/github.com/tokopedia/tg-monitoring/new-relic/monitors/directorate/tech-graphql/conditions/v0.35.13")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		dirname := "/Users/sarthaksrivastava-mbp/go/src/github.com/tokopedia/tg-monitoring/new-relic/monitors/directorate/tech-graphql/conditions/v0.35.13"
		dirname = dirname + "/" + f.Name()
		fil, er := ioutil.ReadDir(dirname)
		if er != nil {
			log.Fatal(err)
		}
		for _, f2 := range fil {
			original := dirname
			original = original + "/" + f2.Name()
			newPath := dirname
			newPath = dirname + "/" + "traffic"
			e := os.Rename(original, newPath)
			if e != nil {
				continue
			}
			//println(f2.Name())
		}

		println("---------\n")
	}
}
