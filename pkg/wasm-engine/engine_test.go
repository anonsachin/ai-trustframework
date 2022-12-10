package wasmengine_test

import (
	wasmengine "ai-trustframework/pkg/wasm-engine"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestRunningWASI(t *testing.T) {
	// base setup
	l := logrus.NewEntry(logrus.StandardLogger())
	engine, err := wasmengine.NewWASIRunner(l)
	assert.Nil(t,err)

	wasmPath := "test-data/echo/target/wasm32-wasi/debug/echo.wasm"

	tt := []struct{
		Name string
		Args []string
		MapingDirextories map[string]string
	}{
		{
			Name:              "Base",
			Args:              []string{"echo"},
			MapingDirextories: map[string]string{},
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name,func(t *testing.T) {
			dir := t.TempDir()
			// the stdout and stderr
			stdout, err := os.CreateTemp(dir,"stdout-*");
			assert.Nil(t,err)
			stderr, err := os.CreateTemp(dir,"stderr-*");
			assert.Nil(t,err)

			// run wasm
			err = engine.Run(wasmPath,stdout.Name(),stderr.Name(),tc.MapingDirextories,tc.Args)

			assert.Nil(t,err)

			if tc.Args != nil {
				str, err := io.ReadAll(stdout)
				assert.Nil(t,err)
				assert.Equal(t,fmt.Sprintf("\"%s\"\n",tc.Args[0]),string(str))
			}
			

		})
	}

}