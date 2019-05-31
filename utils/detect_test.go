package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsInteger(t *testing.T) {
	assert := require.New(t)

	assert.False(IsInteger(nil))
	assert.False(IsInteger(``))
	assert.False(IsInteger(` `))
	assert.False(IsInteger(` 0 `))
	assert.False(IsInteger(`obviously not a number`))
	assert.False(IsInteger(`1 `))
	assert.False(IsInteger(`3.14`))
	assert.False(IsInteger(3.14))

	assert.True(IsInteger(0))
	assert.True(IsInteger(1))
	assert.True(IsInteger(`0`))
	assert.True(IsInteger(`1`))
}

func TestIsFloat(t *testing.T) {
	assert := require.New(t)

	assert.False(IsFloat(nil))
	assert.False(IsFloat(``))
	assert.False(IsFloat(` `))
	assert.False(IsFloat(` 0 `))
	assert.False(IsFloat(`1 `))
	assert.False(IsFloat(`obviously not a number`))

	assert.True(IsFloat(`3.14`))
	assert.True(IsFloat(3.14))
	assert.True(IsFloat(0))
	assert.True(IsFloat(1))
	assert.True(IsFloat(`0`))
	assert.True(IsFloat(`1`))
}

func TestIsHexadecimal(t *testing.T) {
	assert := require.New(t)

	assert.False(IsHexadecimal(nil))
	assert.False(IsHexadecimal(``))
	assert.False(IsHexadecimal(` `))
	assert.False(IsHexadecimal(` 0 `))
	assert.False(IsHexadecimal(`1 `))
	assert.False(IsHexadecimal(`obviously not a number`))
	assert.False(IsHexadecimal(`3.14`))
	assert.False(IsHexadecimal(3.14))

	assert.False(IsHexadecimal(0))
	assert.False(IsHexadecimal(1))
	assert.False(IsHexadecimal(`0`))
	assert.False(IsHexadecimal(`1`))
	assert.False(IsHexadecimal(0x0))
	assert.False(IsHexadecimal(0x1))
	assert.False(IsHexadecimal(0xA))
	assert.True(IsHexadecimal(`0x0`))
	assert.True(IsHexadecimal(`0x1`))
	assert.True(IsHexadecimal(`0x2`))
	assert.True(IsHexadecimal(`0x3`))
	assert.True(IsHexadecimal(`0x4`))
	assert.True(IsHexadecimal(`0x5`))
	assert.True(IsHexadecimal(`0x6`))
	assert.True(IsHexadecimal(`0x7`))
	assert.True(IsHexadecimal(`0x8`))
	assert.True(IsHexadecimal(`0x9`))
	assert.True(IsHexadecimal(`0xA`))
	assert.True(IsHexadecimal(`0xB`))
	assert.True(IsHexadecimal(`0xC`))
	assert.True(IsHexadecimal(`0xD`))
	assert.True(IsHexadecimal(`0xE`))
	assert.True(IsHexadecimal(`0xF`))
	assert.False(IsHexadecimal(`0xG`))
	assert.True(IsHexadecimal(`0xa`))
	assert.True(IsHexadecimal(`0xb`))
	assert.True(IsHexadecimal(`0xc`))
	assert.True(IsHexadecimal(`0xd`))
	assert.True(IsHexadecimal(`0xe`))
	assert.True(IsHexadecimal(`0xf`))
	assert.False(IsHexadecimal(`0xg`))

	assert.False(IsHexadecimal(`0`))
	assert.False(IsHexadecimal(`1`))
	assert.False(IsHexadecimal(`2`))
	assert.False(IsHexadecimal(`3`))
	assert.False(IsHexadecimal(`4`))
	assert.False(IsHexadecimal(`5`))
	assert.False(IsHexadecimal(`6`))
	assert.False(IsHexadecimal(`7`))
	assert.False(IsHexadecimal(`8`))
	assert.False(IsHexadecimal(`9`))
	assert.False(IsHexadecimal(`A`))
	assert.False(IsHexadecimal(`B`))
	assert.False(IsHexadecimal(`C`))
	assert.False(IsHexadecimal(`D`))
	assert.False(IsHexadecimal(`E`))
	assert.False(IsHexadecimal(`F`))
	assert.False(IsHexadecimal(`G`))
	assert.False(IsHexadecimal(`a`))
	assert.False(IsHexadecimal(`b`))
	assert.False(IsHexadecimal(`c`))
	assert.False(IsHexadecimal(`d`))
	assert.False(IsHexadecimal(`e`))
	assert.False(IsHexadecimal(`f`))
	assert.False(IsHexadecimal(`g`))

	assert.False(IsHexadecimal(`dEaDbEeF`))
	assert.True(IsHexadecimal(`0xdEaDbEeF`))
	assert.False(IsHexadecimal(`DEADBEEF`))
	assert.True(IsHexadecimal(`0xDEADBEEF`))
	assert.False(IsHexadecimal(`deadbeef`))
	assert.True(IsHexadecimal(`0xdeadbeef`))
}
