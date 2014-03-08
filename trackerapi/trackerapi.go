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

// MeResponse .
type MeResponse struct {
	APIToken string `json:"api_token"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Initials string `json:"initials"`
	Timezone struct {
		Kind      string `json:"kind"`
		Offset    string `json:"offset"`
		OlsonName string `json:"olson_name"`
	} `json:"time_zone"`
}

func CacheCredentials() error {
	usr, pwd, err := getCredentials()
	parse(makeMeRequest(usr, pwd))
	ioutil.WriteFile(FileLocation, []byte(currentUser.APIToken), 0644)
	return err
}

func makeMeRequest(usr, pwd string) []byte {
	client := &http.Client{}
	req, err := http.NewRequest("GET", URL, nil)
	req.SetBasicAuth(usr, pwd)
	resp, err := client.Do(req)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("\n****\nAPI response: \n%s\n", string(body))
	return body
}

func parse(body []byte) {
	var meResp = new(MeResponse)
	err := json.Unmarshal(body, &meResp)
	if err != nil {
		fmt.Println("error:", err)
	}

	currentUser.APIToken = meResp.APIToken
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
