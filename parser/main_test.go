package parser

import (
  "github.com/EntropiaTally/entropia-tally-cli/internal/misc"
  "reflect"
  "testing"
)

var testPlayerName = "Test Player Name"

func parseAndCompare(t *testing.T, want *misc.Event, row, playerName string) {
  event, ok := Parse(row, playerName)

  if !ok {
    t.Errorf("got %t; want %t", ok, true)
  }

  if !reflect.DeepEqual(event, want) {
    eventStr, _ := event.String()
    wantStr, _ := want.String()

    t.Errorf("got %s; want %s", eventStr, wantStr)
  }
}

func TestParseLoot(t *testing.T) {
  want := &misc.Event{
    Event:   "loot",
    Date:    "2020-12-24 18:35:50",
    Channel: "system",
    Values: &map[string]string{
      "value":  "6.2885",
      "name":   "Shrapnel",
      "amount": "62885",
    },
  }

  parseAndCompare(t, want, "2020-12-24 18:35:50 [System] [] You received Shrapnel x (62885) Value: 6.28 PED", testPlayerName)

  want = &misc.Event{
    Event:   "loot",
    Date:    "2020-12-24 18:35:50",
    Channel: "system",
    Values: &map[string]string{
      "value":  "0.1877",
      "name":   "(╯°□°)╯︵ ┻━┻)",
      "amount": "1877",
    },
  }

  parseAndCompare(t, want, "2020-12-24 18:35:50 [System] [] You received (╯°□°)╯︵ ┻━┻) x (1877) Value: 0.1877 PED", testPlayerName)
}

func TestParseSkill(t *testing.T) {
  want := &misc.Event{
    Event:   "skill",
    Date:    "2020-12-24 18:35:50",
    Channel: "system",
    Values: &map[string]string{
      "name":  "Wounding",
      "value": "6.5432",
    },
  }

  parseAndCompare(t, want, "2020-12-24 18:35:50 [System] [] You have gained 6.5432 experience in your Wounding skill", testPlayerName)
}

func TestParseAltSkill(t *testing.T) {
  want := &misc.Event{
    Event:   "skill",
    Date:    "2020-12-24 18:35:50",
    Channel: "system",
    Values: &map[string]string{
      "name":  "Dexterity",
      "value": "0.0426",
    },
  }

  parseAndCompare(t, want, "2020-12-24 18:35:50 [System] [] You have gained 0.0426 Dexterity", testPlayerName)
}

func TestParseHeal(t *testing.T) {
  want := &misc.Event{
    Event:   "heal",
    Date:    "2020-12-24 18:35:50",
    Channel: "system",
    Values: &map[string]string{
      "target": testPlayerName,
      "amount": "38.2",
    },
  }

  parseAndCompare(t, want, "2020-12-24 18:35:50 [System] [] You healed yourself 38.2 points", testPlayerName)
}

func TestParseHealAlt(t *testing.T) {
  want := &misc.Event{
    Event:   "heal",
    Date:    "2020-12-24 18:35:50",
    Channel: "system",
    Values: &map[string]string{
      "target": "Someones name",
      "amount": "38.2",
    },
  }

  parseAndCompare(t, want, "2020-12-24 18:35:50 [System] [] You healed Someones name with 38.2 points", testPlayerName)
}

func TestParseAttribute(t *testing.T) {
  want := &misc.Event{
    Event:   "attribute",
    Date:    "2020-12-24 18:35:50",
    Channel: "system",
    Values: &map[string]string{
      "name":  "Agility",
      "value": "1.2345",
    },
  }

  parseAndCompare(t, want, "2020-12-24 18:35:50 [System] [] Your Agility has improved by 1.2345", testPlayerName)
}

