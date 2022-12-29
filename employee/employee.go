package employee

import (
	"airstack-tech-assignment/types"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func GetEmployee(id int64) (*types.Employee, error){

	c := &http.Client{Timeout: time.Millisecond * 15000}

	log.Printf("Calling API for id %d", id)

	baseURL := "https://dummy.restapiexample.com/api/v1/employee/%d"

	ur := fmt.Sprintf(baseURL, id)

	req, err := http.NewRequest(http.MethodGet, ur, nil)
	if err != nil {
		//log.Printf("error creating a request for term %d :: error is %+v", num, err)
		return nil, err
	}
	res, err := c.Do(req)
	if err != nil {
		//log.Printf("error querying for term %d :: error is %+v", num, err)
		return nil, err
	}
	defer res.Body.Close()
	getEmployeeApiResponse, err := ioutil.ReadAll(res.Body)
	if err != nil {
		//log.Printf("error reading response body :: error is %+v", err)
		return nil, err
	}

	var employee types.GetEmployeeResponse
	if err := json.Unmarshal(getEmployeeApiResponse, &employee); err != nil {
		return nil, err
	}

	return &employee.Data , nil
}