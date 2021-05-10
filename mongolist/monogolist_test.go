package mongolist

import (
	"testing"

	"github.com/dogpakk/lib/mongoutil"
)

func TestFindQuery(t *testing.T) {
	testCases := []struct {
		name      string
		listState ListState
		expected  mongoutil.FindQuery
		expectErr bool
	}{
		{
			name:      "blank",
			listState: ListState{},
			expected:  mongoutil.FindQuery{},
			expectErr: false,
		},
	}

	for _, test := range testCases {
		_, err := test.listState.FindQuery()
		if test.expectErr && err == nil {
			t.Fatalf("Testing %s.  Expected error but didn't get one", test.name)
		}
		if !test.expectErr && err != nil {
			t.Fatalf("Testing %s.  Not Expecting error but get one: %s", test.name, err)
		}
	}
}