func TestParseLocation(t *testing.T) {
  want := &misc.Event{
    Event:   "position",
    Date:    "2020-12-24 18:35:50",
    Channel: "system",
    Values: &map[string]string{
      "lat":  "12345",
      "lon":  "54321",
      "alt":  "987",
      "name": "Calypso",
    },
  }

  parseAndCompare(t, want, "2020-12-24 18:35:50 [System] [] [Calypso, 12345, 54321, 987, Waypoint]", testPlayerName)
}

func TestParseDamageInflicted(t *testing.T) {
  want := &misc.Event{
    Event:   "damage_inflicted",
    Date:    "2020-12-24 18:35:50",
    Channel: "system",
    Values: &map[string]string{
      "amount":   "5.4",
      "critical": "0",
    },
  }

  parseAndCompare(t, want, "2020-12-24 18:35:50 [System] [] You inflicted 5.4 points of damage", testPlayerName)
}

func TestParseCriticalDamageInflicted(t *testing.T) {
  want := &misc.Event{
    Event:   "damage_inflicted",
    Date:    "2020-12-24 18:35:50",
    Channel: "system",
    Values: &map[string]string{
      "amount":   "5.4",
      "critical": "1",
    },
  }

  parseAndCompare(t, want, "2020-12-24 18:35:50 [System] [] Critical hit - Additional damage! You inflicted 5.4 points of damage", testPlayerName)
}

func TestParseDamageTaken(t *testing.T) {
  want := &misc.Event{
    Event:   "damage_taken",
    Date:    "2020-12-24 18:35:50",
    Channel: "system",
    Values: &map[string]string{
      "amount":   "50.9",
      "critical": "0",
    },
  }

  parseAndCompare(t, want, "2020-12-24 18:35:50 [System] [] You took 50.9 points of damage", testPlayerName)
}

func TestParseCriticalDamageTaken(t *testing.T) {
  want := &misc.Event{
    Event:   "damage_taken",
    Date:    "2020-12-24 18:35:50",
    Channel: "system",
    Values: &map[string]string{
      "amount":   "50.9",
      "critical": "1",
    },
  }

  parseAndCompare(t, want, "2020-12-24 18:35:50 [System] [] Critical hit - Armor penetration! You took 50.9 points of damage", testPlayerName)
}

func TestParseCriticalDamageTakenAlt(t *testing.T) {
  want := &misc.Event{
    Event:   "damage_taken",
    Date:    "2020-12-24 18:35:50",
    Channel: "system",
    Values: &map[string]string{
      "amount":   "50.9",
      "critical": "1",
    },
  }

  parseAndCompare(t, want, "2020-12-24 18:35:50 [System] [] Critical hit - Additional damage! You took 50.9 points of damage", testPlayerName)
}

func TestParseEnemyMiss(t *testing.T) {
  want := &misc.Event{
    Event:   "enemy_miss",
    Date:    "2020-12-24 18:35:50",
    Channel: "system",
    Values:  &map[string]string{},
  }

  parseAndCompare(t, want, "2020-12-24 18:35:50 [System] [] The attack missed you", testPlayerName)
}

func TestParseEnemyDodge(t *testing.T) {
  want := &misc.Event{
    Event:   "enemy_dodge",
    Date:    "2020-12-24 18:35:50",
    Channel: "system",
    Values:  &map[string]string{},
  }

  parseAndCompare(t, want, "2020-12-24 18:35:50 [System] [] The target Dodged your attack", testPlayerName)
}

func TestParseEnemyEvade(t *testing.T) {
  want := &misc.Event{
    Event:   "enemy_evade",
    Date:    "2020-12-24 18:35:50",
    Channel: "system",
    Values:  &map[string]string{},
  }

  parseAndCompare(t, want, "2020-12-24 18:35:50 [System] [] The target Evaded your attack", testPlayerName)
}

func TestParsePlayerDodge(t *testing.T) {
  want := &misc.Event{
    Event:   "player_dodge",
    Date:    "2020-12-24 18:35:50",
    Channel: "system",
    Values:  &map[string]string{},
  }

  parseAndCompare(t, want, "2020-12-24 18:35:50 [System] [] You Dodged the attack", testPlayerName)
}

