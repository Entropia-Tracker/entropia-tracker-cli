package events

import (
  "reflect"
  "testing"
)

var testPlayerName = "Test Player Name"

func mustParseAndCompare(t *testing.T, want *Event, row, playerName string) {
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
  want := &Event{
    Event:   "loot",
    Date:    "2020-12-24 18:35:50",
    Channel: "system",
    Values: LootEventValues{
      Amount: 62885,
      Name:   "Shrapnel",
      Value:  6.2885,
    },
  }

  mustParseAndCompare(t, want, "2020-12-24 18:35:50 [System] [] You received Shrapnel x (62885) Value: 6.28 PED", testPlayerName)

  want = &Event{
    Event:   "loot",
    Date:    "2020-12-24 18:35:50",
    Channel: "system",
    Values: LootEventValues{
      Amount: 62885,
      Name:   "Explosive Projectile",
      Value:  6.2885,
    },
  }

  mustParseAndCompare(t, want, "2020-12-24 18:35:50 [System] [] You received Explosive Projectile x (62885) Value: 6.28 PED", testPlayerName)

  want = &Event{
    Event:   "loot",
    Date:    "2020-12-24 18:35:50",
    Channel: "system",
    Values: LootEventValues{
      Amount: 1877,
      Name:   "(╯°□°)╯︵ ┻━┻)",
      Value:  0.1877,
    },
  }

  mustParseAndCompare(t, want, "2020-12-24 18:35:50 [System] [] You received (╯°□°)╯︵ ┻━┻) x (1877) Value: 0.1877 PED", testPlayerName)
}

func TestParseSkill(t *testing.T) {
  want := &Event{
    Event:   "points_gained",
    Date:    "2020-12-24 18:35:50",
    Channel: "system",
    Values: PointsGainedEventValues{
      Name:  "Wounding",
      Type:  "skill",
      Value: 6.5432,
    },
  }

  mustParseAndCompare(t, want, "2020-12-24 18:35:50 [System] [] You have gained 6.5432 experience in your Wounding skill", testPlayerName)
}

func TestParseAltSkill(t *testing.T) {
  want := &Event{
    Event:   "points_gained",
    Date:    "2020-12-24 18:35:50",
    Channel: "system",
    Values: PointsGainedEventValues{
      Name:  "Dexterity",
      Type:  "skill",
      Value: 0.0426,
    },
  }

  mustParseAndCompare(t, want, "2020-12-24 18:35:50 [System] [] You have gained 0.0426 Dexterity", testPlayerName)
}

func TestParseHeal(t *testing.T) {
  want := &Event{
    Event:   "heal",
    Date:    "2020-12-24 18:35:50",
    Channel: "system",
    Values: HealEventValues{
      Target: testPlayerName,
      Amount: 38.2,
    },
  }

  mustParseAndCompare(t, want, "2020-12-24 18:35:50 [System] [] You healed yourself 38.2 points", testPlayerName)
}

func TestParseHealAlt(t *testing.T) {
  want := &Event{
    Event:   "heal",
    Date:    "2020-12-24 18:35:50",
    Channel: "system",
    Values: HealEventValues{
      Target: "Someones name",
      Amount: 38.2,
    },
  }

  mustParseAndCompare(t, want, "2020-12-24 18:35:50 [System] [] You healed Someones name with 38.2 points", testPlayerName)
}

func TestParseAttribute(t *testing.T) {
  want := &Event{
    Event:   "points_gained",
    Date:    "2020-12-24 18:35:50",
    Channel: "system",
    Values: PointsGainedEventValues{
      Type:  "attribute",
      Name:  "Agility",
      Value: 1.2345,
    },
  }

  mustParseAndCompare(t, want, "2020-12-24 18:35:50 [System] [] Your Agility has improved by 1.2345", testPlayerName)
}

