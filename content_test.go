package strfs_test

import (
	"github.com/reiver/go-strfs"

	"io"

	"testing"
)

func TestContent(t *testing.T) {

	tests := []struct{
		Content string
	}{
		{
			Content: "",
		},



		{
			Content: "once",
		},
		{
			Content: "once twice",
		},
		{
			Content: "once twice thrice",
		},
		{
			Content: "once twice thrice fource",
		},
	}

	for testNumber, test := range tests {

		var content strfs.Content
		{
			if strfs.EmptyContent() != content {
				t.Errorf("For test #%d, expected content to be empty but actually wasn't.", testNumber)
				t.Logf("CONTENT: %q", test.Content)
				continue
			}
			if expected, actual := true, content.Closed(); expected != actual {
				t.Errorf("For test #%d, expected content to be closed but actually wasn't.", testNumber)
				t.Logf("EXPECTED CONTENT-CLOSED: %t", expected)
				t.Logf("ACTUAL   CONTENT-CLOSED: %t", actual)
				t.Logf("CONTENT: %q", test.Content)
				continue
			}
			if expected, actual := int64(0), content.Size(); expected != actual {
				t.Errorf("For test #%d, expected content to be zero but actually wasn't.", testNumber)
				t.Logf("EXPECTED CONTENT-SIZE: %d", expected)
				t.Logf("ACTUAL   CONTENT-SIZE: %d", actual)
				t.Logf("CONTENT: %q", test.Content)
				continue
			}
			{
				var b [1]byte
				var p []byte = b[:]

				n, err := content.Read(p)
				if expected, actual := 0, n; expected != actual {
					t.Errorf("For test #%d, expected number of bytes read to zero but actually wasn't.", testNumber)
					t.Logf("EXPECTED NUMBER-BYTES-READ: %d", expected)
					t.Logf("ACTUAL   NUMBER-BYTES-READ: %d", actual)
					t.Logf("CONTENT: %q", test.Content)
					continue
				}
				if nil == err {
					t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
					t.Logf("CONTENT: %q", test.Content)
					continue
				}
				if expected, actual := "closed", err.Error(); expected != actual {
					t.Errorf("For test #%d, the actual error was not what was expected.", testNumber)
					t.Logf("EXPECTED ERROR: %q", expected)
					t.Logf("ACTUAL   ERROR: %q", actual)
					t.Logf("CONTENT: %q", test.Content)
					continue
				}
			}
		}
		{

			content = strfs.CreateContent(test.Content)
			if strfs.EmptyContent() == content {
				t.Errorf("For test #%d, did not expect content to be nil but actually was.", testNumber)
				t.Logf("CONTENT: %q", test.Content)
				continue
			}
			if expected, actual := false, content.Closed(); expected != actual {
				t.Errorf("For test #%d, did not expect content to be closed but actually was.", testNumber)
				t.Logf("EXPECTED CONTENT-CLOSED: %t", expected)
				t.Logf("ACTUAL   CONTENT-CLOSED: %t", actual)
				t.Logf("CONTENT: %q", test.Content)
				continue
			}
			if expected, actual := int64(len(test.Content)), content.Size(); expected != actual {
				t.Errorf("For test #%d, the actual content-size was not what was expected", testNumber)
				t.Logf("EXPECTED CONTENT-SIZE: %d", expected)
				t.Logf("ACTUAL   CONTENT-SIZE: %d", actual)
				t.Logf("CONTENT: %q", test.Content)
				continue
			}
		}

		{
			var reader io.Reader = &content

			actualBytes, err := io.ReadAll(reader)
			if nil != err {
				t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
				t.Logf("ERROR: (%T) %s", err, err)
				t.Logf("FILE-CONTENT: %q", test.Content)
				continue
			}

			var actual   string = string(actualBytes)
			var expected string = test.Content

			if expected != actual {
				t.Errorf("For test #%d, the actual content is not what was expected.", testNumber)
				t.Logf("EXPECTED FILE-CONTENT: %q", expected)
				t.Logf("ACTUAL   FILE-CONTENT: %q", actual)
				continue
			}
		}

		{
			if expected, actual := false, content.Closed(); expected != actual {
				t.Errorf("For test #%d, did not expect content to be closed but actually was." ,testNumber)
				t.Logf("EXPECTED CONTENT-CLOSED: %t", expected)
				t.Logf("ACTUAL   CONTENT-CLOSED: %t", actual)
				continue
			}
		}
		{
			var expected int64 = int64(len(test.Content))
			var actual   int64 = content.Size()

			if expected != actual {
				t.Errorf("For test #%d, the actual content-size is not what was expected.", testNumber)
				t.Logf("EXPECTED FILE-SIZE: %d", expected)
				t.Logf("ACTUAL   FILE-SIZE: %d", actual)
				t.Logf("FILE-CONTENT: %q", test.Content)
				continue
			}
		}
	}
}
