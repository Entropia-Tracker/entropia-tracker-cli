package parser

import (
  "regexp"
)

// 0000-00-00 00:00:00 [System] [] You received Shrapnel x (1000) Value: 0.0000 PED
// date, channel, item, amount, value
var lootRegexp = regexp.MustCompile(`^(?P<date>\d{4}-\d{2}-\d{2}\s\d{2}:\d{2}:\d{2})\s\[(?P<channel>\S+)\]\s\[\]\sYou\sreceived\s(?P<name>.*)\sx\s\((?P<amount>\d+)\)\sValue:\s(?P<value>[\d.]+)\sPED$`)

// 0000-00-00 00:00:00 [System] [] You have gained 0.0000 experience in your Wounding skill
// date, channel, name, value
var skillRegexp = regexp.MustCompile(`^(?P<date>\d{4}-\d{2}-\d{2}\s\d{2}:\d{2}:\d{2})\s\[(?P<channel>\S+)\]\s\[\]\sYou\shave\sgained\s(?P<value>[\d.]+)\sexperience\sin\syour\s(?P<name>.*)\sskill$`)

// 0000-00-00 00:00:00 [System] [] You have gained 0.0426 Dexterity
// date, channel, name, value
var skillAltRegexp = regexp.MustCompile(`^(?P<date>\d{4}-\d{2}-\d{2}\s\d{2}:\d{2}:\d{2})\s\[(?P<channel>\S+)\]\s\[\]\sYou\shave\sgained\s(?P<value>[\d.]+)\s(?P<name>.*)$`)

// 0000-00-00 00:00:00 [System] [] Your Agility has improved by 0.0001
// date, channel, name, value
var attributeRegexp = regexp.MustCompile(`^(?P<date>\d{4}-\d{2}-\d{2}\s\d{2}:\d{2}:\d{2})\s\[(?P<channel>\S+)\]\s\[\]\sYour\s(?P<name>.*)\shas\simproved\sby\s(?P<value>[\d.]+)$`)

// 0000-00-00 00:00:00 [System] [] [Calypso, 00000, 00000, 000, Waypoint]
// date, channel, name, lat, lon, alt
var positionRegexp = regexp.MustCompile(`^(?P<date>\d{4}-\d{2}-\d{2}\s\d{2}:\d{2}:\d{2})\s\[(?P<channel>\S+)\]\s\[\]\s\[(?P<name>.*),\s(?P<lat>\d+),\s(?P<lon>\d+),\s(?P<alt>\d+),\sWaypoint\]$`)

// 0000-00-00 00:00:00 [System] [] You inflicted 10.0 points of damage
// date, channel, amount
var damageInflictedRegexp = regexp.MustCompile(`^(?P<date>\d{4}-\d{2}-\d{2}\s\d{2}:\d{2}:\d{2})\s\[(?P<channel>\S+)\]\s\[\]\sYou\sinflicted\s(?P<amount>[\d.]+)\spoints\sof\sdamage$`)

// 0000-00-00 00:00:00 [System] [] Critical hit - Additional damage! You inflicted 10.0 points of damage
// date, channel, amount
var criticalDamageInflictedRegexp = regexp.MustCompile(`^(?P<date>\d{4}-\d{2}-\d{2}\s\d{2}:\d{2}:\d{2})\s\[(?P<channel>\S+)\]\s\[\]\sCritical\shit\s-\sAdditional\sdamage!\sYou\sinflicted\s(?P<amount>[\d.]+)\spoints\sof\sdamage$`)

// 0000-00-00 00:00:00 [System] []  You took 50.9 points of damage
// date, channel, amount
var damageTakenRegexp = regexp.MustCompile(`^(?P<date>\d{4}-\d{2}-\d{2}\s\d{2}:\d{2}:\d{2})\s\[(?P<channel>\S+)\]\s\[\]\sYou\stook\s(?P<amount>[\d.]+)\spoints\sof\sdamage$`)

