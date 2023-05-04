package util

import (
	"encoding/xml"
	"fmt"
	"strconv"
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
		newSuite := TestSuite{Cases: []TestCase{}}
		for i, testInput := range suiteInput.Testcases {
			var skipped bool
			var skippedmessage string
			if testInput.Skipped != nil {
				skipped = true
				skippedmessage = testInput.Skipped.Message
			}
			var stderr, stdout string
			if testInput.SystemErr != nil {
				stderr = testInput.SystemErr.Body
			}
			if testInput.SystemOut != nil {
				stdout = testInput.SystemOut.Body
			}
			var errdetails, errorstacktrace string
			status := "PASSED"

			if testInput.Failure != nil {
				errdetails = testInput.Failure.Body
				errorstacktrace = testInput.Failure.Type
				status = "FAIL"
			}

			caseStruct := TestCase{
				ID:              fmt.Sprintf("%d", i+1),
				Name:            testInput.Name,
				Duration:        float64(testInput.Time),
				StartTime:       0,
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
		output.TestResult.Suites = append(output.TestResult.Suites, newSuite)
		totalfailed += failed
		totalpassed += newSuite.Passed
		totalskipped += skipped
	}
	output.Failed = totalfailed
	output.Passed = totalpassed
	output.Skipped = totalskipped
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
