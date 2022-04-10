# Events

All events share a common format, the table below is the content in the `values`
key.

```json
{
    "event": "event_name",
    "date": "YYYY-MM-DD HH:mm:ss",
    "channel": "system",
    "values": {
      "key": <value>
    }
}
```

## Values

| Event                | Key       | Type     | Example value              | Description                                            |
|----------------------|-----------|----------|----------------------------|--------------------------------------------------------|
| **points_gained**    |           |          |                            | Attribute or skill points gained                       |
|                      | name      | string   | `Agility`                  | Name                                                   |
|                      | value     | float    | `50.8`                     | Points value                                           |
|                      | type      | string   | `skill`                    | Type of point `skill` or `attribute`                   |
| **damage_inflicted** |           |          |                            | Player inflicted damage                                |
|                      | amount    | float    | `50.0`                     | Amount of damage inflicted                             |
|                      | critical  | bool     | `true`                     | Critical hit                                           |
| **damage_taken**     |           |          |                            | Player took damage                                     |
|                      | amount    | float    | `50.0`                     | Amount of damage taken                                 |
|                      | critical  | bool     | `false`                    | Critical hit                                           |
| **enemy_evade**      |           |          |                            | Enemy evaded damage                                    |
|                      | reason    | string   | `dodge`                    | Possible values: `dodge`, `evade` or `jam`             |
| **enemy_miss**       |           |          |                            | Enemy missed an attack                                 |
| **enhancer_break**   |           |          |                            | Enhancer broke                                         |
|                      | name      | string   | `Weapon Damage Enhancer 1` | Name of the enhancer                                   |
|                      | item      | string   | `Omegaton M83 Predator`    | Item it was attached to                                |
|                      | remaining | int      | `200`                      | Currently remaining on item                            |
|                      | value     | float    | `0.8000`                   | PED value                                              |
| **special_loot**     |           |          |                            | Player got a special loot**                            |
|                      | value     | int      | `15`                       | PED value                                              |
|                      | player    | string   | `Some Player Name`         | Name of the player                                     |
|                      | enemy     | string?  | `Kerberos Young`           | Enemy name                                             |
|                      | item      | string?  | `Holy Grail`               | Only available on `rare` loots                         |
|                      | location  | string?  | `Some cave`                | Location
|                      | type      | string   | `global`                   | Possible values: `global`, `hall_of_fame` or `rare`    |
| **heal**             |           |          |                            | Player healed someone                                  |
|                      | target    | string   | `Some Player Name`         | Name of the player healed*                             |
|                      | amount    | float    | `50.0`                     | Amount healed                                          |
| **loot**             |           |          |                            | Player received loot                                   |
|                      | name      | string   | `Shrapnel`                 | Looted item                                            |
|                      | amount    | int      | `50`                       | Item amount                                            |
|                      | value     | float    | `5.50`                     | PED value                                              |
| **player_evade**     |           |          |                            | Player evaded damage                                   |
|                      | reason    | string   | `dodge`                    | Possible values: `dodge`, `evade` or `deflect`         |
| **player_miss**      |           |          |                            | Player missed an attack                                |
| **position**         |           |          |                            | Player position                                        |
|                      | lat       | int      | `12345`                    | Latitude                                               |
|                      | lon       | int      | `6789`                     | Longitude                                              |
|                      | alt       | int?     | `50`                       | Altitude                                               |
| **tier_up**          |           |          |                            | Item gained a tier                                     |
|                      | item      | string   | `Arsonistic Chip 2 (L)`    | Item name                                              |
|                      | tier      | float    | `1.12`                     | New tier                                               |


*Note: Types marked with `?` are optional and can be `null`, for example `string?`.*

*\* Note: Player name must be supplied or heals on yourself will return `yourself` instead of your name.*

*\*\* Note: Player name must be supplied or these will not be tracked.*
