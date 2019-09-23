package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	//  First lets login to onepeloton

	part1 := "{\"username_or_email\":\""
	pelotonID := "fred"
	part2 := "\",\"password\":\""
	pelotonPass := "abc123"
	part3 := "\"}"

	fullString := part1 + pelotonID + part2 + pelotonPass + part3

	payload := strings.NewReader(fullString)

	url := "https://api.onepeloton.com/auth/login"

	// Set the http headers
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("Content-Type", "text/plain")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Host", "api.onepeloton.com")
	req.Header.Add("Content-Length", "53")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("cache-control", "no-cache")

	var mycookieValue string
	res, _ := http.DefaultClient.Do(req)

	// Get the cookie so we can access my data
	for _, cookie := range res.Cookies() {
		if cookie.Name == "peloton_session_id" {
			mycookieValue = cookie.Value
		}
	}

	//Set the cookie for the next lot of calls
	cookie := http.Cookie{Name: "peloton_session_id", Value: mycookieValue}
	req.AddCookie(&cookie)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	//Get me user details
	url = "https://api.onepeloton.com/api/me"
	req, _ = http.NewRequest("GET", url, payload)
	req.AddCookie(&cookie)

	res, _ = http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ = ioutil.ReadAll(res.Body)

	var myProfile MyProfileStruct
	err := json.Unmarshal(body, &myProfile)
	if err != nil {
		fmt.Println(err)
	}

	//  As we now have my ID,  try and get work out IDs
	url = "https://api.onepeloton.com/api/user/" + myProfile.ID + "/workouts"
	req, _ = http.NewRequest("GET", url, payload)
	req.AddCookie(&cookie)
	res, _ = http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ = ioutil.ReadAll(res.Body)

	var workoutSum WorkoutSummary

	err = json.Unmarshal(body, &workoutSum)
	if err != nil {
		fmt.Println(err)
	}
	// So we now loop through all the workout IDs listed in the workout summary and get the low level detail of each workout
	var count float64
	for _, v := range workoutSum.Data {
		url = "https://api.onepeloton.co.uk/api/workout/" + v.ID + "/performance_graph?every_n=5"
		req, _ = http.NewRequest("GET", url, payload)
		req.AddCookie(&cookie)

		res, _ = http.DefaultClient.Do(req)

		defer res.Body.Close()
		body, _ = ioutil.ReadAll(res.Body)
		var performance PerformanceStruct

		err = json.Unmarshal(body, &performance)
		if err != nil {
			fmt.Println(err)
		}

		//   Now I have access to all the detailed metrics, just for the sake of it, I am getting the calorie count
		for _, v := range performance.Summaries {
			if v.DisplayName == "Calories" {
				count = count + v.Value
				//fmt.Println("Cumulative count :" ,  count)
			}
		}
	}
	fmt.Println("Total calories burned in peloton rides ", count)

}
