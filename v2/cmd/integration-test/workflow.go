package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/julienschmidt/httprouter"

	"github.com/onaio/nuclei/v2/core/testutils"
)

var workflowTestcases = map[string]testutils.TestCase{
	"workflow/basic.yaml":               &workflowBasic{},
	"workflow/condition-matched.yaml":   &workflowConditionMatched{},
	"workflow/condition-unmatched.yaml": &workflowConditionUnmatch{},
	"workflow/matcher-name.yaml":        &workflowMatcherName{},
}

type workflowBasic struct{}

// Execute executes a test case and returns an error if occurred
func (h *workflowBasic) Execute(filePath string) error {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		httpDebugRequestDump(r)
		fmt.Fprintf(w, "This is test matcher text")
	})
	ts := httptest.NewServer(router)
	defer ts.Close()

	results, err := testutils.RunNucleiWorkflowAndGetResults(filePath, ts.URL, debug)
	if err != nil {
		return err
	}
	if len(results) != 2 {
		return errIncorrectResultsCount(results)
	}
	return nil
}

type workflowConditionMatched struct{}

// Execute executes a test case and returns an error if occurred
func (h *workflowConditionMatched) Execute(filePath string) error {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		httpDebugRequestDump(r)
		fmt.Fprintf(w, "This is test matcher text")
	})
	ts := httptest.NewServer(router)
	defer ts.Close()

	results, err := testutils.RunNucleiWorkflowAndGetResults(filePath, ts.URL, debug)
	if err != nil {
		return err
	}
	if len(results) != 1 {
		return errIncorrectResultsCount(results)
	}
	return nil
}

type workflowConditionUnmatch struct{}

// Execute executes a test case and returns an error if occurred
func (h *workflowConditionUnmatch) Execute(filePath string) error {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		httpDebugRequestDump(r)
		fmt.Fprintf(w, "This is test matcher text")
	})
	ts := httptest.NewServer(router)
	defer ts.Close()

	results, err := testutils.RunNucleiWorkflowAndGetResults(filePath, ts.URL, debug)
	if err != nil {
		return err
	}
	if len(results) != 0 {
		return errIncorrectResultsCount(results)
	}
	return nil
}

type workflowMatcherName struct{}

// Execute executes a test case and returns an error if occurred
func (h *workflowMatcherName) Execute(filePath string) error {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		httpDebugRequestDump(r)
		fmt.Fprintf(w, "This is test matcher text")
	})
	ts := httptest.NewServer(router)
	defer ts.Close()

	results, err := testutils.RunNucleiWorkflowAndGetResults(filePath, ts.URL, debug)
	if err != nil {
		return err
	}
	if len(results) != 1 {
		return errIncorrectResultsCount(results)
	}
	return nil
}
