package util

import (
  "strconv"
)

// Items that should have their value calculated based on the amount instead of
// reported value.
var specialValueItems = []string{"Shrapnel", "Explosive Projectile"}

// StringToInt converts a string to int64, will return 0 if
// any error occurs.
func StringToInt(input string) int64 {
  result, err := strconv.ParseInt(input, 10, 64)
  if err != nil {
    return 0
  }

  return result
}

// StringToFloat converts a string to float64, will return 0.0 if
// any error occurs.
func StringToFloat(input string) float64 {
  result, err := strconv.ParseFloat(input, 64)
  if err != nil {
    return 0.0
  }

  return result
}

// StringToFloatAmount converts a string to float64, but with the added functionality
// of handling a few items where the reported values are off in the logs compared to
// their real PED value, like Shrapnel. Calculate the value based on the amount instead.
func StringToFloatAmount(input, item string, amount int64) float64 {
  if !sliceContainsString(item, specialValueItems) {
    return StringToFloat(input)
  }

  value := float64(amount) / 10_000
  return value
}