func TestParseLocation(t *testing.T) {
  alt := int64(987)
  name := "Calypso"

  want := &Event{
    Event:   "position",
    Date:    "2020-12-24 18:35:50",
    Channel: "system",
    Values: PositionEventValues{
      Lat:  12345,
      Lon:  54321,
      Alt:  &alt,
      Name: &name,
    },
  }

  mustParseAndCompare(t, want, "2020-12-24 18:35:50 [System] [] [Calypso, 12345, 54321, 987, Waypoint]", testPlayerName)
}

func TestParseDamageInflicted(t *testing.T) {
  want := &Event{
    Event:   "damage_inflicted",
    Date:    "2020-12-24 18:35:50",
    Channel: "system",
    Values: DamageInflictedEventValues{
      Amount:   5.4,
      Critical: false,
    },
  }

  mustParseAndCompare(t, want, "2020-12-24 18:35:50 [System] [] You inflicted 5.4 points of damage", testPlayerName)
}

func TestParseCriticalDamageInflicted(t *testing.T) {
  want := &Event{
    Event:   "damage_inflicted",
    Date:    "2020-12-24 18:35:50",
    Channel: "system",
    Values: DamageInflictedEventValues{
      Amount:   5.4,
      Critical: true,
    },
  }

  mustParseAndCompare(t, want, "2020-12-24 18:35:50 [System] [] Critical hit - Additional damage! You inflicted 5.4 points of damage", testPlayerName)
}

func TestParseDamageTaken(t *testing.T) {
  want := &Event{
    Event:   "damage_taken",
    Date:    "2020-12-24 18:35:50",
    Channel: "system",
    Values: DamageTakenEventValues{
      Amount:   50.9,
      Critical: false,
    },
  }

  mustParseAndCompare(t, want, "2020-12-24 18:35:50 [System] [] You took 50.9 points of damage", testPlayerName)
}

func TestParseCriticalDamageTaken(t *testing.T) {
  want := &Event{
    Event:   "damage_taken",
    Date:    "2020-12-24 18:35:50",
    Channel: "system",
    Values: DamageTakenEventValues{
      Amount:   50.9,
      Critical: true,
    },
  }

  mustParseAndCompare(t, want, "2020-12-24 18:35:50 [System] [] Critical hit - Armor penetration! You took 50.9 points of damage", testPlayerName)
}

func TestParseCriticalDamageTakenAlt(t *testing.T) {
  want := &Event{
    Event:   "damage_taken",
    Date:    "2020-12-24 18:35:50",
    Channel: "system",
    Values: DamageTakenEventValues{
      Amount:   50.9,
      Critical: true,
    },
  }

  mustParseAndCompare(t, want, "2020-12-24 18:35:50 [System] [] Critical hit - Additional damage! You took 50.9 points of damage", testPlayerName)
}

func TestParseEnemyMiss(t *testing.T) {
  want := &Event{
    Event:   "enemy_miss",
    Date:    "2020-12-24 18:35:50",
    Channel: "system",
    Values:  nil,
  }

  mustParseAndCompare(t, want, "2020-12-24 18:35:50 [System] [] The attack missed you", testPlayerName)
}

func TestParseEnemyDodge(t *testing.T) {
  want := &Event{
    Event:   "enemy_evade",
    Date:    "2020-12-24 18:35:50",
    Channel: "system",
    Values: EnemyEvadeEventValues{
      Reason: "dodge",
    },
  }

  mustParseAndCompare(t, want, "2020-12-24 18:35:50 [System] [] The target Dodged your attack", testPlayerName)
}

func TestParseEnemyJam(t *testing.T) {
  want := &Event{
    Event:   "enemy_evade",
    Date:    "2020-12-24 18:35:50",
    Channel: "system",
    Values: EnemyEvadeEventValues{
      Reason: "jam",
    },
  }

  mustParseAndCompare(t, want, "2020-12-24 18:35:50 [System] [] The target Jammed your attack", testPlayerName)
}

