package main

import (
	"io/ioutil"
	"net/http"
	"html/template"
	"fmt"
	"encoding/json"
	"time"
)
type Corona  struct {
	Country     string    `json:"Country"`
	CountryCode string    `json:"CountryCode"`
	Lat         string    `json:"Lat"`
	Lon         string    `json:"Lon"`
	Cases       int       `json:"Cases"`
	Status      string    `json:"Status"`
	Date        time.Time `json:"Date"`
	con   struct {
		India string `json:"India"`
	}
}
func main(){
	disp:=func(w http.ResponseWriter, r *http.Request){



		endpoint := "https://api.covid19api.com/dayone/country/"

		Country := r.FormValue("Country")

		res := endpoint + Country

		req,err:=http.Get(res)

		defer req.Body.Close()

		fmt.Println(Country)

		if(err!=nil){

			fmt.Println(err)
		}
		body, err := ioutil.ReadAll(req.Body)
		var c [] Corona
		err = json.Unmarshal(body, &c)
		fmt.Println(string(body))
		if(err!=nil){

			fmt.Println(err)
		}
		wr,er := template.ParseFiles("corona.html")

		if er != nil{

			fmt.Println(Country)
			fmt.Println(res)
			fmt.Println("wrong something")

		}else{
			wr.Execute(w,c)
		}
		for l := range c {
			fmt.Printf("Id = %v, Name = %v", c[l].Country, c[l].CountryCode)
			fmt.Println()
		}
		fmt.Println(c)

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
