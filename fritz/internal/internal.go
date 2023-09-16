package internal

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/user"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

const (
	FritzTcpPort      = 49000
	DefaultConfigPath = ".config/fritz/config.yml"
)

type Config struct {
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Fritz struct {
	Client *Client
	config *Config
}

func NewFritz() (*Fritz, error) {
	usr, err := user.Current()
	if err != nil {
		return nil, err
	}
	file, err := os.ReadFile(filepath.Join(usr.HomeDir, DefaultConfigPath))
	var config Config
	err = yaml.Unmarshal(file, &config)

	baseUrl, err := url.Parse(config.Host)
	if err != nil {
		return nil, err
	}
	client := &Client{
		client:  http.DefaultClient,
		BaseURL: baseUrl,
	}
	return &Fritz{Client: client, config: &config}, nil
}

func (f *Fritz) Connect() error {
	if err := f.Client.Auth(f.config.Username, f.config.Password); err != nil {
		return err
	}
	defer f.Client.session.Close()
	fmt.Printf("Connected to FRITZ!Box: %v \n", f.Client.session.Sid)

    // "DeviceInfo1", "GetInfo"
    // session := f.Client.session
    // req, err := f.Client.NewRequest("POST", "/?lp=reboot", nil)
    // if err != nil {
    // }
	return nil
}

func (f *Fritz) Reconnect() error {
    // TODO: think i need a soap client for this shit
    // service, action
    // self.call_action("WANIPConn1", "ForceTermination")
    return fmt.Errorf("TODO: fritz reconnect")
}
func (f *Fritz) Reboot() error    {
     // self.call_action("DeviceConfig1", "Reboot")
    return fmt.Errorf("TODO: fritz reboot")
}

// func (a *Fritz) reboot() {
// 	fmt.Println("=== Rebooting Router ===")
//
// 	err := a.connect()
// 	if err != nil {
// 		log.Fatalf("Error connecting to FRITZ!Box: %v", err)
// 	}
//
// 	fmt.Printf("Model:   %s\n", a.Client.ModelName())
// 	fmt.Printf("Host:    %s\n", a.Client.Host())
//
// 	fmt.Println("Rebooting router...")
// 	err = a.Client.Reboot()
// 	if err != nil {
// 		log.Fatalf("Error rebooting router: %v", err)
// 	}
//
// 	fmt.Println("Reboot command sent successfully.")
// 	fmt.Println("Please wait for the router to restart.")
// }
//
// func (a *Fritz) reconnect() {
// 	fmt.Println("=== Reconnecting to ISP ===")
//
// 	err := a.connect()
// 	if err != nil {
// 		log.Fatalf("Error connecting to FRITZ!Box: %v", err)
// 	}
//
// 	fmt.Printf("Current Public IP: %s\n", currentIP())
//
// 	fmt.Println("Reconnecting...")
// 	err = a.Client.Reconnect()
// 	if err != nil {
// 		log.Fatalf("Error reconnecting: %v", err)
// 	}
//
// 	fmt.Printf("Reconnection completed. New Public IP: %s\n", currentIP())
// }
//
// func currentIP() string {
// 	res, err := http.Get("https://httpbin.org/ip")
// 	if err != nil {
// 		log.Fatalf("Error getting current IP: %v", err)
// 	}
// 	defer res.Body.Close()
//
// 	body, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		log.Fatalf("Error reading response body: %v", err)
// 	}
//
// 	return string(body)
// }
