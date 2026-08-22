![brand image](./assets/brand.png)

# FSLedger

A highly customizable personal ledger based on the Flow &amp; Stock model.

## ToDo
- DB indexing
- Multi User & Login System

## Memo

- Store the currency exponent with the currency, and use it to scale the integer amount for display in the UI.

## Database

All columns are `NOT NULL`, and empty strings (`""`) are used in some text columns.
If a user wants to delete a stock record, they must select a fallback stock for any flows that currently use it.

```mermaid
---
title: "FSLedger ER Diagram"
---

erDiagram
    stocks{
        uuid id PK "DEFAULT gen_random_uuid()"
        text name "UNIQUE"
        boolean has_amount "Whether the total amount is meaningful"
        text currency
        integer currency_exponent
        text description
    }

    stock_tags{
        uuid id PK "DEFAULT gen_random_uuid()"
        text name "UNIQUE"
    }

    stock_tag_relations{
        uuid stock_id PK "FK stocks.id (ON DELETE CASCADE)"
        uuid tag_id PK "FK stock_tags.id (ON DELETE CASCADE)"
    }
    
    transactions{
        uuid id PK "DEFAULT gen_random_uuid()"
        text description
        date occurred_at
    }

    transaction_tags{
        uuid id PK "DEFAULT gen_random_uuid()"
        text name "UNIQUE"
    }

    transaction_tag_relations{
        uuid transaction_id PK "FK transactions.id (ON DELETE CASCADE)"
        uuid tag_id PK "FK transaction_tags.id (ON DELETE CASCADE)"
    }

    flows{
        uuid id PK "DEFAULT gen_random_uuid()"
        uuid transaction_id "FK transactions.id (ON DELETE CASCADE)"
        uuid from_stock_id "FK stocks.id (ON DELETE NO ACTION)"
        uuid to_stock_id "FK stocks.id (ON DELETE NO ACTION)"
        bigint amount
    }

    flow_tags{
        uuid id PK "DEFAULT gen_random_uuid()"
        text name "UNIQUE"
    }

    flow_tag_relations{
        uuid flow_id PK "FK flows.id (ON DELETE CASCADE)"
        uuid tag_id PK "FK flow_tags.id (ON DELETE CASCADE)"
    }

    stocks ||--o{ stock_tag_relations : has
    stock_tags ||--o{ stock_tag_relations : has

    transactions ||--o{ transaction_tag_relations : has
    transaction_tags ||--o{ transaction_tag_relations : has

    flows ||--o{ flow_tag_relations : has
    flow_tags ||--o{ flow_tag_relations : has

    transactions ||--o{ flows : contains
    stocks ||--o{ flows : "to"
    stocks ||--o{ flows : "from"
```
