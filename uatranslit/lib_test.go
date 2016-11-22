package uatranslit

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReplaceUARune(t *testing.T) {
	assert := assert.New(t)
	assert.Equal([]rune{'m'}, ReplaceUARune('м', true), "rune with no first position (first)")
	assert.Equal([]rune{'m'}, ReplaceUARune('м', false), "rune with no first position (other)")
	assert.Equal([]rune{'y', 'u'}, ReplaceUARune('ю', true), "rune with first position (first)")
	assert.Equal([]rune{'i', 'u'}, ReplaceUARune('ю', false), "rune with first position (other)")
	assert.Equal([]rune{'I', 'a'}, ReplaceUARune('Я', false), "rune with first position upper case")
	assert.NotEqual([]rune{'I', 'a'}, ReplaceUARune('я', false), "rune with first position lower case")
	assert.Equal([]rune{'w'}, ReplaceUARune('w', true), "test check replace non UA rune")
}

func TestReplaceUARunes(t *testing.T) {
	assert := assert.New(t)
	assert.Equal([]rune("moloko"), ReplaceUARunes([]rune("молоко")), "convert simple word")
	assert.Equal([]rune("moloko, maslo"), ReplaceUARunes([]rune("молоко, масло")), "convert two simple words")
	assert.Equal([]rune("tsia yalynka"), ReplaceUARunes([]rune("ця ялинка")), "test detecting first position after space character")
	assert.Equal([]rune("Baba-yaha"), ReplaceUARunes([]rune("Баба-яга")), "test detecting first position after non space character")
	assert.Equal([]rune("Halia, Yana"), ReplaceUARunes([]rune("Галя, Яна")), "test rune with position")
}
