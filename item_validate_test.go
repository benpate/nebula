package nebula

import (
	"testing"

	"github.com/benpate/derp"
	"github.com/stretchr/testify/require"
)

func TestItemValidate_Empty(t *testing.T) {

	item := Item{}

	require.NotNil(t, item.Validate(""))
	require.Equal(t, "Item.Validate: Item is empty", derp.Message(item.Validate("")))
}

func TestItemValidate_Checksum(t *testing.T) {

	item := Item{
		Type:  "HTML",
		Check: "123456789101112",
	}

	require.Nil(t, item.Validate("123456789101112"))
}

func TestItemValidate_BadChecksum(t *testing.T) {

	item := Item{
		Type:  "HTML",
		Check: "123456789101112",
	}

	require.NotNil(t, item.Validate(""))
	require.NotNil(t, item.Validate("123456789"))
}
