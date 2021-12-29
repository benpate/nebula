package transaction

import (
	"testing"

	"github.com/benpate/datatype"
	"github.com/benpate/nebula"
	"github.com/stretchr/testify/require"
)

func TestAddItem_InsertContainer_Above(t *testing.T) {

	c := nebula.Container{
		{
			Type:  nebula.ItemTypeLayout,
			Check: "123",
			Refs:  []int{1},
			Data: datatype.Map{
				"style": nebula.LayoutStyleRows,
			},
		},
		{
			Type:  nebula.ItemTypeWYSIWYG,
			Check: "123",
			Data: datatype.Map{
				"html": "This is the first item",
			},
		},
	}
	txn := NewItem{
		ItemID:   1,
		Place:    nebula.LayoutPlaceAbove,
		ItemType: nebula.ItemTypeHTML,
		Check:    "123",
	}

	id, err := txn.Execute(&c)

	require.Nil(t, err)
	require.Equal(t, 3, c.Len())
	require.Equal(t, 0, id)
	require.Equal(t, nebula.ItemTypeLayout, c[0].Type)
	require.Equal(t, nebula.LayoutStyleRows, c[0].Data["style"])
	require.Equal(t, []int{2, 1}, c[0].Refs)

	require.Equal(t, nebula.ItemTypeWYSIWYG, c[1].Type)
	require.Equal(t, "This is the first item", c[1].Data["html"])

	require.Equal(t, nebula.ItemTypeHTML, c[2].Type)
	require.Empty(t, c[2].Data["html"])
}

func TestAddItem_InsertContainer_Above2(t *testing.T) {

	c := nebula.Container{
		{
			Type:  nebula.ItemTypeLayout,
			Check: "123",
			Refs:  []int{1, 2},
			Data: datatype.Map{
				"style": nebula.LayoutStyleRows,
			},
		},
		{
			Type:  nebula.ItemTypeWYSIWYG,
			Check: "123",
			Data: datatype.Map{
				"html": "This is the first item",
			},
		},
		{
			Type:  nebula.ItemTypeWYSIWYG,
			Check: "123",
			Data: datatype.Map{
				"html": "This is the second item",
			},
		},
	}
	txn := NewItem{
		ItemID:   2,
		Place:    nebula.LayoutPlaceAbove,
		ItemType: nebula.ItemTypeHTML,
		Check:    "123",
	}

	id, err := txn.Execute(&c)

	require.Nil(t, err)
	require.Equal(t, 4, c.Len())
	require.Equal(t, 0, id)
	require.Equal(t, nebula.ItemTypeLayout, c[0].Type)
	require.Equal(t, nebula.LayoutStyleRows, c[0].Data["style"])
	require.Equal(t, []int{1, 3, 2}, c[0].Refs)

	require.Equal(t, nebula.ItemTypeWYSIWYG, c[1].Type)
	require.Equal(t, "This is the first item", c[1].Data["html"])

	require.Equal(t, nebula.ItemTypeWYSIWYG, c[2].Type)
	require.Equal(t, "This is the second item", c[2].Data["html"])

	require.Equal(t, nebula.ItemTypeHTML, c[3].Type)
	require.Empty(t, c[3].Data["html"])
}

func TestAddItem_InsertContainer_Below(t *testing.T) {

	c := nebula.Container{
		{
			Type:  nebula.ItemTypeLayout,
			Check: "123",
			Refs:  []int{1},
			Data: datatype.Map{
				"style": nebula.LayoutStyleRows,
			},
		},
		{
			Type:  nebula.ItemTypeWYSIWYG,
			Check: "123",
			Data: datatype.Map{
				"html": "This is the first item",
			},
		},
	}
	txn := NewItem{
		ItemID:   1,
		Place:    nebula.LayoutPlaceBelow,
		ItemType: nebula.ItemTypeHTML,
		Check:    "123",
	}

	id, err := txn.Execute(&c)

	require.Nil(t, err)
	require.Equal(t, 3, c.Len())
	require.Equal(t, 0, id)
	require.Equal(t, nebula.ItemTypeLayout, c[0].Type)
	require.Equal(t, nebula.LayoutStyleRows, c[0].Data["style"])
	require.Equal(t, []int{1, 2}, c[0].Refs)

	require.Equal(t, nebula.ItemTypeWYSIWYG, c[1].Type)
	require.Equal(t, "This is the first item", c[1].Data["html"])

	require.Equal(t, nebula.ItemTypeHTML, c[2].Type)
	require.Empty(t, c[2].Data["html"])
}

func TestAddItem_InsertContainer_Left(t *testing.T) {

	c := nebula.Container{
		{
			Type:  nebula.ItemTypeLayout,
			Check: "123",
			Refs:  []int{1},
			Data: datatype.Map{
				"style": nebula.LayoutStyleColumns,
			},
		},
		{
			Type:  nebula.ItemTypeWYSIWYG,
			Check: "123",
			Data: datatype.Map{
				"html": "This is the first item",
			},
		},
	}
	txn := NewItem{
		ItemID:   1,
		Place:    nebula.LayoutPlaceLeft,
		ItemType: nebula.ItemTypeHTML,
		Check:    "123",
	}

	id, err := txn.Execute(&c)

	require.Nil(t, err)
	require.Equal(t, 3, c.Len())
	require.Equal(t, 0, id)
	require.Equal(t, nebula.ItemTypeLayout, c[0].Type)
	require.Equal(t, nebula.LayoutStyleColumns, c[0].Data["style"])
	require.Equal(t, []int{2, 1}, c[0].Refs)

	require.Equal(t, nebula.ItemTypeWYSIWYG, c[1].Type)
	require.Equal(t, "This is the first item", c[1].Data["html"])

	require.Equal(t, nebula.ItemTypeHTML, c[2].Type)
	require.Empty(t, c[2].Data["html"])
}

func TestAddItem_InsertContainer_Right(t *testing.T) {

	c := nebula.Container{
		{
			Type:  nebula.ItemTypeLayout,
			Check: "123",
			Refs:  []int{1},
			Data: datatype.Map{
				"style": nebula.LayoutStyleColumns,
			},
		},
		{
			Type:  nebula.ItemTypeWYSIWYG,
			Check: "123",
			Data: datatype.Map{
				"html": "This is the first item",
			},
		},
	}
	txn := NewItem{
		ItemID:   1,
		Place:    nebula.LayoutPlaceRight,
		ItemType: nebula.ItemTypeHTML,
		Check:    "123",
	}

	id, err := txn.Execute(&c)

	require.Nil(t, err)
	require.Equal(t, 3, c.Len())
	require.Equal(t, 0, id)
	require.Equal(t, nebula.ItemTypeLayout, c[0].Type)
	require.Equal(t, nebula.LayoutStyleColumns, c[0].Data["style"])
	require.Equal(t, []int{1, 2}, c[0].Refs)

	require.Equal(t, nebula.ItemTypeWYSIWYG, c[1].Type)
	require.Equal(t, "This is the first item", c[1].Data["html"])

	require.Equal(t, nebula.ItemTypeHTML, c[2].Type)
	require.Empty(t, c[2].Data["html"])
}
