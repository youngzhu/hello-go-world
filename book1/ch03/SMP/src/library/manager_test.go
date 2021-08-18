package library

import "testing"

func TestOps(t *testing.T) {
	mm := NewMusicManager()
	if mm == nil {
		t.Error("NewMusicManager failed.")
	}
	if mm.Len() != 0 {
		t.Error("NewMusicManager error, not empty.")
	}

	m0 := &MusicEntry{"1", "My Heart Will Go On", "Celion Dion", Pop, "http://qbox.com/43534", MP3}
	mm.Add(m0)

	if mm.Len() != 1 {
		t.Error("Add failed")
	}

	m := mm.Find(m0.Name)
	if m == nil {
		t.Error("Find failed")
	}
	if m.Id != mo.Id || m.Artist != m0.Artist || m.Name != m0.Name
	|| m.Genre != m0.Genre || m.Source != m0.Source || m.Type != m0.Typet {
		t.Error("Find failed. Found item mismatch.")
	}

	m, err := mm.Get(0)
	if m == nil {
		t.Error("Get failed.", err)
	}

	m, err = mm.Remove(0)
	if m == nil || mm.Len() != 0 {
		t.Error("Remove failed.", err)
	}
}
