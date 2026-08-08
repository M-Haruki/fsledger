![brand image](./assets/brand.png)

# FSLedger

A highly customizable personal ledger based on the Flow &amp; Stock model.

## Database

All columns are `NOT NULL`, and empty strings (`""`) are used in some text columns.

```mermaid
---
title: "FSLedger ER Diagram"
---

erDiagram
    stocks{
        uuid uuid PK "DEFAULT gen_random_uuid()"
        text name
        boolean has_amount "Whether the total amount is meaningful"
        text currency
        text desc
    }

    stock_tag{
        uuid uuid PK "DEFAULT gen_random_uuid()"
        text name
    }

    stock_tag_relation{
        uuid stock_uuid PK "FK stocks.uuid"
        uuid tag_uuid PK "FK stock_tag.uuid"
    }
    
    transactions{
        uuid uuid PK "DEFAULT gen_random_uuid()"
        text desc
        timestamptz occurred_at "DEFAULT now()"
    }

    transaction_tag{
        uuid uuid PK "DEFAULT gen_random_uuid()"
        text name
    }

    transaction_tag_relation{
        uuid transaction_uuid PK "FK transactions.uuid"
        uuid tag_uuid PK "FK transaction_tag.uuid"
    }

    flows{
        uuid uuid PK "DEFAULT gen_random_uuid()"
        uuid transaction_uuid "FK transactions.uuid"
        uuid from_stock_uuid "FK stocks.uuid"
        uuid to_stock_uuid "FK stocks.uuid"
        numeric amount
    }

    flow_tag{
        uuid uuid PK "DEFAULT gen_random_uuid()"
        text name
    }

    flow_tag_relation{
        uuid flow_uuid PK "FK flows.uuid"
        uuid tag_uuid PK "FK flow_tag.uuid"
    }

    stocks ||--o{ stock_tag_relation : has
    stock_tag ||--o{ stock_tag_relation : has

    transactions ||--o{ transaction_tag_relation : has
    transaction_tag ||--o{ transaction_tag_relation : has

    flows ||--o{ flow_tag_relation : has
    flow_tag ||--o{ flow_tag_relation : has

    transactions ||--o{ flows : contains
    stocks ||--o{ flows : "to"
    stocks ||--o{ flows : "from"
```
