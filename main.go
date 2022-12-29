package main

import (
	"airstack-tech-assignment/employee"
	"fmt"
	"log"
	"sync"
	"time"
)

const MAX_CONNCURRENT_REQ_EMPLOYEE_API int = 10

func main() {
	start := time.Now()
	employeeIDs := []int64{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
		16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
	}
	processFetchEmpDetailRequest(employeeIDs)
	log.Print("Done")
	log.Print(time.Now().Sub(start).Seconds())
}

func processFetchEmpDetailRequest(employeeIDs []int64) {
	var (
		wg                   sync.WaitGroup
		getEmpDetailsJobChan = make(chan int64, MAX_CONNCURRENT_REQ_EMPLOYEE_API)
		responseChan         = make(chan string, len(employeeIDs))
	)

	wg.Add(MAX_CONNCURRENT_REQ_EMPLOYEE_API)

	// Spinning up the workers
	for i := 0; i < MAX_CONNCURRENT_REQ_EMPLOYEE_API; i++ {
		go worker(getEmpDetailsJobChan, responseChan, &wg)
	}

	// Queueing Jobs in channel
	for i := 0; i < len(employeeIDs); {
		if queueJob(employeeIDs[i], getEmpDetailsJobChan) {
			i++
		}
	}

	close(getEmpDetailsJobChan)

	wg.Wait()

	for noOfReqProcesssed:=0; noOfReqProcesssed<len(employeeIDs); {
		select {
		case response:= <- responseChan:
			fmt.Printf("%+v\n", response)
			noOfReqProcesssed++
		}
	}

}

func queueJob(job int64, jobChan chan<- int64) bool {
	select {
	case jobChan <- job:
		return true
	default:
		return false
	}
}

func worker(jobChan <-chan int64, responseChan chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Worker is waiting for jobs")

	for employeeId := range jobChan {
		employee, err := employee.GetEmployee(employeeId)
		var response string
		if err != nil {
			response = fmt.Sprintf("ERROR: %s", err.Error())
		} else {
			response = fmt.Sprintf("%+v", *employee)
		}
		responseChan <- response
	}
}
