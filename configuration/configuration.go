package configuration

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

const configFileLocation string = "./content/data/config.json"
const jsConfigLocation string = "./content/static/js/config.js"

// Contains the configuration information for the system.
type Configuration struct {
	ServerUrl string `json:"serverUrl"`
	HttpPort  int    `json:"httpPort"`
	HttpsPort int    `json:"httpsPort"`
	CertFile  string `json:"certFile"`
	KeyFile   string `json:"keyFile"`
}

// The Load method opens the config.json file for the system, and extracts the configuration
// parameters from the JSON. It then returns this completed configuration. If the config
// file doesn't exist, it is created automatically with default settings. If there is an error
// reading or creating the file, the system will exit with a fatal error.
// After successfully reading (or creating) the config, this method will call sendToJsConfig
// to export the necessary config information to the JavaScript directory.
func Load() *Configuration {
	configFile, err := os.Open(configFileLocation)
	if err != nil {
		newConfigJson, _ := json.Marshal(Configuration{
			ServerUrl: "localhost",
			HttpPort:  8080,
			HttpsPort: 8081,
			CertFile:  "./content/certificates/cert.pem",
			KeyFile:   "./content/certificates/key.pem",
		})
		err = ioutil.WriteFile(configFileLocation, newConfigJson, 0644)
		if err != nil {
			log.Fatalf("Failed to load and write to config file - check permissions")
		}
		configFile, err = os.Open(configFileLocation)
		if err != nil {
			log.Fatalf("Failed to load and write to config file - check permissions")
		}
	}
	defer configFile.Close()

	bytes, err := ioutil.ReadAll(configFile)
	if err != nil {
		log.Fatalf("Failed to read configuration file from %s", configFileLocation)
		// return
	}

	config := Configuration{}

	json.Unmarshal(bytes, &config)

	sendToJsConfig(&config)

	return &config
}

// sendToJsConfig creates a JavaScript file in the specified JavaScript directory,
// containing the config information that is necessary for that JavaScript to run correctly.
// The previously generated JavaScript file is overwritten.
func sendToJsConfig(config *Configuration) {
	jsText := "export const SERVER_URL = 'https://" + config.ServerUrl + ":" + strconv.Itoa(config.HttpsPort) + "';"

	err := ioutil.WriteFile(jsConfigLocation, []byte(jsText), 0644)
	if err != nil {
		log.Fatalf("Failed to write to JS config file - check permissions")
	}
}
