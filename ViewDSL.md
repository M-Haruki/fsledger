# View DSL

Basically you'll need all the items.

## Root

'base-date' is base date for relative dates.

```json
{
  "base-date": # string (yyyy-MM-dd),
  "rule": # Logic or Condition
}
```


## Logic

### Common (all view)

#### AND

```json
{
  "AND": [
    # Logic or Condition
  ]
}
```

#### OR

```json
{
  "OR": [
    # Logic or Condition
  ]
}
```

#### NOT

```json
{
  "NOT": # Logic or Condition
}
```

## Conditions

### Stocks (stock view only)

#### id

If you set some ids, they are treated as a OR operation.

```json
{
  "target": "stocks",
  "id": [
    # uuid
  ]
}
```

#### tags

If you set some tags, they are treated as a OR operation.

```json
{
  "target": "stocks",
  "tags": [
    # uuid
  ]
}
```


### Stocks (flow view only)

#### from stocks / id

If you set some ids, they are treated as a OR operation.

```json
{
  "target": "stocks-from",
  "id": [
    # uuid
  ]
}
```

#### from stocks / tag

If you set some tags, they are treated as a OR operation.

```json
{
  "target": "stocks-from",
  "tags": [
    # uuid
  ]
}
```

#### to stocks / id

If you set some ids, they are treated as a OR operation.

```json
{
  "target": "stocks-to",
  "id": [
    # uuid
  ]
}
```

#### to stocks / tag

If you set some tags, they are treated as a OR operation.

```json
{
  "target": "stocks-to",
  "tags": [
    # uuid
  ]
}
```

### Transactions (transaction view and flow view)

#### tag

If you set some tags, they are treated as a OR operation.

```json
{
  "target": "transactions",
  "tags": [
    # uuid
  ]
}
```

#### date / absolute / num

```json
{
  "target": "transactions",
  "mode": "absolute",
  "relation": # 'before' or 'before-or-equal' or 'point' or 'after-or-equal' or 'after',
  "unit" # 'day' or 'month' or 'year'
  "num": # int32
}
```

#### date / absolute / date

```json
{
  "target": "transactions",
  "mode": "absolute",
  "relation": # 'before' or 'before-or-equal' or 'point' or 'after-or-equal' or 'after',
  "unit" # 'year-month' or 'year-month-day'
  "date": # If unit is 'year-month', use the string (yyyy-MM); if it is 'year-month-day', use string (year-MM-dd).
}
```

#### date / relative

```json
{
  "target": "transactions",
  "mode": "relative",
  "relation": # 'before' or 'before-or-equal' or 'point' or 'after-or-equal' or 'after',
  "unit" # 'day' or 'month' or 'year'
  "num": # int32
}
```

### Flows (flow view only)

#### tag

If you set some tags, they are treated as a OR operation.

```json
{
  "target": "flows",
  "tags": [
    # uuid
  ]
}
```

#### amount

You must set either “amount-min” or “amount-max,” or both.
If both are set, they are treated as an AND operation.

```json
{
  "target": "flows",
  "amount-min": # int64,
  "amount-max": # int64
}
```

## Examples

### stock view

#### includes any of multiple tags, but excludes a specific ID

```json
{
  "base-date": "2026-01-01",
  "rule": {
    "AND": [
      {
        "target": "stocks",
        "tags": ["aaaa0000-0000-0000-0000-000000000001", "bbbb0000-0000-0000-0000-000000000002"]
      },
      {
        "NOT": {
          "target": "stocks",
          "id": ["cccc0000-0000-0000-0000-000000000003"]
        }
      }
    ]
  }
}
```

### transaction view

#### Multiple tags OR + Date lower bound

```json
{
  "base-date": "2026-01-01",
  "rule": {
    "AND": [
      {
        "target": "transactions",
        "tags": ["ffff0000-0000-0000-0000-000000000006", "aaaa1111-0000-0000-0000-000000000007"]
      },
      {
        "target": "transactions",
        "mode": "absolute",
        "relation": "after-or-equal",
        "unit": "year-month",
        "date": "2026-01"
      }
    ]
  }
}
```

#### within the last 3 months

```json
{
  "base-date": "2026-01-01",
  "rule": {
    "target": "transactions",
    "mode": "relative",
    "relation": "after-or-equal",
    "unit": "month",
    "num": -3
  }
}
```

### flow view

#### specific from/to tags + transaction tag + amount range

```json
{
  "base-date": "2026-01-01",
  "rule": {
    "AND": [
      { "target": "stocks-from", "tags": ["dddd0000-0000-0000-0000-000000000004"] },
      { "target": "stocks-to", "tags": ["eeee0000-0000-0000-0000-000000000005"] },
      { "target": "transactions", "tags": ["ffff0000-0000-0000-0000-000000000006"] },
      { "target": "flows", "amount-min": 1000, "amount-max": 5000 }
    ]
  }
}
```