func TestParseEnemyEvade(t *testing.T) {
  want := &Event{
    Event:   "enemy_evade",
    Date:    "2020-12-24 18:35:50",
    Channel: "system",
    Values: EnemyEvadeEventValues{
      Reason: "evade",
    },
  }

  mustParseAndCompare(t, want, "2020-12-24 18:35:50 [System] [] The target Evaded your attack", testPlayerName)
}

func TestParsePlayerDodge(t *testing.T) {
  want := &Event{
    Event:   "player_evade",
    Date:    "2020-12-24 18:35:50",
    Channel: "system",
    Values: PlayerEvadeEventValues{
      Reason: "dodge",
    },
  }

  mustParseAndCompare(t, want, "2020-12-24 18:35:50 [System] [] You Dodged the attack", testPlayerName)
}

func TestParsePlayerEvade(t *testing.T) {
  want := &Event{
    Event:   "player_evade",
    Date:    "2020-12-24 18:35:50",
    Channel: "system",
    Values: PlayerEvadeEventValues{
      Reason: "evade",
    },
  }

  mustParseAndCompare(t, want, "2020-12-24 18:35:50 [System] [] You Evaded the attack", testPlayerName)
}

func TestParsePlayerDeflect(t *testing.T) {
  want := &Event{
    Event:   "player_evade",
    Date:    "2020-12-24 18:35:50",
    Channel: "system",
    Values: PlayerEvadeEventValues{
      Reason: "deflect",
    },
  }

  mustParseAndCompare(t, want, "2020-12-24 18:35:50 [System] [] Damage deflected!", testPlayerName)
}

func TestParsePlayerMiss(t *testing.T) {
  want := &Event{
    Event:   "player_miss",
    Date:    "2020-12-24 18:35:50",
    Channel: "system",
    Values:  nil,
  }

  mustParseAndCompare(t, want, "2020-12-24 18:35:50 [System] [] You missed", testPlayerName)
}

func TestParseGlobal(t *testing.T) {
  enemy := "Kerberos Young"

  want := &Event{
    Event:   "special_loot",
    Date:    "2020-12-24 18:35:50",
    Channel: "globals",
    Values: SpecialLootEventValues{
      Type:   "global",
      Player: "Test Player Name",
      Enemy:  &enemy,
      Value:  15,
    },
  }

  mustParseAndCompare(t, want, "2020-12-24 18:35:50 [Globals] [] Test Player Name killed a creature (Kerberos Young) with a value of 15 PED!", testPlayerName)
}

func TestParseGlobalWithLocation(t *testing.T) {
  enemy := "Kerberos Young"
  location := "DSEC-9"

  want := &Event{
    Event:   "special_loot",
    Date:    "2020-12-24 18:35:50",
    Channel: "globals",
    Values: SpecialLootEventValues{
      Type:     "global",
      Player:   "Test Player Name",
      Enemy:    &enemy,
      Value:    15,
      Location: &location,
    },
  }

  mustParseAndCompare(t, want, "2020-12-24 18:35:50 [Globals] [] Test Player Name killed a creature (Kerberos Young) with a value of 15 PED at DSEC-9!", testPlayerName)
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
  enemy := "Kerberos Young"

  want := &Event{
    Event:   "special_loot",
    Date:    "2020-12-24 18:35:50",
    Channel: "globals",
    Values: SpecialLootEventValues{
      Type:   "hall_of_fame",
      Player: "Test Player Name",
      Enemy:  &enemy,
      Value:  1500,
    },
  }

  mustParseAndCompare(t, want, "2020-12-24 18:35:50 [Globals] [] Test Player Name killed a creature (Kerberos Young) with a value of 1500 PED! A record has been added to the Hall of Fame!", testPlayerName)
}