// 0000-00-00 00:00:00 [System] [] Critical hit - Armor penetration! You took 50.9 points of damage
// 0000-00-00 00:00:00 [System] [] Critical hit - Additional damage! You took 800.5 points of damage
// date, channel, amount
var criticalDamageTakenRegexp = regexp.MustCompile(`^(?P<date>\d{4}-\d{2}-\d{2}\s\d{2}:\d{2}:\d{2})\s\[(?P<channel>\S+)\]\s\[\]\sCritical\shit\s-\s(Armor\spenetration|Additional\sdamage)!\sYou\stook\s(?P<amount>[\d.]+)\spoints\sof\sdamage$`)

// 0000-00-00 00:00:00 [System] [] The attack missed you
// date, channel
var enemyMissRegexp = regexp.MustCompile(`^(?P<date>\d{4}-\d{2}-\d{2}\s\d{2}:\d{2}:\d{2})\s\[(?P<channel>\S+)\]\s\[\]\sThe\sattack\smissed\syou$`)

// 0000-00-00 00:00:00 [System] [] The target Dodged your attack
// date, channel
var enemyDodgeRegexp = regexp.MustCompile(`^(?P<date>\d{4}-\d{2}-\d{2}\s\d{2}:\d{2}:\d{2})\s\[(?P<channel>\S+)\]\s\[\]\sThe\starget\sDodged\syour\sattack$`)

// 0000-00-00 00:00:00 [System] [] The target Evaded your attack
// date, channel
var enemyEvadeRegexp = regexp.MustCompile(`^(?P<date>\d{4}-\d{2}-\d{2}\s\d{2}:\d{2}:\d{2})\s\[(?P<channel>\S+)\]\s\[\]\sThe\starget\sEvaded\syour\sattack$`)

// 0000-00-00 00:00:00 [System] [] The target Jammed your attack
// date, channel
var enemyJamRegexp = regexp.MustCompile(`^(?P<date>\d{4}-\d{2}-\d{2}\s\d{2}:\d{2}:\d{2})\s\[(?P<channel>\S+)\]\s\[\]\sThe\starget\sJammed\syour\sattack$`)

// 0000-00-00 00:00:00 [System] [] You Dodged the attack
// date, channel
var playerDodgeRegexp = regexp.MustCompile(`^(?P<date>\d{4}-\d{2}-\d{2}\s\d{2}:\d{2}:\d{2})\s\[(?P<channel>\S+)\]\s\[\]\sYou\sDodged\sthe\sattack$`)

// 0000-00-00 00:00:00 [System] [] You Evaded the attack
// date, channel
var playerEvadeRegexp = regexp.MustCompile(`^(?P<date>\d{4}-\d{2}-\d{2}\s\d{2}:\d{2}:\d{2})\s\[(?P<channel>\S+)\]\s\[\]\sYou\sEvaded\sthe\sattack$`)

// 0000-00-00 00:00:00 [System] [] Damage deflected!
// date, channel
var playerDeflectRegexp = regexp.MustCompile(`^(?P<date>\d{4}-\d{2}-\d{2}\s\d{2}:\d{2}:\d{2})\s\[(?P<channel>\S+)\]\s\[\]\sDamage\sdeflected!$`)

// 0000-00-00 00:00:00 [System] [] You missed
// date, channel
var playerMissRegexp = regexp.MustCompile(`^(?P<date>\d{4}-\d{2}-\d{2}\s\d{2}:\d{2}:\d{2})\s\[(?P<channel>\S+)\]\s\[\]\sYou\smissed$`)

// 0000-00-00 00:00:00 [Globals] [] Example Player Name killed a creature (Kerberos Young) with a value of 15 PED!
// date, channel, player, enemy, value
var globalRegexp = regexp.MustCompile(`^(?P<date>\d{4}-\d{2}-\d{2}\s\d{2}:\d{2}:\d{2})\s\[(?P<channel>\S+)\]\s\[\]\s(?P<player>.*)\skilled\sa\screature\s\((?P<enemy>.*?)\)\swith\sa\svalue\sof\s(?P<value>\d+)\sPED!$`)

