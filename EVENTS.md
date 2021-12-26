# Events

All events share a common format, the table below is the content in the `values`
key.

```json
{
    "event": "event_name",
    "date": "YYYY-MM-DD HH:mm:ss",
    "channel": "system",
    "values": {
      "key": "value"
    }
}
```

## Values


| Event            | Key      | Example value      | Description                                |
|------------------|----------|--------------------|--------------------------------------------|
| attribute        | `-`      | `-`                | Attribute points gained                    |
|                  | name     | `Agility`          | Attribute name                             |
|                  | value    | `10.0`             | Number of points gained                    |
| damage_inflicted | `-`      | `-`                | Player inflicted damage                    |
|                  | amount   | `50.0`             | Amount of damage inflicted                 |
|                  | critical | `0`                | Critical hit, available values: `1` or `0` |
| damage_taken     | `-`      | `-`                | Player took damage                         |
|                  | amount   | `50.0`             | Amount of damage taken                     |
|                  | critical | `0`                | Critical hit, available values: `1` or `0` |
| enemy_dodge      | `-`      | `-`                | Enemy dodged player attack                 |
| enemy_evade      | `-`      | `-`                | Enemy evaded player attack                 |
| enemy_jam        | `-`      | `-`                | Enemy jammed player attack                 |
| enemy_miss       | `-`      | `-`                | Enemy missed an attack                     |
| global           | `-`      | `-`                | Player got a global**                      |
|                  | player   | `Some Player Name` | Name of the player                         |
|                  | enemy    | `Kerberos Young`   | Mob that dropped it                        |
|                  | value    | `15`               | PED value of the global                    |
| hall_of_fame     | `-`      | `-`                | Player got a Hall of Fame entry**          |
|                  | player   | `Some Player Name` | Name of the player                         |
|                  | enemy    | `Kerberos Mature`  | Mob that dropped it                        |
|                  | value    | `1500`             | PED value                                  |
| heal             | `-`      | `-`                | Player healed someone                      |
|                  | target   | `Some Player Name` | Name of the player healed*                 |
|                  | amount   | `50.0`             | Amount healed                              |
| loot             | `-`      | `-`                | Player received loot                       |
|                  | name     | `Shrapnel`         | Item name                                  |
|                  | amount   | `50`               | Item amount                                |
|                  | value    | `5.50`             | PED value                                  |
| player_deflect   | `-`      | `-`                | Player deflected damage                    |
| player_dodge     | `-`      | `-`                | Player dodged an attack                    |
| player_evade     | `-`      | `-`                | Player evaded an attack                    |
| player_miss      | `-`      | `-`                | Player missed an attack                    |
| position         | `-`      | `-`                | Player position                            |
|                  | lat      | `12345`            | Latitude                                   |
|                  | lon      | `6789`             | Longitude                                  |
|                  | alt      | `50`               | Altitude                                   |
| rare_loot        | `-`      | `-`                | Player received rare loot**                |
|                  | player   | `Some Player Name` | Player name                                |
|                  | item     | `Holy Grail`       | Item name                                  |
|                  | value    | `5000`             | PED value                                  |
| skill            | `-`      | `-`                | Player gained skill points                 |
|                  | name     | `Wounding`         | Skill name                                 |
|                  | value    | `6.5432`           | Points value                               |


*\* Note: Player name must be supplied or heals on yourself will return `yourself` instead of your name.*

*\*\* Note: Player name must be supplied or these will not be tracked.*
