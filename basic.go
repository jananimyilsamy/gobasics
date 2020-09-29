package main

import (
	"net/http"
	"html/template"
	"fmt"
	"encoding/json"
)
type Corona struct {
	Conditions []interface{} `json:"conditions"`
	Extras     struct {
	} `json:"extras"`
	Question struct {
		Explanation interface{} `json:"explanation"`
		Extras      struct {
		} `json:"extras"`
		Items []struct {
			Choices []struct {
				ID    string `json:"id"`
				Label string `json:"label"`
			} `json:"choices"`
			Explanation string `json:"explanation"`
			ID          string `json:"id"`
			Name        string `json:"name"`
		} `json:"items"`
		Text string `json:"text"`
		Type string `json:"type"`
	} `json:"question"`
	ShouldStop bool `json:"should_stop"`
}

func main(){
	
	weatherFinder := func(w http.ResponseWriter, r *http.Request){

		endpoint := "https://api.infermedica.com/covid19/diagnosis"

		appId := "XXXXXXXX"

		name  := r.FormValue("Name ")

		req := endpoint +name  + "&APPID=" + appId

		fmt.Println("janani")

		resp,er := http.Get(req)

		if er != nil{
			print(er)
		}

		defer resp.Body.Close()

		decoder := json.NewDecoder(resp.Body)

		var Data Corona

		err := decoder.Decode(&Data)

		if err == nil{
			print(err)
		}

		wr,er := template.ParseFiles("weather.html")

		if er != nil{
			print("Something went wrong")
		}else{
			wr.Execute(w,Data)
		}

		fmt.Println(wr)
	}

	Iweather := func(w http.ResponseWriter, r *http.Request){
		t,err := template.ParseFiles("getCity.html")
		if err != nil{
			fmt.Println("panic")
		}
		data := "fine"
		t.Execute(w,data)
	}

	http.HandleFunc("/",Iweather)

	http.HandleFunc("/FindWeather",weatherFinder)

	http.ListenAndServe(":9229",nil)

}
