package util

import "encoding/xml"

type TestResults struct {
	TestResult TestResult `json:"testResults"`
}
type TestResult struct {
	Suites  []TestSuite `json:"suites"`
	Failed  int         `json:"failed"`
	Passed  int         `json:"passed"`
	Skipped int         `json:"skipped"`
}

// JSON Structs
type TestCase struct {
	ID              string  `json:"id"`
	Name            string  `json:"name"`
	Duration        float64 `json:"duration"`
	StartTime       int64   `json:"startTime"`
	CompletedTime   int64   `json:"completedTime"`
	ErrorDetails    *string `json:"errorDetails"`
	ErrorStackTrace *string `json:"errorStackTrace"`
	Skipped         bool    `json:"skipped"`
	SkippedMessage  *string `json:"skippedMessage"`
	Status          string  `json:"status"`
	Stderr          string  `json:"stderr"`
	Stdout          string  `json:"stdout"`
}

type TestsuiteInput struct {
	FileName   string
	XMLName    xml.Name   `xml:"testsuite"`
	Name       string     `xml:"name,attr"`
	Time       float64    `xml:"time,attr"`
	Tests      string     `xml:"tests,attr"`
	Errors     string     `xml:"errors,attr"`
	Skipped    string     `xml:"skipped,attr"`
	Failures   string     `xml:"failures,attr"`
	Testcases  []Testcase `xml:"testcase"`
	Properties []Property `xml:"properties>property"`
}
type TestSuite struct {
	Cases           []TestCase `json:"cases"`
	Duration        float64    `json:"duration"`
	StartTime       int64      `json:"startTime"`
	CompletedTime   int64      `json:"completedTime"`
	ErrorDetails    *string    `json:"errorDetails"`
	ErrorStackTrace *string    `json:"errorStackTrace"`
	Id              string     `json:"id"`
	Name            string     `json:"name"`
	Stderr          string     `json:"stderr"`
	Stdout          string     `json:"stdout"`
	Failed          int        `json:"failed"`
	Passed          int        `json:"passed"`
	Skipped         int        `json:"skipped"`
}

// XML Structs input
type TestSuites struct {
	//	XMLName xml.Name `xml:"testsuites"`

	TestSuite []Testsuite `xml:"testsuite"`
}

type Testsuite struct {
	FileName   string
	XMLName    xml.Name   `xml:"testsuite"`
	Name       string     `xml:"name,attr"`
	Time       float64    `xml:"time,attr"`
	Tests      string     `xml:"tests,attr"`
	Errors     string     `xml:"errors,attr"`
	Skipped    string     `xml:"skipped,attr"`
	Failures   string     `xml:"failures,attr"`
	Testcases  []Testcase `xml:"testcase"`
	Properties []Property `xml:"properties>property"`
	TimeStamp  string     `xml:"timestamp,attr"`
}

type Testcase struct {
	Name         string        `xml:"name,attr"`
	Classname    string        `xml:"classname,attr"`
	Time         float64       `xml:"time,attr"`
	Skipped      *Skipped      `xml:"skipped,omitempty"`
	Failure      *Failure      `xml:"failure,omitempty"`
	RerunFailure *RerunFailure `xml:"rerunFailure,omitempty"`
	FlakyFailure *FlakyFailure `xml:"flakyFailure,omitempty"`
	TestError    *TestError    `xml:"error,omitempty"`
	RerunError   *RerunError   `xml:"rerunError,omitempty"`
	FlakyError   *FlakyError   `xml:"flakyError,omitempty"`
	SystemOut    *SystemOut    `xml:"system-out,omitempty"`
	SystemErr    *SystemErr    `xml:"system-err,omitempty"`
}

type Skipped struct {
	Message string `xml:"message,attr"`
}

type Failure struct {
	Message string `xml:"message,attr"`

	Type string `xml:"type,attr"`

	Body string `xml:",cdata"`
}

type RerunFailure struct {
	Message string `xml:"message,attr"`

	Type string `xml:"type,attr"`

	Body string `xml:",cdata"`
}

type FlakyFailure struct {
	Message string `xml:"message,attr"`

	Type string `xml:"type,attr"`

	Body string `xml:",cdata"`
}

type TestError struct {
	Message string `xml:"message,attr"`

	Type string `xml:"type,attr"`

	Body string `xml:",cdata"`
}

type RerunError struct {
	Message string `xml:"message,attr"`
	Type    string `xml:"type,attr"`
	Body    string `xml:",cdata"`
}
type FlakyError struct {
	Message string `xml:"message,attr"`
	Type    string `xml:"type,attr"`
	Body    string `xml:",cdata"`
}
type Property struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}
type SystemOut struct {
	Body string `xml:",cdata"`
}
type SystemErr struct {
	Body string `xml:",cdata"`
}

/*type Testsuites struct {
	Testsuites []Testsuite  `xml:"testsuites"`
}


type Testsuite struct {
	XMLName                   xml.Name `xml:"testsuite"`
	Text                      string   `xml:",chardata"`
	Xsi                       string   `xml:"xsi,attr"`
	NoNamespaceSchemaLocation string   `xml:"noNamespaceSchemaLocation,attr"`
	Name                      string   `xml:"name,attr"`
	Time                      string   `xml:"time,attr"`
	Tests                     string   `xml:"tests,attr"`
	Errors                    string   `xml:"errors,attr"`
	Skipped                   string   `xml:"skipped,attr"`
	Failures                  string   `xml:"failures,attr"`
	Properties                struct {
		Text     string `xml:",chardata"`
		Property []struct {
			Text  string `xml:",chardata"`
			Name  string `xml:"name,attr"`
			Value string `xml:"value,attr"`
		} `xml:"property"`
	} `xml:"properties"`
	Testcase []struct {
		Text      string `xml:",chardata"`
		Name      string `xml:"name,attr"`
		Classname string `xml:"classname,attr"`
		Time      string `xml:"time,attr"`
		Skipped   struct {
			Text    string `xml:",chardata"`
			Message string `xml:"message,attr"`
		} `xml:"skipped"`
		SystemOut string `xml:"system-out"`
		Failure   struct {
			Text    string `xml:",chardata"`
			Message string `xml:"message,attr"`
			Type    string `xml:"type,attr"`
		} `xml:"failure"`
	} `xml:"testcase"`
} */
