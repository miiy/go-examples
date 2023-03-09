package mongo_objectid

import "testing"

func TestNewObjectID(t *testing.T) {
	t.Log(NewObjectID())
	t.Log(NewObjectID())
	t.Log(NewObjectID())
	t.Log(NewObjectID())
}
