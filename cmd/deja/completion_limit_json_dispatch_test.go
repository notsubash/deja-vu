package main

import "testing"

func TestLimitJSONCompletionDispatch(t *testing.T) {
	checkCompletionDispatch(t, []completionDispatchCase{
		{[]string{"deja", "fix", "--"}, []string{"--limit", "--json"}, []string{"--since", "--re", "--settle", "--unsettle"}},
		{[]string{"deja", "friction", "--"}, []string{"--limit", "--json", "--settle", "--unsettle"}, []string{"--since", "--re"}},
		{[]string{"deja", "tests", "--"}, []string{"--limit", "--json"}, []string{"--since", "--re", "--settle", "--unsettle"}},
	})
}