func TestParsePlayerEvade(t *testing.T) {
  want := &misc.Event{
    Event:   "player_evade",
    Date:    "2020-12-24 18:35:50",
    Channel: "system",
    Values:  &map[string]string{},
  }

  parseAndCompare(t, want, "2020-12-24 18:35:50 [System] [] You Evaded the attack", testPlayerName)
}

func TestParsePlayerDeflect(t *testing.T) {
  want := &misc.Event{
    Event:   "player_deflect",
    Date:    "2020-12-24 18:35:50",
    Channel: "system",
    Values:  &map[string]string{},
  }

  parseAndCompare(t, want, "2020-12-24 18:35:50 [System] [] Damage deflected!", testPlayerName)
}

func TestParseGlobal(t *testing.T) {
  want := &misc.Event{
    Event:   "global",
    Date:    "2020-12-24 18:35:50",
    Channel: "globals",
    Values: &map[string]string{
      "player": "Test Player Name",
      "enemy":  "Kerberos Young",
      "value":  "15",
    },
  }

  parseAndCompare(t, want, "2020-12-24 18:35:50 [Globals] [] Test Player Name killed a creature (Kerberos Young) with a value of 15 PED!", testPlayerName)
}

func TestParseGlobalWrongName(t *testing.T) {
  result, ok := Parse("2020-12-24 18:35:50 [Globals] [] Test Player Name killed a creature (Kerberos Young) with a value of 15 PED!", "Other Player Name")

  if ok != false {
    t.Errorf("got %t; want %t", ok, false)
  }

  if result != nil {
    wantStr, _ := result.String()
    t.Errorf("got %s; want nil", wantStr)
  }
}

func TestParseHallOfFame(t *testing.T) {
  want := &misc.Event{
    Event:   "hall_of_fame",
    Date:    "2020-12-24 18:35:50",
    Channel: "globals",
    Values: &map[string]string{
      "player": "Test Player Name",
      "enemy":  "Kerberos Young",
      "value":  "1500",
    },
  }

  parseAndCompare(t, want, "2020-12-24 18:35:50 [Globals] [] Test Player Name killed a creature (Kerberos Young) with a value of 1500 PED! A record has been added to the Hall of Fame!", testPlayerName)
}

func TestParseHallOfFameWrongName(t *testing.T) {
  result, ok := Parse("2020-12-24 18:35:50 [Globals] [] Test Player Name killed a creature (Kerberos Young) with a value of 1500 PED! A record has been added to the Hall of Fame!", "Other Player Name")

  if ok != false {
    t.Errorf("got %t; want %t", ok, false)
  }

  if result != nil {
    wantStr, _ := result.String()
    t.Errorf("got %s; want nil", wantStr)
  }
}

func TestParseRareLoot(t *testing.T) {
  want := &misc.Event{
    Event:   "rare_loot",
    Date:    "2020-12-24 18:35:50",
    Channel: "globals",
    Values: &map[string]string{
      "player": "Test Player Name",
      "item":   "Holy Grail",
      "value":  "5000",
    },
  }

  parseAndCompare(t, want, "2020-12-24 18:35:50 [Globals] [] Test Player Name has found a rare item (Holy Grail) with a value of 5000 PED! A record has been added to the Hall of Fame!", testPlayerName)
}

func TestParseRareLootWrongName(t *testing.T) {
  result, ok := Parse("2020-12-24 18:35:50 [Globals] [] Test Player Name has found a rare item (Holy Grail) with a value of 5000 PED! A record has been added to the Hall of Fame!", "Other Player Name")

  if ok != false {
    t.Errorf("got %t; want %t", ok, false)
  }

  if result != nil {
    wantStr, _ := result.String()
    t.Errorf("got %s; want nil", wantStr)
  }
}
