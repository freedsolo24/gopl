package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var m1 = Movie{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}}
var m2 = Movie{Title: "Cool Hand Luke", Year: 1967, Color: true, Actors: []string{"Paul Newman"}}
var m3 = Movie{Title: "Bullitt", Year: 1968, Color: true, Actors: []string{"Steve McQueen", "Jacqueline Bisset"}}

var movies = []*Movie{&m1, &m2, &m3}

func jsonMarshal() {
	data1, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling failed:%s", err)
	}
	fmt.Printf("%s\n", data1)
	// [{"Title":"Casablanca","released":1942,"Actors":["Humphrey Bogart","Ingrid Bergman"]},{"Title":"Cool Hand Luke","released":1967,"color":true,"Actors":["Paul Newman"]},{"Title":"Bullitt","released":1968,"color":true,"Actors":["Steve McQueen","Jacqueline Bisset"]}]
	// [ ... ] 最外层的方括号, 说明整体是一个切片
	// [ {...}, {...}, {...} ] 里面每一个尖括号, 说明是每一个结构体实例
	// [ {k:v, k:v, k:["","",""] }, {...}, {...} ]
	// 每个结构体的实例映射成json的kv对, 最后一个key对应的value是一个方括号, 说明里面是一个切片

	data2, err := json.MarshalIndent(movies, "", "   ")
	if err != nil {
		log.Fatalf("JSON marshaling failed:%s", err)
	}
	fmt.Printf("%s\n", data2)
	// [
	//
	//	{
	//	   "Title": "Casablanca",
	//	   "released": 1942,
	//	   "Actors": [
	//	      "Humphrey Bogart",
	//	      "Ingrid Bergman"
	//	   ]
	//	},
	//	{
	//	   "Title": "Cool Hand Luke",
	//	   "released": 1967,
	//	   "color": true,
	//	   "Actors": [
	//	      "Paul Newman"
	//	   ]
	//	},
	//	{
	//	   "Title": "Bullitt",
	//	   "released": 1968,
	//	   "color": true,
	//	   "Actors": [
	//	      "Steve McQueen",
	//	      "Jacqueline Bisset"
	//	   ]
	//	}
	//
	// ]

	type title struct {
		Title string
		Year  int `json:"released"`
	}
	var titles []title

	if err := json.Unmarshal(data1, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s\n", err)
	}
	fmt.Printf("%v\n", titles)

}
