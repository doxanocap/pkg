package aws

import (
	"github.com/doxanocap/pkg/lg"
	"testing"
)

func Test(t *testing.T) {
	aws := Init()
	err := aws.InitS3()
	if err != nil {
		lg.Fatal(err)
	}
}
