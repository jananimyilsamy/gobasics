package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"time"
)
type coronaReport struct {
	Global struct {
		NewConfirmed   int `json:"NewConfirmed"`
		TotalConfirmed int `json:"TotalConfirmed"`
		NewDeaths      int `json:"NewDeaths"`
		TotalDeaths    int `json:"TotalDeaths"`
		NewRecovered   int `json:"NewRecovered"`
		TotalRecovered int `json:"TotalRecovered"`
	} `json:"Global"`
	Countries []struct {
		Country        string    `json:"Country"`
		CountryCode    string    `json:"CountryCode"`
		Slug           string    `json:"Slug"`
		NewConfirmed   int       `json:"NewConfirmed"`
		TotalConfirmed int       `json:"TotalConfirmed"`
		NewDeaths      int       `json:"NewDeaths"`
		TotalDeaths    int       `json:"TotalDeaths"`
		NewRecovered   int       `json:"NewRecovered"`
		TotalRecovered int       `json:"TotalRecovered"`
		Date           time.Time `json:"Date"`
	} `json:"Countries"`
	Date time.Time `json:"Date"`
}

func main(){
	disp:=func(w http.ResponseWriter, r *http.Request){



		endpoint := "https://api.covid19api.com/"

		Country := r.FormValue("Country")

		res := endpoint + Country

		req,err:=http.Get(res)

		defer req.Body.Close()

		fmt.Println(Country)

		if(err!=nil){

			fmt.Println(err)
		}



		var c  coronaReport

		decoder:=json.NewDecoder(req.Body)

		er1 := decoder.Decode(&c)
		fmt.Println(c)
		if(er1!=nil){

			fmt.Println(er1)
		}
		wr,er := template.ParseFiles("corona.html")

		if er != nil{
			fmt.Println(res)
			fmt.Println("wrong something")

		}else{
			wr.Execute(w,c)
		}
		fmt.Println(wr)


	}

	input:=func(w http.ResponseWriter, r *http.Request){

		wr,er := template.ParseFiles("getCountry.html")

		if er != nil{

			fmt.Print(" wrong")

		}else{

			wr.Execute(w,nil)

			fmt.Print("redirected")
		}
	}

	http.HandleFunc("/c",input)

	http.HandleFunc("/FindWeather",disp)
	http.ListenAndServe(":9229",nil)
}
