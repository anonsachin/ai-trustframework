package files_test

import (
	"ai-trustframework/pkg/files"
	"path/filepath"
	"strings"
	"testing"
)


func TestCreateFolderStructure(t *testing.T) {

	tt := []struct {
		Name string
		Directories []string
		Files map[string]string
		RootDir string
	}{
		{
			Name: "Base case",
			Directories: []string{"test/input","test/output"},
			Files: map[string]string{
				"test/input/test.tx": "This is a test",
			},
			RootDir: "data",
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name,func(t *testing.T) {
			// cleanup
			defer func(t *testing.T,root string) {
				err := files.RemoveAll(root)
				if err != nil {
					t.Fail()
				}
			}(t,tc.RootDir)
			// start test
	
			for _, dir := range tc.Directories {
				err := files.CreateDir(filepath.Join(tc.RootDir,dir))
				if err != nil {
					t.Fail()
				}
			}

			for fileName, data := range tc.Files {
				data_reader := strings.NewReader(data)
				err := files.CopyToNewFile(data_reader,filepath.Join(tc.RootDir,fileName),0644)
				if err != nil {
					t.Fail()
				}
			}
		})
	}

}

func TestCreateFolderStructureAndTar(t *testing.T) {

	tt := []struct {
		Name string
		Directories []string
		Files map[string]string
		RootDir string
		TarName string
		TarFiles []string
	}{
		{
			Name: "Base case",
			Directories: []string{"test/input","test/output"},
			Files: map[string]string{
				"test/input/test.txt": "This is a test",
			},
			RootDir: "data",
			TarName: "data.tar.gz",
			TarFiles: []string{"test/input/test.txt"},
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name,func(t *testing.T) {
			// cleanup
			defer func(t *testing.T,root string) {
				err := files.RemoveAll(root)
				if err != nil {
					t.Fail()
				}
			}(t,tc.RootDir)
			// start test
	
			for _, dir := range tc.Directories {
				err := files.CreateDir(filepath.Join(tc.RootDir,dir))
				if err != nil {
					t.Fail()
				}
			}

			for fileName, data := range tc.Files {
				data_reader := strings.NewReader(data)
				err := files.CopyToNewFile(data_reader,filepath.Join(tc.RootDir,fileName),0644)
				if err != nil {
					t.Fail()
				}
			}

			filepaths := make([]string,0)

			for _, file := range tc.TarFiles {
				filepaths = append(filepaths, filepath.Join(tc.RootDir,file))
			}

			err := files.CreateTar(filepath.Join(tc.RootDir,tc.TarName),filepaths)
			if err != nil {
				t.Logf("Unable to create tar : %v",err)
				t.Fail()
			}
		})
	}

}

func TestCreateFolderStructureAndTarThenUntar(t *testing.T) {

	tt := []struct {
		Name string
		Directories []string
		Files map[string]string
		RootDir string
		TarName string
		TarFiles []string
		UnTarDir string
	}{
		{
			Name: "Base case",
			Directories: []string{"test/input","test/output"},
			Files: map[string]string{
				"test/input/test.txt": "This is a test",
			},
			RootDir: "data",
			TarName: "data.tar.gz",
			// currently buggy untar, it is not able to untar, if all folder structres are
			// added before hand :- "test/input","test/input/test.txt"
			// just :- "test/input/test.txt" will through an error during untar
			TarFiles: []string{"test/output","test/input","test/input/test.txt"},
			UnTarDir: "UnTar",
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name,func(t *testing.T) {
			// cleanup
			defer func(t *testing.T,root string) {
				err := files.RemoveAll(root)
				if err != nil {
					t.Fail()
				}
			}(t,tc.RootDir)
			// start test
	
			for _, dir := range tc.Directories {
				err := files.CreateDir(filepath.Join(tc.RootDir,dir))
				if err != nil {
					t.Fail()
				}
			}

			for fileName, data := range tc.Files {
				data_reader := strings.NewReader(data)
				err := files.CopyToNewFile(data_reader,filepath.Join(tc.RootDir,fileName),0644)
				if err != nil {
					t.Fail()
				}
			}

			filepaths := make([]string,0)

			for _, file := range tc.TarFiles {
				filepaths = append(filepaths, filepath.Join(tc.RootDir,file))
			}

			err := files.CreateTar(filepath.Join(tc.RootDir,tc.TarName),filepaths)
			if err != nil {
				t.Logf("Unable to create tar : %v",err)
				t.Fail()
			}

			err = files.CreateDir(filepath.Join(tc.RootDir,tc.UnTarDir))
			if err != nil {
				t.Fail()
			}
			err = files.Untar(filepath.Join(tc.RootDir,tc.TarName),filepath.Join(tc.RootDir,tc.UnTarDir))
			if err != nil {
				t.Logf("Unable to untar : %v",err)
				t.Fail()
			}
		})
	}

}