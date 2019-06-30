package confighelper

import (
	"fmt"
	"log"
	"os"

	"github.com/tkanos/gonfig"
)

type Configuration struct {
	Port              int
	Static_Variable   string
	Connection_String string
}

func GetConfiguration() Configuration {
	basePath, err := os.Getwd()
	configuration := Configuration{}
	err = gonfig.GetConf(basePath+"/confighelper/config.json", &configuration)
	log.Fatal(err)
	fmt.Println(configuration)

	return configuration
}
