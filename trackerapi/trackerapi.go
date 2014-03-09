package trackerapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/user"
)

const URL string = "https://www.pivotaltracker.com/services/v5/me"

var (
	FileLocation = homeDir() + "/.tracker"
	currentUser  = new(pivotalUser)
	Stdout       = os.Stdout
)

func CacheCredentials() error {
	usr, pwd, err := getCredentials()
	apiToken, err := getAPIToken(usr, pwd)
	if err != nil {
		return err
	}

	ioutil.WriteFile(FileLocation, []byte(apiToken), 0644)
	return err
}

func getAPIToken(usr, pwd string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return "", err
	}

	req.SetBasicAuth(usr, pwd)
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var meResp struct {
		APIToken string `json:"api_token"`
	}

	err = json.Unmarshal(body, &meResp)
	if err != nil {
		return "", err
	}
	return meResp.APIToken, nil
}

func getCredentials() (usr, pwd string, err error) {
	fmt.Fprint(Stdout, "Username: ")
	usr, err = readLine()

	if err != nil {
		return
	}

	silenceStty()
	defer unsilenceStty()

	fmt.Fprint(Stdout, "Password: ")
	pwd, err = readLine()

	return usr, pwd, err
}

func homeDir() string {
	usr, _ := user.Current()
	return usr.HomeDir
}