func TestParseHallOfFameWithLocation(t *testing.T) {
  enemy := "Kerberos Young"
  location := "DSEC-9"

  want := &Event{
    Event:   "special_loot",
    Date:    "2020-12-24 18:35:50",
    Channel: "globals",
    Values: SpecialLootEventValues{
      Type:     "hall_of_fame",
      Player:   "Test Player Name",
      Enemy:    &enemy,
      Value:    1500,
      Location: &location,
    },
  }

  mustParseAndCompare(t, want, "2020-12-24 18:35:50 [Globals] [] Test Player Name killed a creature (Kerberos Young) with a value of 1500 PED at DSEC-9! A record has been added to the Hall of Fame!", testPlayerName)
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
  item := "Holy Grail"

  want := &Event{
    Event:   "special_loot",
    Date:    "2020-12-24 18:35:50",
    Channel: "globals",
    Values: SpecialLootEventValues{
      Type:   "rare_loot",
      Player: "Test Player Name",
      Item:   &item,
      Value:  5000,
    },
  }

  mustParseAndCompare(t, want, "2020-12-24 18:35:50 [Globals] [] Test Player Name has found a rare item (Holy Grail) with a value of 5000 PED! A record has been added to the Hall of Fame!", testPlayerName)
}

func TestParseRareLootWithLocation(t *testing.T) {
  item := "Holy Grail"
  location := "DSEC-9"

  want := &Event{
    Event:   "special_loot",
    Date:    "2020-12-24 18:35:50",
    Channel: "globals",
    Values: SpecialLootEventValues{
      Type:     "rare_loot",
      Player:   "Test Player Name",
      Item:     &item,
      Value:    5000,
      Location: &location,
    },
  }

  mustParseAndCompare(t, want, "2020-12-24 18:35:50 [Globals] [] Test Player Name has found a rare item (Holy Grail) with a value of 5000 PED at DSEC-9! A record has been added to the Hall of Fame!", testPlayerName)
}

func TestParseRareLootPEC(t *testing.T) {
  item := "Holy Grail"
  want := &Event{
    Event:   "special_loot",
    Date:    "2020-12-24 18:35:50",
    Channel: "globals",
    Values: SpecialLootEventValues{
      Type:   "rare_loot",
      Player: "Test Player Name",
      Item:   &item,
      Value:  0.02,
    },
  }

  mustParseAndCompare(t, want, "2020-12-24 18:35:50 [Globals] [] Test Player Name has found a rare item (Holy Grail) with a value of 2 PEC! A record has been added to the Hall of Fame!", testPlayerName)
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

func TestEnhancerBreak(t *testing.T) {
  want := &Event{
    Event:   "enhancer_break",
    Date:    "2020-12-24 18:35:50",
    Channel: "system",
    Values: EnhancerBreakEventValues{
      Name:      "Weapon Damage Enhancer 1",
      Item:      "Omegaton M83 Predator",
      Remaining: 246,
      Value:     0.8000,
    },
  }

  mustParseAndCompare(t, want, "2020-12-24 18:35:50 [System] [] Your enhancer Weapon Damage Enhancer 1 on your Omegaton M83 Predator broke. You have 246 enhancers remaining on the item. You received 0.8000 PED Shrapnel.", testPlayerName)
  mustParseAndCompare(t, want, "2020-12-24 18:35:50 [System] [] Your enhancer Weapon Damage Enhancer 1 on your Omegaton M83 Predator broke. You have 246 enhancers remaining on the item. You received 0.8000 PED Shrapnel. ", testPlayerName)
}

func TestTierUp(t *testing.T) {
  want := &Event{
    Event:   "tier_up",
    Date:    "2020-12-24 18:35:50",
    Channel: "system",
    Values: TierUpEventValues{
      Item: "Arsonistic Chip 2 (L)",
      Tier: 1.12,
    },
  }

  mustParseAndCompare(t, want, "2020-12-24 18:35:50 [System] [] Your Arsonistic Chip 2 (L) has reached tier 1.12", testPlayerName)
}
