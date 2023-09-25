package internal

import (
	"encoding/json"
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
	return nil
}

type FritzInfo struct {
	Data any `json:"data"`
}

func (f *Fritz) Info() error {
	if err := f.Connect(); err != nil {
		return err
	}
	defer f.Client.session.Close()
	req, _ := f.Client.NewRequest("POST", "data.lua", url.Values{
		"xhr":         {"1"},
		"sid":         {f.Client.session.Sid},
		"lang":        {"de"},
		"page":        {"overview"},
		"xhrId":       {"all"},
		"useajax":     {"1"},
		"no_sidrenew": {""},
	})
	info := FritzInfo{}
	fmt.Println("Getting info...")
	_, err := f.Client.Do(req, &info)
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(info.Data, " ", "  ")
	if err != nil {
		return nil
	}
	fmt.Println(string(data))
	return nil
}
func (f *Fritz) Reconnect() error {
	if err := f.Connect(); err != nil {
		return err
	}
	defer f.Client.session.Close()

	req, _ := f.Client.NewRequest(
		"GET",
		"internet/inetstat_monitor.lua?myXhr=1&action=disconnect&useajax=1&xhr=1&t1695669022799=nocache",
		url.Values{})
	req1, _ := f.Client.NewRequest(
		"GET",
		"internet/inetstat_monitor.lua?myXhr=1&action=connect&useajax=1&xhr=1&t1695669022799=nocache",
		url.Values{})

	fmt.Println("Disconnecting...")
	if _, err := f.Client.Do(req, nil); err != nil {
		return err
	}

	fmt.Println("Connecting...")
	if _, err := f.Client.Do(req1, nil); err != nil {
		return err
	}
	fmt.Println("OK. This can take up to 30s to take full effect..")
	return nil
}

func (f *Fritz) Reboot() error { return nil }
