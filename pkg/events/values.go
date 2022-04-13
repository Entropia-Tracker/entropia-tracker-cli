package events

type LootEventValues struct {
  Name   string  `json:"name"`
  Amount int64   `json:"amount"`
  Value  float64 `json:"value"`
}

type PointsGainedEventValues struct {
  Name  string  `json:"name"`
  Value float64 `json:"value"`
  Type  string  `json:"type"`
}

type DamageInflictedEventValues struct {
  Amount   float64 `json:"amount"`
  Critical bool    `json:"critical"`
}

type DamageTakenEventValues struct {
  Amount   float64 `json:"amount"`
  Critical bool    `json:"critical"`
}

type EnemyEvadeEventValues struct {
  Reason string `json:"reason"`
}

type EnhancerBreakEventValues struct {
  Name      string  `json:"name"`
  Item      string  `json:"item"`
  Remaining int64   `json:"remaining"`
  Value     float64 `json:"value"`
}

type SpecialLootEventValues struct {
  Value    float64 `json:"value"`
  Player   string  `json:"player"`
  Enemy    *string `json:"enemy"`
  Item     *string `json:"item"`
  Location *string `json:"location"`
  Type     string  `json:"type"`
}

type HealEventValues struct {
  Amount float64 `json:"amount"`
  Target string  `json:"target"`
}

type PlayerEvadeEventValues struct {
  Reason string `json:"reason"`
}

type PositionEventValues struct {
  Lat  int64   `json:"lat"`
  Lon  int64   `json:"lon"`
  Alt  *int64  `json:"alt"`
  Name *string `json:"name"`
}

type TierUpEventValues struct {
  Item string  `json:"item"`
  Tier float64 `json:"tier"`
}
