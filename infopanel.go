package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type weather struct {
	Liveweer []struct {
		Plaats     string
		Temp       string
		Samenv     string
		Lv         string
		Windr      string
		Windms     string
		Winds      string
		Windk      string
		Windkmh    string
		Luchtd     string
		Ldmmhg     string
		Dauwp      string
		Zicht      string
		Verw       string
		Sup        string
		Sunder     string
		Image      string
		D0Weer     string
		D0Tmax     string
		D0Tmin     string
		D0Windk    string
		Gtemp      string
		D0Windknp  string
		D0Windms   string
		D0Windkmh  string
		D0Windr    string
		D0Neerslag string
		Alarm      string
	}
}

var myClient = &http.Client{Timeout: 10 * time.Second}

const weeronline_API_key string = "2aa4dd2560"

func GetWeather(Location string, target interface{}) error {
	fmt.Println("Starting the application...")
	response, err := myClient.Get(fmt.Sprintf("http://weerlive.nl/api/json-data-10min.php?key=%s&locatie=%s", weeronline_API_key, Location))
	if err != nil {
		fmt.Printf("Applicatoin failed to connect to API: %s\n", err)
	}
	defer response.Body.Close()
	return json.NewDecoder(response.Body).Decode(target)
}

func main() {
	data1 := new(weather)
	GetWeather("Utrecht", data1)
	fmt.Println(data1.Liveweer[0].Plaats)
}

//#post data example
//	jsonData := map[string]string{"firstname": "Nic", "lastname": "Raboy"}
//	jsonValue, _ := json.Marshal(jsonData)
//	response, err = http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jsonValue))
//	if err != nil {
//		fmt.Printf("The HTTP request failed with error %s\n", err)
//	} else {
//		data, _ := ioutil.ReadAll(response.Body)
//		fmt.Println(string(data))
//	}
//	fmt.Println("Terminating the application...")
//}
