package maputil

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMapMerge(t *testing.T) {
	assert := require.New(t)

	out, err := Merge(nil, nil)
	assert.NoError(err)
	assert.Empty(out)

	// ------------------------------------------------------------------------
	out, err = Merge(map[string]interface{}{
		`name`: `First`,
	}, map[string]interface{}{
		`age`: 2,
	})

	assert.NoError(err)
	assert.Equal(map[string]interface{}{
		`name`: `First`,
		`age`:  2,
	}, out)

	// ------------------------------------------------------------------------
	out, err = Merge(map[string]interface{}{
		`name`:    `First`,
		`enabled`: true,
	}, nil)

	assert.NoError(err)
	assert.Equal(map[string]interface{}{
		`name`:    `First`,
		`enabled`: true,
	}, out)

	// ------------------------------------------------------------------------
	out, err = Merge(nil, map[string]interface{}{
		`name`:    `Second`,
		`enabled`: true,
	})

	assert.NoError(err)
	assert.Equal(map[string]interface{}{
		`name`:    `Second`,
		`enabled`: true,
	}, out)

	// ------------------------------------------------------------------------
	out, err = Merge(map[string]interface{}{
		`name`: `First`,
	}, map[string]interface{}{
		`name`: `Second`,
		`age`:  2,
	})

	assert.NoError(err)
	assert.Equal(map[string]interface{}{
		`name`: []interface{}{`First`, `Second`},
		`age`:  2,
	}, out)

	// ------------------------------------------------------------------------
	out, err = Merge(map[string]interface{}{
		`name`:    `First`,
		`enabled`: nil,
	}, map[string]interface{}{
		`name`:    `Second`,
		`enabled`: true,
	})

	assert.NoError(err)
	assert.Equal(map[string]interface{}{
		`name`:    []interface{}{`First`, `Second`},
		`enabled`: true,
	}, out)

	// ------------------------------------------------------------------------
	out, err = Merge(map[string]interface{}{
		`name`: `First`,
		`age`:  `yes`,
	}, map[string]interface{}{
		`name`: `Second`,
		`age`:  42,
	})

	assert.NoError(err)
	assert.Equal(map[string]interface{}{
		`name`: []interface{}{`First`, `Second`},
		`age`:  []interface{}{`yes`, 42},
	}, out)
}
