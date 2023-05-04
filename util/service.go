package util

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"time"
)

func ConvertResultJson(byteValue []byte) (*TestResults, error) {
	var input TestSuites
	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'users' which we defined above
	err := xml.Unmarshal(byteValue, &input)

	if err != nil {
		return nil, err
	}
	// output := TestResults{TestResult{Suites: []TestSuite{}}}
	output := TestResults{}
	var totalfailed, totalpassed, totalskipped int
	for j, suiteInput := range input.TestSuite {
		//	testsuite:=TestResult{Suites: []TestSuite{}}
		startTime, endTime := GetStartAndEndTime(suiteInput.TimeStamp, int(suiteInput.Time))
		newSuite := TestSuite{Cases: []TestCase{}}
		for i, testInput := range suiteInput.Testcases {
			var skipped bool
			var skippedmessage *string
			testCaseStartTime, testCaseEndTime := GetStartAndEndTime(suiteInput.TimeStamp, int(testInput.Time))
			if testInput.Skipped != nil {
				skipped = true
				skippedmessage = &testInput.Skipped.Message
			}
			var stderr, stdout string
			if testInput.SystemErr != nil {
				stderr = testInput.SystemErr.Body
			}
			if testInput.SystemOut != nil {
				stdout = testInput.SystemOut.Body
			}
			var errdetails, errorstacktrace *string
			status := "PASSED"
			//errdetails = new(string)
			//errorstacktrace = "null"

			if testInput.Failure != nil {
				errdetails = &testInput.Failure.Body
				errorstacktrace = &testInput.Failure.Type
				status = "FAIL"
			}

			caseStruct := TestCase{
				ID:              fmt.Sprintf("%d", i+1),
				Name:            testInput.Name,
				Duration:        float64(testInput.Time),
				StartTime:       testCaseStartTime,
				CompletedTime:   testCaseEndTime,
				Skipped:         skipped,
				SkippedMessage:  skippedmessage,
				Status:          status,
				ErrorDetails:    errdetails,
				ErrorStackTrace: errorstacktrace,
				Stderr:          stderr,
				Stdout:          stdout,
			}

			newSuite.Cases = append(newSuite.Cases, caseStruct)
		}
		newSuite.Name = suiteInput.Name
		newSuite.Id = fmt.Sprintf("%d", j+1)
		newSuite.Name = suiteInput.Name
		failed, _ := strconv.Atoi(suiteInput.Failures)
		newSuite.Failed = failed
		skipped, _ := strconv.Atoi(suiteInput.Skipped)
		newSuite.Skipped = skipped
		tests, _ := strconv.Atoi(suiteInput.Tests)
		newSuite.Passed = tests - (skipped + failed)

		newSuite.Duration = suiteInput.Time
		newSuite.StartTime = startTime
		newSuite.CompletedTime = endTime
		output.TestResult.Suites = append(output.TestResult.Suites, newSuite)
		totalfailed += failed
		totalpassed += newSuite.Passed
		totalskipped += skipped
	}
	output.TestResult.Failed = totalfailed
	output.TestResult.Passed = totalpassed
	output.TestResult.Skipped = totalskipped
	// res, err := json.Marshal(output)
	// fmt.Println(output)
	// if err != nil {
	// 	fmt.Println("Error marshalling output json:", err)
	// 	return nil, err
	// }
	// err = ioutil.WriteFile("output.json", res, 0644)
	// if err != nil {
	// 	return nil, err
	// }
	return &output, nil
}

func GetStartAndEndTime(timestamp string, sec int) (int64, int64) {
	//timestampStr := "2021-04-02T15:48:23"

	// Parse the timestamp string into a time.Time value
	t, err := time.Parse("2006-01-02T15:04:05", timestamp)
	if err != nil {
		fmt.Println("Error parsing timestamp:", err)
		//return
	}

	// Get the Unix timestamp
	unixtime := t.Unix()
	//fmt.Println(unixtime)
	// Define the Unix timestamp
	timest := int64(unixtime) // Unix timestamp in seconds

	// Define the total duration of the test suite run in seconds
	duration := int64(sec) // 1 hour

	// Convert the Unix timestamp to a time.Time value
	start := time.Unix(timest, 0)

	// Calculate the end time by adding the duration to the start time
	end := start.Add(time.Duration(duration) * time.Second)

	return timest, end.Unix()

}

func ConvertResultJsonForGitHubAction(byteValue []byte) (*TestResults, error) {

	var input Testsuite
	output := TestResults{}
	// we unmarshal our byteArray which contains our

	// xmlFiles content into 'users' which we defined above

	err := xml.Unmarshal(byteValue, &input)

	if err != nil {

		return nil, err

	}

	var startTime, endTime int64

	newSuite := TestSuite{Cases: []TestCase{}}

	for i, testInput := range input.Testcases {

		var skipped bool

		var skippedmessage *string

		var testCaseStartTime, testCaseEndTime int64 //:= GetStartAndEndTime(suiteInput.TimeStamp, int(testInput.Time))

		if testInput.Skipped != nil {

			skipped = true

			skippedmessage = &testInput.Skipped.Message

		}

		var stderr, stdout string

		if testInput.SystemErr != nil {

			stderr = testInput.SystemErr.Body

		}

		if testInput.SystemOut != nil {

			stdout = testInput.SystemOut.Body

		}

		var errdetails, errorstacktrace *string

		status := "PASSED"

		if testInput.Failure != nil {

			errdetails = &testInput.Failure.Body

			errorstacktrace = &testInput.Failure.Type

			status = "FAIL"

		}

		caseStruct := TestCase{

			ID: fmt.Sprintf("%d", i+1),

			Name: testInput.Name,

			Duration: float64(testInput.Time),

			StartTime: testCaseStartTime,

			CompletedTime: testCaseEndTime,

			Skipped: skipped,

			SkippedMessage: skippedmessage,

			Status: status,

			ErrorDetails: errdetails,

			ErrorStackTrace: errorstacktrace,

			Stderr: stderr,

			Stdout: stdout,
		}

		newSuite.Cases = append(newSuite.Cases, caseStruct)

	}

	newSuite.Name = input.Name

	newSuite.Id = fmt.Sprintf("%d", 1)

	newSuite.Name = input.Name

	failed, _ := strconv.Atoi(input.Failures)

	newSuite.Failed = failed

	skipped, _ := strconv.Atoi(input.Skipped)

	newSuite.Skipped = skipped

	tests, _ := strconv.Atoi(input.Tests)

	newSuite.Passed = tests - (skipped + failed)

	newSuite.Duration = input.Time

	newSuite.StartTime = startTime

	newSuite.CompletedTime = endTime

	output.TestResult.Suites = append(output.TestResult.Suites, newSuite)
	output.TestResult.Failed = failed
	output.TestResult.Passed = newSuite.Passed
	output.TestResult.Skipped = skipped

	return &output, nil

}
