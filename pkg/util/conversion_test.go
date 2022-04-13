package util

import (
  "reflect"
  "testing"
)

func assertEqual(t *testing.T, a interface{}, b interface{}) {
  if a != b {
    t.Errorf("Received %v (type %v), expected %v (type %v)", a, reflect.TypeOf(a), b, reflect.TypeOf(b))
  }
}

func TestStringToInt(t *testing.T) {
  assertEqual(t, int64(5), StringToInt("5"))
  assertEqual(t, int64(5000), StringToInt("5000"))
  assertEqual(t, int64(0), StringToInt("invalid"))
  assertEqual(t, int64(0), StringToInt("5.0"))
}

func TestStringToFloat(t *testing.T) {
  assertEqual(t, float64(5), StringToFloat("5.0"))
  assertEqual(t, float64(5.5), StringToFloat("5.5"))
  assertEqual(t, float64(5000.9999), StringToFloat("5000.9999"))
  assertEqual(t, float64(0), StringToFloat("invalid"))
}

func TestStringToFloatAmount(t *testing.T) {
  assertEqual(t, float64(5), StringToFloatAmount("123", "Shrapnel", 50000))
  assertEqual(t, float64(123), StringToFloatAmount("123", "Other", 50000))
  assertEqual(t, float64(0), StringToFloatAmount("invalid", "Other", 50000))
}
