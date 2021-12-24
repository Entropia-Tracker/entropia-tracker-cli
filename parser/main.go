package parser

import (
  "fmt"
  "github.com/EntropiaTally/entropia-tally-cli/internal/misc"
  "strconv"
  "strings"
)

func preprocess(event *misc.Event) {
  // Handle damage inflicted
  if event.Event == "damage_inflicted" {
    (*event.Values)["critical"] = "0"
  } else if event.Event == "critical_damage_inflicted" {
    event.Event = "damage_inflicted"
    (*event.Values)["critical"] = "1"
  }

  // Handle damage taken
  if event.Event == "damage_taken" {
    (*event.Values)["critical"] = "0"
  } else if event.Event == "critical_damage_taken" {
    event.Event = "damage_taken"
    (*event.Values)["critical"] = "1"
  }

  // Handle skill
  if event.Event == "skill_alt" {
    event.Event = "skill"
  }
}

func postprocess(event *misc.Event) {
  // Calculate real shrapnel PED value
  if event.Event == "loot" && (*event.Values)["name"] == "Shrapnel" {
    amount, err := strconv.Atoi((*event.Values)["amount"])
    if err != nil {
      return
    }

    value := float64(amount) / 10000
    (*event.Values)["value"] = fmt.Sprintf("%.4f", value)
  }

  // Handle PEC values in rare loot events
  if event.Event == "rare_loot" {
    if (*event.Values)["unit"] == "PEC" {
      pecValue, err := strconv.Atoi((*event.Values)["value"])
      if err != nil {
        return
      }

      pedValue := float64(pecValue) / 100
      (*event.Values)["value"] = fmt.Sprintf("%.2f", pedValue)
    }

    delete(*event.Values, "unit")
  }
}

func validEvent(event *misc.Event, playerName string) bool {
  if event.Event == "global" || event.Event == "hall_of_fame" || event.Event == "rare_loot" {
    if (*event.Values)["player"] != playerName {
      return false
    }
  }

  if event.Event == "skill" && strings.HasPrefix((*event.Values)["name"], "experience in your") {
    return false
  }

  return true
}

// Parse a message and return the extracted data
func Parse(input, playerName string) (*misc.Event, bool) {
  for event, reg := range regexps {
    matches := reg.FindStringSubmatch(input)

    if len(matches) == 0 {
      continue
    }

    result := misc.NewEvent(event)
    preprocess(result)

    for i, name := range reg.SubexpNames() {
      // Skip first and stay within limits
      if i == 0 || i > len(matches) || name == "" {
        continue
      }

      switch {
      case name == "date":
        result.Date = matches[i]
      case name == "channel":
        result.Channel = strings.ToLower(matches[i])

      // Replace "yourself" target with playername if set on heals
      case event == "heal" && name == "target" && matches[i] == "yourself" && playerName != "":
        (*result.Values)[name] = playerName
      case true:
        (*result.Values)[name] = matches[i]
      }
    }

    postprocess(result)

    if validEvent(result, playerName) == false {
      continue
    }

    return result, true
  }

  return nil, false
}
