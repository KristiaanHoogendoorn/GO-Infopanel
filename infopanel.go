package main

import (
	//"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

const weeronline_API_key string = "2aa4dd2560"

func Get_wether(Location string) string {
	fmt.Println("Starting the application...")
	response, err := http.Get(fmt.Sprintf("http://weerlive.nl/api/json-data-10min.php?key=%s&locatie=%s", weeronline_API_key, Location))
	if err != nil {
		fmt.Printf("Applicatoin failed to connect to API: %s\n", err)
	} else {
		//data := { "liveweer": [{"plaats": "Utrecht", "temp": "9.2", "gtemp": "6.2", "samenv": "Af en toe lichte regen", "lv": "90", "windr": "West", "windms": "6", "winds": "4", "windk": "11.7", "windkmh": "21.6", "luchtd": "1005.5", "ldmmhg": "754", "dauwp": "7", "zicht": "4", "verw": "Geleidelijk enkele buien. Eerste kerstdag nog een enkele bui", "sup": "08:44", "sunder": "16:33", "image": "buien", "d0weer": "halfbewolkt", "d0tmax": "10", "d0tmin": "6", "d0windk": "3", "d0windknp": "10", "d0windms": "5", "d0windkmh": "19", "d0windr": "NO", "d0neerslag": "33", "d0zon": "28", "d1weer": "halfbewolkt", "d1tmax": "9", "d1tmin": "6", "d1windk": "3", "d1windknp": "8", "d1windms": "4", "d1windkmh": "15", "d1windr": "W", "d1neerslag": "50", "d1zon": "30", "d2weer": "halfbewolkt", "d2tmax": "5", "d2tmin": "2", "d2windk": "2", "d2windknp": "6", "d2windms": "3", "d2windkmh": "11", "d2windr": "Z", "d2neerslag": "40", "d2zon": "20", "alarm": "0"}]}

		data, _ := ioutil.ReadAll(response.Body)
		return string(data)
	}
	return ""
}

//func ()  {
//
//}

func main() {
	fmt.Println(Get_wether("Utrecht"))
	//for i := range Get_wether("Utrecht"){
	//	fmt.Println(Get_wether("Utrecht"))
	//	fmt.Println(i)
	//}
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
