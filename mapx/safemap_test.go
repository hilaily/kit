package mapx

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSafeMapRace(t *testing.T) {
	m := NewSafeMap[string, string]()
	m.Set("user1", "ace")
	go func() {
		m.Set("user2", "bob")
	}()
	m.Get("user1")
}

func TestSafeMap(t *testing.T) {
	Convey("TestSafeMap", t, func() {
		m := NewSafeMap[string, string]()
		m.Set("user1", "ace")

		v, ok := m.Get("user1")
		So(ok, ShouldBeTrue)
		So(v, ShouldEqual, "ace")

		m.Set("user2", "bob")
		l := m.Len()
		So(l, ShouldEqual, 2)
	})
}
