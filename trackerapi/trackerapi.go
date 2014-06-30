package trackerapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
)

const URL string = "https://www.pivotaltracker.com/services/v5/me"

var FileLocation string = fromHome("/.tracker")

func CacheCredentials() error {
	apiToken, err := getAPIToken()
	if err != nil {
		return err
	}
	fmt.Println(string(apiToken))
	return err
}

func getAPIToken() ([]byte, error) {

	if _, err := os.Stat(FileLocation); err == nil {
		return ioutil.ReadFile(FileLocation)
	}

	usr, pwd, err := getCredentials()
	client := &http.Client{}
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(usr, pwd)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var meResp struct {
		APIToken []byte `json:"api_token"`
	}

	err = json.Unmarshal(body, &meResp)
	if err != nil {
		return nil, err
	}
	ioutil.WriteFile(FileLocation, []byte(meResp.APIToken), 0644)
	return meResp.APIToken, nil
}

func getCredentials() (usr, pwd string, err error) {
	fmt.Fprint(os.Stdout, "Username: ")
	usr, err = readLine()

	if err != nil {
		return
	}

	silenceStty()
	defer unsilenceStty()

	fmt.Fprint(os.Stdout, "Password: ")
	pwd, err = readLine()

	return usr, pwd, err
}

func fromHome(file string) string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	return filepath.Join(usr.HomeDir, file)
}
