package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	start := time.Now()
	terms := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 
		16, 17, 18,19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
	}

	c := &http.Client{Timeout: time.Millisecond * 15000}

	for i, num := range terms {
		callApi(num, i, c)
	}
	log.Print("Done")
	log.Print(time.Now().Sub(start).Seconds())
}

func callApi(num, id int, c *http.Client) {

	log.Printf("Calling API for id %d", id)

	baseURL := "https://dummy.restapiexample.com/api/v1/employee/%d"

	ur := fmt.Sprintf(baseURL, num)

	req, err := http.NewRequest(http.MethodGet, ur, nil)
	if err != nil {
		//log.Printf("error creating a request for term %d :: error is %+v", num, err)
		return
	}
	res, err := c.Do(req)
	if err != nil {
		//log.Printf("error querying for term %d :: error is %+v", num, err)
		return
	}
	defer res.Body.Close()
	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		//log.Printf("error reading response body :: error is %+v", err)
		return
	}
	//log.Printf("%d  :: ok", id)
}
