package transmission

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mergeIDsAndSHAsJustIDs(t *testing.T) {
	ids := mergeIDsAndSHAs(WhichTorrents{IDs: []int{1, 2}})
	assert.Equal(t, ids, []interface{}{1, 2})
}

func Test_mergeIDsAndSHAsJustSHAs(t *testing.T) {
	ids := mergeIDsAndSHAs(WhichTorrents{SHAs: []string{"a"}})
	assert.Equal(t, ids, []interface{}{"a"})
}

func Test_mergeIDsAndSHAs(t *testing.T) {
	ids := mergeIDsAndSHAs(WhichTorrents{IDs: []int{1}, SHAs: []string{"a", "b"}})
	assert.Equal(t, ids, []interface{}{1, "a", "b"})
}
