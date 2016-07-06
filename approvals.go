package ApprovalTests_go

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/Approvals/ApprovalTests_go/reporters"
)

func Verify(t *testing.T, reader io.Reader) error {
	namer, err := getApprovalName()
	if err != nil {
		return err
	}

	reporter := getReporter()
	err = namer.compare(namer.getApprovalFile(".txt"), reader)
	if err != nil {
		reporter.Report(namer.getApprovalFile(".txt"), namer.getReceivedFile(".txt"))
		t.Fail()
	} else {
		os.Remove(namer.getReceivedFile(".txt"))
	}

	return err
}

func VerifyString(t *testing.T, s string) {
	reader := strings.NewReader(s)
	Verify(t, reader)
}

func getReporter() reporters.Reporter {
	return reporters.NewDiffReporter()
}