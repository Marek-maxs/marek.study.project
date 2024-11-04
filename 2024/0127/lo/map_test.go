package lo

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
*
* Author: Marek
* Date: 2024-01-27 16:43
* Email: 364021318@qq.com
*
 */

func TestKeys(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	r1 := Keys(map[string]int{"foo": 1, "bar": 2})
	sort.Strings(r1)

	is.Equal(r1, []string{"bar", "foo"})
}
