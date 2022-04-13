package events

import (
  "encoding/json"
  "github.com/EntropiaTally/entropia-tally-cli/pkg/util"
  "strings"
)

type Event struct {
  Event   string      `json:"event"`
  Date    string      `json:"date"`
  Channel string      `json:"channel"`
  Values  interface{} `json:"values"`
}

// String representation of the Event
func (e *Event) String() (string, bool) {
  res, ok := e.JSON()
  if !ok {
    return "", false
  }

  return res, true
}

func (e *Event) JSON() (string, bool) {
  str, err := json.Marshal(e)
  if err != nil {
    return "", false
  }

  return string(str), true
}

// Parse an input and create an Event instance if a match is found.
func Parse(input, playerName string) (*Event, bool) {
  // Go through all available events
  for key, reg := range regexps {
    // Check for match against input
    matches := reg.FindStringSubmatch(input)

    // Skip if no match
    if len(matches) == 0 {
      continue
    }

    // Store matched raw data in map
    data := make(map[string]string)

    for i, name := range reg.SubexpNames() {
      // Skip first and stay within limits
      if i == 0 || i > len(matches) || name == "" {
        continue
      }

      data[name] = matches[i]
    }

    event, ok := NewEvent(key, playerName, data)
    if !ok {
      continue
    }

    return event, true
  }

  return nil, false
}

// NewEvent creates an Event instance based on the type of event and data, handles
// converting to the correct types for values.
func NewEvent(key, playerName string, data map[string]string) (*Event, bool) {
  event := &Event{
    Channel: strings.ToLower(data["channel"]),
    Date:    data["date"],
    Event:   getEventKey(key),
  }

  switch key {

  case "loot":
    amount := util.StringToInt(data["amount"])
    name := data["name"]

    event.Values = LootEventValues{
      Amount: amount,
      Name:   name,
      Value:  util.StringToFloatAmount(data["value"], name, amount),
    }

  case "attribute":
    event.Values = PointsGainedEventValues{
      Name:  data["name"],
      Type:  "attribute",
      Value: util.StringToFloat(data["value"]),
    }

  case "skill", "skill_alt":
    // Edge-case for certain rows
    if strings.HasPrefix(data["name"], "experience in your") {
      return nil, false
    }

    event.Values = PointsGainedEventValues{
      Name:  data["name"],
      Type:  "skill",
      Value: util.StringToFloat(data["value"]),
    }

  case "damage_inflicted", "critical_damage_inflicted":
    critical := key == "critical_damage_inflicted"
    event.Values = DamageInflictedEventValues{
      Amount:   util.StringToFloat(data["amount"]),
      Critical: critical,
    }

  case "damage_taken", "critical_damage_taken":
    critical := key == "critical_damage_taken"
    event.Values = DamageTakenEventValues{
      Amount:   util.StringToFloat(data["amount"]),
      Critical: critical,
    }

  case "enemy_dodge", "enemy_evade", "enemy_jam":
    reason := strings.Replace(key, "enemy_", "", 1)
    event.Values = EnemyEvadeEventValues{
      Reason: reason,
    }

  case "enhancer_break":
    event.Values = EnhancerBreakEventValues{
      Item:      data["item"],
      Name:      data["name"],
      Remaining: util.StringToInt(data["remaining"]),
      Value:     util.StringToFloat(data["value"]),
    }

  case "global", "hall_of_fame", "rare_loot":
    if data["player"] != playerName {
      return nil, false
    }

    enemy := getOptionalStringValue("enemy", data)
    item := getOptionalStringValue("item", data)
    location := getOptionalStringValue("location", data)
    value := util.StringToFloat(data["value"])

    // Handle values received in PEC instead of PED
    unit := getOptionalStringValue("unit", data)
    if unit != nil && *unit == "PEC" {
      value = value / 100
    }

    event.Values = SpecialLootEventValues{
      Enemy:    enemy,
      Item:     item,
      Location: location,
      Player:   data["player"],
      Type:     key,
      Value:    value,
    }

  case "heal":
    target := data["target"]
    if target == "yourself" && playerName != "" {
      target = playerName
    }

    event.Values = HealEventValues{
      Amount: util.StringToFloat(data["amount"]),
      Target: target,
    }

  case "player_deflect", "player_dodge", "player_evade":
    reason := strings.Replace(key, "player_", "", 1)
    event.Values = PlayerEvadeEventValues{
      Reason: reason,
    }

  case "position":
    alt := getOptionalIntValue("alt", data)
    name := getOptionalStringValue("name", data)

    event.Values = PositionEventValues{
      Lat:  util.StringToInt(data["lat"]),
      Lon:  util.StringToInt(data["lon"]),
      Alt:  alt,
      Name: name,
    }

  case "tier_up":
    event.Values = TierUpEventValues{
      Item: data["item"],
      Tier: util.StringToFloat(data["tier"]),
    }
  }

  return event, true
}

func getOptionalStringValue(key string, data map[string]string) *string {
  if val, ok := data[key]; ok && val != "" {
    return &val
  }

  return nil
}

func getOptionalIntValue(key string, data map[string]string) *int64 {
  if val, ok := data[key]; ok && val != "" {
    intVal := util.StringToInt(val)
    return &intVal
  }

  return nil
}
