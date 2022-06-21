//usr/bin/env go run $0 "$@"; exit
package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	api_key := "f90d99975cf2748d40dd7a5ad7df132d"
	lat := "49.1952"
	lon := "16.608"
	url := "https://api.openweathermap.org/data/2.5/weather?lat=" + lat + "&lon=" + lon + "&appid=" + api_key + "&units=metric&lang=cz"
	path := "/home/antak/.config/conky/Graffias"
	body, stat, err := GetDataFromWeb(url)
	if err != nil {
		log.Println(err)
		return
	}
	if len(body) > 2 || stat == 200 {
		type weather struct {
			Coord struct {
				Lon float64 `json:"lon"`
				Lat float64 `json:"lat"`
			} `json:"coord"`
			Weather []struct {
				Id          int    `json:"id"`
				Main        string `json:"main"`
				Description string `json:"description"`
				Icon        string `json:"icon"`
			} `json:"weather"`
			Base string `json:"base"`
			Main struct {
				Temp      float64 `json:"temp"`
				FeelsLike float64 `json:"feels_like"`
				TempMin   float64 `json:"temp_min"`
				TempMax   float64 `json:"temp_max"`
				Pressure  int     `json:"pressure"`
				Humidity  int     `json:"humidity"`
			} `json:"main"`
			Visibility int `json:"visibility"`
			Wind       struct {
				Speed float64 `json:"speed"`
				Deg   int     `json:"deg"`
			} `json:"wind"`
			Clouds struct {
				All int `json:"all"`
			} `json:"clouds"`
			Dt  int64 `json:"dt"`
			Sys struct {
				Type    int    `json:"type"`
				Id      int    `json:"id"`
				Country string `json:"country"`
				Sunrise int64  `json:"sunrise"`
				Sunset  int64  `json:"sunset"`
			} `json:"sys"`
			Timezone int64  `json:"timezone"`
			Id       int    `json:"id"`
			Name     string `json:"name"`
			Cod      int    `json:"cod"`
			Sunrise  string
			Sunset   string
		}
		data := weather{}
		//str := string(body)
		//fmt.Println(str)
		//var dataw []map[string]interface{}
		err = json.Unmarshal(body, &data)
		if err != nil {
			log.Println("Problem with Unmarshalling dataw to structure.")
			return
		}
		data.Sunrise = GetTimeFromTimestamp(data.Sys.Sunrise + data.Timezone)
		data.Sunset = GetTimeFromTimestamp(data.Sys.Sunset + data.Timezone)
		file, _ := json.MarshalIndent(data, "", " ")
		WriteJSON("weather", path, &file)
	}
	GetPrices()
}
func GetDataFromWeb(url string) (body []byte, status int, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Problem with creating a request.")
		return
	}
	req.Header.Add("Accept", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Problem with connections to API server.")
		return
	}

	defer func(Body io.ReadCloser, err *error) {
		*err = Body.Close()
		if *err != nil {
			log.Println(*err)
		}
	}(res.Body, &err)
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("Problem with reading file")
		return
	}
	if err != nil {
		return
	}
	return body, res.StatusCode, nil
}
func WriteJSON(fileName, path string, file *[]byte) {
	FindFolder(path, true)
	err := ioutil.WriteFile(path+"/"+fileName+".json", *file, 0644)
	if err != nil {
		log.Println("File wasn't created make sure you have path or filename in correct order.")
	}
}

func FindFolder(path string, create bool) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		if create {
			err = os.Mkdir(path, 0755)
			if err != nil {
				log.Fatal(err)
			}
			log.Println("folder does not exist in /" + path + "/ but was created.")
		}
		return false
	}
	return true
	//log.Println("File does exist in /" + path + "/.")
}

func GetTimeFromTimestamp(ti int64) string {
	t := time.Unix(ti, 0).UTC()
	return GetTimeFormat(int64(t.Hour()), int64(t.Minute()), int64(t.Second()))
}
func GetTimeFormat(hour, min, sec int64) string {
	var hou, minu, se string

	hou = strconv.FormatInt(hour, 10)
	minu = strconv.FormatInt(min, 10)
	se = strconv.FormatInt(sec, 10)
	return hou + ":" + minu + ":" + se
}
func GetPrices() {
	url := "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin%2Cethereum%2Cavalanche-2%2Ccardano%2Csolana%2Cpolkadot&vs_currencies=eur%2Cbtc"
	path := "/home/antak/.config/conky/Graffias"
	body, stat, err := GetDataFromWeb(url)
	if err != nil {
		log.Println(err)
		return
	}
	if stat == 200 {
		type price struct {
			Polkadot struct {
				Eur float64 `json:"eur"`
				Btc float64 `json:"btc"`
			} `json:"polkadot"`
			Bitcoin struct {
				Eur float64 `json:"eur"`
				Btc float64 `json:"btc"`
			} `json:"bitcoin"`
			Ethereum struct {
				Eur float64 `json:"eur"`
				Btc float64 `json:"btc"`
			} `json:"ethereum"`
			Cardano struct {
				Eur float64 `json:"eur"`
				Btc float64 `json:"btc"`
			} `json:"cardano"`
			Avalanche2 struct {
				Eur float64 `json:"eur"`
				Btc float64 `json:"btc"`
			} `json:"avalanche-2"`
			Solana struct {
				Eur float64 `json:"eur"`
				Btc float64 `json:"btc"`
			} `json:"solana"`
		}
		data := price{}
		//str := string(body)
		//fmt.Println(str)
		err = json.Unmarshal(body, &data)
		if err != nil {
			log.Println("Problem with Unmarshalling dataw to structure.")
			return
		}
		file, _ := json.MarshalIndent(data, "", " ")
		WriteJSON("prices", path, &file)
	}
}
