package util

import (
  "testing"
)

func TestSliceContainsString(t *testing.T) {
  s := []string{"a", "b", "c"}

  assertEqual(t, true, sliceContainsString("a", s))
  assertEqual(t, false, sliceContainsString("d", s))
}
