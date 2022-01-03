package nebula

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestItemDecode(t *testing.T) {

	var item Item
	encoded := `{"type":"TEXT", "data":{"text": "Hello There"}}`

	err := json.Unmarshal([]byte(encoded), &item)

	require.Nil(t, err)
	require.Equal(t, "TEXT", item.Type)
	require.Equal(t, "Hello There", item.GetString("text"))
}
