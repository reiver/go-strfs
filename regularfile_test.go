package strfs_test

import (
	"github.com/reiver/go-strfs"

	"io"
	"io/fs"
	"time"

	"testing"
)

func TestRegularFile(t *testing.T) {

	tests := []struct{
		FileContent string
		FileName    string
		FileModTime time.Time
	}{
		{
			FileContent: "",
			FileName:    "empty.txt",
			FileModTime: time.Now(),
		},



		{
			FileContent: "once",
			FileName:    "file1.txt",
			FileModTime: time.Date(2022, 12, 12, 10, 30, 14, 2, time.UTC),
		},
		{
			FileContent: "once twice",
			FileName:    "file2.html",
			FileModTime: time.Date(1984, 01, 14, 9, 10, 11, 12, time.Local),
		},
		{
			FileContent: "once twice thrice",
			FileName:    "file3.gmni",
			FileModTime: time.Date(1974, 12, 18, 4, 5, 6, 7, time.Local),
		},
		{
			FileContent: "once twice thrice fource",
			FileName:    "file4.fngr",
		},
	}


	for testNumber, test := range tests {

		var regularfile strfs.RegularFile
		{
			fileinfo, err := regularfile.Stat()
			if nil == err {
				t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
				t.Logf("REGULARFILE-NAME: %q", test.FileName)
				t.Logf("REGULARFILE-MODTIME: %v", test.FileModTime)
				t.Logf("REGULARFILE-CONTENT: %q", test.FileContent)
				continue
			}
			if expected, actual := "empty content", err.Error(); expected != actual {
				t.Errorf("For test #%d, the actual error is not what was expected.", testNumber)
				t.Logf("EXPECTED ERROR: %q", expected)
				t.Logf("ACTUAL ERR: %q", actual)
				t.Logf("REGULARFILE-NAME: %q", test.FileName)
				t.Logf("REGULARFILE-MODTIME: %v", test.FileModTime)
				t.Logf("REGULARFILE-CONTENT: %q", test.FileContent)
				continue
			}
			if nil != fileinfo {
				t.Errorf("For test #%d, expected returned fileinfo to be nil but actually wasn't.", testNumber)
				t.Logf("FILEINFO: %#v", fileinfo)
				t.Logf("REGULARFILE-NAME: %q", test.FileName)
				t.Logf("REGULARFILE-MODTIME: %v", test.FileModTime)
				t.Logf("REGULARFILE-CONTENT: %q", test.FileContent)
				continue
				continue
			}
		}
		{
			var b [1]byte
			var p []byte = b[:]

			n, err := regularfile.Read(p)
			if expected, actual := 0, n; expected != actual {
				t.Errorf("For test #%d, the actual number-of-bytes read is not what was expected.", testNumber)
				t.Logf("EXPECTED NUMBER-BYTES-READ: %d", expected)
				t.Logf("ACTUAL   NUMBER-BYTES-READ: %d", actual)
				t.Logf("REGULARFILE-NAME: %q", test.FileName)
				t.Logf("REGULARFILE-MODTIME: %v", test.FileModTime)
				t.Logf("REGULARFILE-CONTENT: %q", test.FileContent)
				continue
			}
			if nil == err {
				t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
				t.Logf("REGULARFILE-NAME: %q", test.FileName)
				t.Logf("REGULARFILE-MODTIME: %v", test.FileModTime)
				t.Logf("REGULARFILE-CONTENT: %q", test.FileContent)
				continue
			}
			if expected, actual := "closed", err.Error(); expected != actual {
				t.Errorf("For test #%d, the actual error is not what was expected.", testNumber)
				t.Logf("EXPECTED ERROR: %q", expected)
				t.Logf("ACTUAL ERR: %q", actual)
				t.Logf("REGULARFILE-NAME: %q", test.FileName)
				t.Logf("REGULARFILE-MODTIME: %v", test.FileModTime)
				t.Logf("REGULARFILE-CONTENT: %q", test.FileContent)
				continue
			}
		}
		{
			regularfile = strfs.RegularFile{
				FileContent: strfs.CreateContent(test.FileContent),
				FileName:    test.FileName,
				FileModTime: test.FileModTime,
			}
		}

		var file fs.File = &regularfile
		if nil == file {
			t.Errorf("For test #%d, did not expect file to be nil but actually was.", testNumber)
			t.Logf("REGULARFILE-NAME: %q", test.FileName)
			t.Logf("REGULARFILE-MODTIME: %v", test.FileModTime)
			t.Logf("REGULARFILE-CONTENT: %q", test.FileContent)
			continue
		}

		{
			actualBytes, err := io.ReadAll(file)
			if nil != err {
				t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
				t.Logf("ERROR: (%T) %s", err, err)
				t.Logf("REGULARFILE-NAME: %q", test.FileName)
				t.Logf("REGULARFILE-MODTIME: %v", test.FileModTime)
				t.Logf("REGULARFILE-CONTENT: %q", test.FileContent)
				continue
			}

			var actual   string = string(actualBytes)
			var expected string = test.FileContent

			if expected != actual {
				t.Errorf("For test #%d, the actual file-content is not what was expected.", testNumber)
				t.Logf("EXPECTED FILE-CONTENT: %q", expected)
				t.Logf("ACTUAL   FILE-CONTENT: %q", actual)
				t.Logf("REGULARFILE-NAME: %q", test.FileName)
				t.Logf("REGULARFILE-MODTIME: %v", test.FileModTime)
				t.Logf("REGULARFILE-CONTENT: %q", test.FileContent)
				continue
			}
		}

		var fileinfo fs.FileInfo
		{
			var err error

			fileinfo, err = file.Stat()
			if nil != err {
				t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
				t.Logf("ERROR: (%T) %s", err, err)
				t.Logf("REGULARFILE-NAME: %q", test.FileName)
				t.Logf("REGULARFILE-MODTIME: %v", test.FileModTime)
				t.Logf("REGULARFILE-CONTENT: %q", test.FileContent)
				continue
			}
			if nil == fileinfo {
				t.Errorf("For test #%d, did not expect file-info to be nil but actually was.", testNumber)
				t.Logf("REGULARFILE-NAME: %q", test.FileName)
				t.Logf("REGULARFILE-MODTIME: %v", test.FileModTime)
				t.Logf("REGULARFILE-CONTENT: %q", test.FileContent)
				continue
			}
		}

		{
			var expected string = test.FileName
			var actual   string = fileinfo.Name()

			if expected != actual {
				t.Errorf("For test #%d, the actual file-name is not what was expected.", testNumber)
				t.Logf("EXPECTED FILE-NAME: %q", expected)
				t.Logf("ACTUAL   FILE-NAME: %q", actual)
				t.Logf("REGULARFILE-NAME: %q", test.FileName)
				t.Logf("REGULARFILE-MODTIME: %v", test.FileModTime)
				t.Logf("REGULARFILE-CONTENT: %q", test.FileContent)
				continue
			}
		}

		{
			var expected int64 = int64(len(test.FileContent))
			var actual   int64 = fileinfo.Size()

			if expected != actual {
				t.Errorf("For test #%d, the actual file-size is not what was expected.", testNumber)
				t.Logf("EXPECTED FILE-SIZE: %d", expected)
				t.Logf("ACTUAL   FILE-SIZE: %d", actual)
				t.Logf("REGULARFILE-NAME: %q", test.FileName)
				t.Logf("REGULARFILE-MODTIME: %v", test.FileModTime)
				t.Logf("REGULARFILE-CONTENT: %q", test.FileContent)
				continue
			}
		}

		{
			const modeRegularFile = 0

			var expected fs.FileMode = modeRegularFile
			var actual   fs.FileMode = fileinfo.Mode()

			if expected != actual {
				t.Errorf("For test #%d, the actual file-mode is not what was expected.", testNumber)
				t.Logf("EXPECTED FILE-MODE: %d", expected)
				t.Logf("ACTUAL   FILE-MODE: %d", actual)
				t.Logf("REGULARFILE-NAME: %q", test.FileName)
				t.Logf("REGULARFILE-MODTIME: %v", test.FileModTime)
				t.Logf("REGULARFILE-CONTENT: %q", test.FileContent)
				continue
			}
		}

		{
			var expected bool = false
			var actual   bool = fileinfo.IsDir()

			if expected != actual {
				t.Errorf("For test #%d, the actual file-is-directory is not what was expected.", testNumber)
				t.Logf("EXPECTED FILE-MODE: %t", expected)
				t.Logf("ACTUAL   FILE-MODE: %t", actual)
				t.Logf("REGULARFILE-NAME: %q", test.FileName)
				t.Logf("REGULARFILE-MODTIME: %v", test.FileModTime)
				t.Logf("REGULARFILE-CONTENT: %q", test.FileContent)
				continue
			}
		}

		{
			var expected time.Time = test.FileModTime
			var actual   time.Time = fileinfo.ModTime()

			if !expected.Equal(actual) {
				t.Errorf("For test #%d, the actual mod-time is not what was expected.", testNumber)
				t.Logf("EXPECTED FILE-MOD-TIME: %v", expected)
				t.Logf("ACTUAL   FILE-MOD-TIME: %v", actual)
				t.Logf("REGULARFILE-NAME: %q", test.FileName)
				t.Logf("REGULARFILE-MODTIME: %v", test.FileModTime)
				t.Logf("REGULARFILE-CONTENT: %q", test.FileContent)
				continue
			}
		}

		{
			var expected string = test.FileContent
			var actual   string = fileinfo.Sys().(string)

			if expected != actual {
				t.Errorf("For test #%d, the actual value for sys was not what was expected.", testNumber)
				t.Logf("EXPECTED SYS: %q", expected)
				t.Logf("ACTUAL   SYS: %q", actual)
				continue
			}
		}
	}
}
