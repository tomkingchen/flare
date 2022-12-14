package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

// Shared variable
var ZoneId string

type Client struct {
	hc *http.Client
}

type ZoneSettings struct {
	Result   []map[string]interface{} `json:"result"`
	Success  bool             `json:"success"`
	Errors   []interface{}    `json:"errors"`
	Messages []interface{}    `json:"messages"`
}

type Zone struct {
	ID                  string      `json:"id"`
	Name                string      `json:"name"`
	Status              string      `json:"status"`
	Paused              bool        `json:"paused"`
	Type                string      `json:"type"`
	DevelopmentMode     int         `json:"development_mode"`
	VerificationKey     string      `json:"verification_key,omitempty"`
	CreatedOn           time.Time   `json:"created_on"`
	ActivatedOn         time.Time   `json:"activated_on"`
	Account struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"account"`
	Plan        struct {
		ID                string `json:"id"`
		Name              string `json:"name"`
	} `json:"plan"`
}

type ZoneResults struct {
	Result []Zone `json:"result"`
	ResultInfo struct {
		Page       int `json:"page"`
		PerPage    int `json:"per_page"`
		TotalPages int `json:"total_pages"`
		Count      int `json:"count"`
		TotalCount int `json:"total_count"`
	} `json:"result_info"`
	Success  bool          `json:"success"`
	Errors   []interface{} `json:"errors"`
	Messages []interface{} `json:"messages"`
}

type cred struct {
	ApiEmail string `yaml:"API_EMAIL"`
	ApiKey   string `yaml:"API_KEY"`
}

// Get Cloudflare API Credential
func (c *cred) getCred() *cred {
	homeDir, err := os.UserHomeDir()
	yamlFilePath := homeDir + "/.flare.yaml"
	yamlFile, err := ioutil.ReadFile(yamlFilePath)
	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "flare",
	Short: "Cli tool to interact with Cloudflare API",
	Long: `
Flare is a CLI tool interact with Cloudflare API, the goal of the tool
is to simplify some common tasks involve working with Cloudflare configuration data`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Generic Query API
func fetchAPI(apiUrl string) string {
	var c cred
	ApiCred := c.getCred()
	ApiEmail := ApiCred.ApiEmail
	ApiKey := ApiCred.ApiKey
	req, err := http.NewRequest("GET", apiUrl, nil)
	req.Header.Add("X-Auth-Email", ApiEmail)
	req.Header.Add("X-Auth-Key", ApiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// return json.Unmarshal(responseData, &responseBuffer)
	return string(responseData)
}
