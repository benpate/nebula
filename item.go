package content

import (
	"github.com/benpate/datatype"
	"github.com/benpate/path"
)

type Item struct {
	Type         string `json:"type"     bson:"type"`
	ItemID       int64  `json:"itemId"   bson:"itemId"`
	Path         string `json:"path"     bson:"path"`
	Children     []Item `json:"children" bson:"children"`
	datatype.Map `json:"map"  bson:"map"`
}

// Implement Path Interface

// GetPath implemnts the path.Getter interface
func GetPath(p path.Path) (interface{}, error) {
	return nil, nil
}

// SetPath implements the path.Setter interface
func SetPath(p path.Path, value interface{}) error {
	return nil
}

// DeletePath implements the path.Deleter interface
func DeletePath(p path.Path) error {
	return nil
}
