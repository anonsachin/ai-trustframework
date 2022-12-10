package ipfs_test

import (
	"ai-trustframework/pkg/ipfs"
	"bytes"
	"io"
	"os/exec"
	"strings"
	"testing"
	// "time"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestIPFSIntegration(t *testing.T) {

	tt := []struct{
		Name string
		Input string
	}{
		{
			Name: "base",
			Input: "Hello",
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name,func(t *testing.T) {
		// 	defer runMakeCommand(t,"ipfs-down")
			
		// 	time.Sleep(20*time.Second)

		// 	runMakeCommand(t,"ipfs-up")

			shell := ipfs.NewShell("localhost:5001",logrus.NewEntry(logrus.StandardLogger()))

			cid, err := shell.Add(strings.NewReader(tc.Input))
			assert.Nil(t,err)

			data, err := shell.Cat(cid)

			if assert.Nil(t,err) {
				defer data.Close()
				b := bytes.NewBuffer(make([]byte, 0))
				_, err = io.Copy(b,data)
				if assert.Nil(t,err) {
					assert.Equal(t,tc.Input,string(b.Bytes()))
				}
			}
		})
	}

}

func runMakeCommand(t *testing.T, command string) {
	out, err := exec.Command("make",command).CombinedOutput()

	if err != nil {
		t.Logf("Unable to run command %v",err)
	}

	t.Logf("The command %s\n%s",command,out)
}