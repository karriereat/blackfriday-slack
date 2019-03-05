package slackdown_test

import (
	"testing"

	"github.com/karriereat/blackfriday-slack"
	bf "github.com/russross/blackfriday/v2"
)

type testData struct {
	input      string
	expected   string
	extensions bf.Extensions
}

func runTest(t *testing.T, tdt []testData) {
	for _, v := range tdt {
		renderer := &slackdown.Renderer{}
		md := bf.New(bf.WithRenderer(renderer), bf.WithExtensions(v.extensions))
		ast := md.Parse([]byte(v.input))
		output := string(renderer.Render(ast))

		if output != v.expected {
			t.Errorf("got:%v\nwant:%v", output, v.expected)
		}
	}
}

func TestHeading(t *testing.T) {
	tdt := []testData{
		{input: "# Head1\n", expected: "*Head1*\n", extensions: bf.CommonExtensions},
		{input: "## Head2\n", expected: "*Head2*\n", extensions: bf.CommonExtensions},
		{input: "### Head3\n", expected: "*Head3*\n", extensions: bf.CommonExtensions},
	}

	runTest(t, tdt)
}

func TestCode(t *testing.T) {
	tdt := []testData{
		{
			input:      "this is `foo`.",
			expected:   "this is `foo`.\n\n",
			extensions: bf.CommonExtensions,
		},
	}

	runTest(t, tdt)
}

func TestList(t *testing.T) {
	tdt := []testData{
		{
			input:      "* list1\n* list2\n* list 3\n",
			expected:   " - list1\n - list2\n - list 3\n\n",
			extensions: bf.CommonExtensions,
		},
	}

	runTest(t, tdt)
}

func TestNestedList(t *testing.T) {
	tdt := []testData{
		{
			input:      "* list1\n* list2\n  * list3\n  * list4",
			expected:   " - list1\n - list2\n    - list3\n    - list4\n\n",
			extensions: bf.CommonExtensions,
		},
		{
			input:      "* list1\n* list2\n  * list3\n  * list4\n* list5",
			expected:   " - list1\n - list2\n    - list3\n    - list4\n - list5\n\n",
			extensions: bf.CommonExtensions,
		},
	}

	runTest(t, tdt)
}

func TestOrderedList(t *testing.T) {
	tdt := []testData{
		{
			input:      "1. list1\n2. list2\n3. list3\n4. list4",
			expected:   " 1. list1\n 2. list2\n 3. list3\n 4. list4\n\n",
			extensions: bf.CommonExtensions,
		},
	}

	runTest(t, tdt)
}

func TestDel(t *testing.T) {
	tdt := []testData{
		{
			input:      "~~del text~~",
			expected:   "~del text~\n\n",
			extensions: bf.CommonExtensions,
		},
	}

	runTest(t, tdt)
}

func TestLink(t *testing.T) {
	tdt := []testData{
		{
			input:      "[google.at](http://www.google.at)",
			expected:   "<http://www.google.at|google.at>\n\n",
			extensions: bf.CommonExtensions,
		},
	}

	runTest(t, tdt)
}

func TestBold(t *testing.T) {
	tdt := []testData{
		{
			input:      "**bold text**",
			expected:   "*bold text*\n\n",
			extensions: bf.CommonExtensions,
		},
	}

	runTest(t, tdt)
}