// 0000-00-00 00:00:00 [Globals] [] Example Player Name killed a creature (Kerberos Young) with a value of 15 PED! A record has been added to the Hall of Fame!
// date, channel, player, enemy, value
var hallOfFameRegexp = regexp.MustCompile(`^(?P<date>\d{4}-\d{2}-\d{2}\s\d{2}:\d{2}:\d{2})\s\[(?P<channel>\S+)\]\s\[\]\s(?P<player>.*)\skilled\sa\screature\s\((?P<enemy>.*?)\)\swith\sa\svalue\sof\s(?P<value>\d+)\sPED!\sA\srecord\shas\sbeen\sadded\sto\sthe\sHall\sof\sFame!$`)

// 0000-00-00 00:00:00 [Globals] [] Example Player Name has found a rare item (Holy Grail) with a value of 5000 PED! A record has been added to the Hall of Fame!
// date, channel, player, item, value
var rareLootRegexp = regexp.MustCompile(`^(?P<date>\d{4}-\d{2}-\d{2}\s\d{2}:\d{2}:\d{2})\s\[(?P<channel>\S+)\]\s\[\]\s(?P<player>.*)\shas\sfound\sa\srare\sitem\s\((?P<item>.*?)\)\swith\sa\svalue\sof\s(?P<value>\d+)\s(?P<unit>PE(D|C))!\sA\srecord\shas\sbeen\sadded\sto\sthe\sHall\sof\sFame!$`)

// 0000-00-00 00:00:00 [System] [] You healed yourself 38.2 points
// date, channel, target, amount
var healRegexp = regexp.MustCompile(`^(?P<date>\d{4}-\d{2}-\d{2}\s\d{2}:\d{2}:\d{2})\s\[(?P<channel>\S+)\]\s\[\]\sYou\shealed\s(?P<target>.*?)\s(with\s)?(?P<amount>[\d.]+)\spoints$`)

// 0000-00-00 00:00:00 [System] [] Your enhancer Weapon Damage Enhancer 1 on your Omegaton M83 Predator broke. You have 246 enhancers remaining on the item. You received 0.8000 PED Shrapnel.
// date, channel, name, item, remaining, value
var enhancerBreakRegexp = regexp.MustCompile(`^(?P<date>\d{4}-\d{2}-\d{2}\s\d{2}:\d{2}:\d{2})\s\[(?P<channel>\S+)\]\s\[\] Your\senhancer\s(?P<name>.*)\son\syour\s(?P<item>.*)\sbroke\.\sYou\shave\s(?P<remaining>\d+)\senhancers\sremaining\son\sthe\sitem\.\sYou\sreceived\s(?P<value>.*)\sPED\sShrapnel\.$`)

// 0000-00-00 00:00:00 [System] [] Your Arsonistic Chip 2 (L) has reached tier 1.12
// date, channel, item, tier
var tierUpRegexp = regexp.MustCompile(`^(?P<date>\d{4}-\d{2}-\d{2}\s\d{2}:\d{2}:\d{2})\s\[(?P<channel>\S+)\]\s\[\]\sYour\s(?P<item>.*)\shas\sreached\stier\s(?P<tier>.*)$`)

var regexps = map[string]*regexp.Regexp{
  "attribute":                 attributeRegexp,
  "critical_damage_inflicted": criticalDamageInflictedRegexp,
  "critical_damage_taken":     criticalDamageTakenRegexp,
  "damage_inflicted":          damageInflictedRegexp,
  "damage_taken":              damageTakenRegexp,
  "enemy_dodge":               enemyDodgeRegexp,
  "enemy_evade":               enemyEvadeRegexp,
  "enemy_jam":                 enemyJamRegexp,
  "enemy_miss":                enemyMissRegexp,
  "enhancer_break":            enhancerBreakRegexp,
  "global":                    globalRegexp,
  "hall_of_fame":              hallOfFameRegexp,
  "heal":                      healRegexp,
  "loot":                      lootRegexp,
  "player_deflect":            playerDeflectRegexp,
  "player_dodge":              playerDodgeRegexp,
  "player_evade":              playerEvadeRegexp,
  "player_miss":               playerMissRegexp,
  "position":                  positionRegexp,
  "rare_loot":                 rareLootRegexp,
  "skill":                     skillRegexp,
  "skill_alt":                 skillAltRegexp,
  "tier_up":                   tierUpRegexp,
}
