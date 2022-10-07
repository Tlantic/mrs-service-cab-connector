package auth

import (
	"fmt"
	"os"
	"sync"

	"github.com/Tlantic/go-util/v4/mrs/configuration"
	"github.com/Tlantic/go-util/v4/mrs/k8s"
	"github.com/Tlantic/go-util/v4/mrs/log"
	"github.com/Tlantic/go-util/v4/mrs/uuid"
	"gopkg.in/yaml.v2"
)

var (
	instance Values
	once     sync.Once
)

// Options ...
type Options struct {
	ConfigName string
}

// Values ...
type Values struct {
	ID          string            `json:"service_id" yaml:"ServiceId"`
	Name        string            `json:"name" yaml:"Name"`
	Environment string            `json:"environment" yaml:"Environment"`
	Username    string            `json:"username" yaml:"Username"`
	Password    string            `json:"password" yaml:"Password"`
	External    *ExternalEndpoint `json:"external" yaml:"External"`
	Ports       ServerPorts       `json:"server_ports" yaml:"Ports"`
}

// Get ...
func Get(options *Options) *Values {

	once.Do(func() {

		// Options
		if options.ConfigName == "" {
			options.ConfigName = os.Getenv("CONFIG_NAME")
		}

		// Sidecar
		service, err := k8s.NewService(os.Getenv("K8S_SIDECAR"))
		if err != nil {
			log.Fatalf("Could not connect to K8S_SIDECAR: %v", err)
		}
		defer service.Close()

		// Configuration Manager
		cm := configuration.GetManager(&configuration.ProviderOptions{
			Kind:       os.Getenv("CONFIG_PROVIDER"),
			K8sService: service,
		})
		if err := cm.Get(options.ConfigName, &instance, os.Getenv("CONFIG_FORMAT")); err != nil {
			log.Fatalln("Unable to retrieve configuration:", err)
		}

		// Service ID
		if value, ok := os.LookupEnv("SERVICE_ID"); ok {
			instance.ID = value
		} else {
			log.Println("env variable SERVICE_ID not defined: generating new id")
			instance.ID = uuid.New().String()
		}

		// Service Name
		if value, ok := os.LookupEnv("SERVICE_NAME"); ok {
			instance.Name = value
		} else {
			log.Fatalln("env variable SERVICE_NAME not defined")
		}

		// Namespace
		if value, ok := os.LookupEnv("K8S_NAMESPACE"); ok {
			instance.Environment = value
		} else {
			log.Fatalln("env variable K8S_NAMESPACE not defined")
		}

		// CAB variables
		// Credentials
		if value, ok := os.LookupEnv("CLIENT_SYSTEM_USERNAME"); ok {
			instance.Username = value
		} else {
			log.Fatalln("env variable CLIENT_SYSTEM_USERNAME not defined")
		}

		if value, ok := os.LookupEnv("CLIENT_SYSTEM_PASSWORD"); ok {
			instance.Password = value
		} else {
			log.Fatalln("env variable CLIENT_SYSTEM_PASSWORD not defined")
		}

		// Endpoint
		if value, ok := os.LookupEnv("CLIENT_ENDPOINT"); ok {
			instance.External.BaseURI = value
		} else {
			log.Fatalln("env variable CLIENT_ENDPOINT not defined")
		}

		if b, err := yaml.Marshal(instance); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("******************** CONFIG ********************")
			fmt.Println(string(b))
			fmt.Println("************************************************")
		}
	})

	return &instance
}
