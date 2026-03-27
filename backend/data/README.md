# Supply Chain Dataset

Place your dataset file here:

```
data/supply_chain_dataset_full_500k.csv
```

## Expected CSV Columns

| Column name              | Type    | Example value    |
|--------------------------|---------|------------------|
| shipment_id              | string  | SHP100001        |
| origin                   | string  | Mumbai           |
| destination              | string  | Delhi            |
| distance_km              | float   | 1400             |
| carrier                  | string  | BlueDart         |
| mode                     | string  | road             |
| expected_delivery_date   | date    | 2024-01-15       |
| delivered_date           | date    | 2024-01-17       |
| weather_severity         | float   | 6.5 (0–10 scale) |
| traffic_condition        | float   | 7.0 (0–10 scale) |

Accepted date formats: `YYYY-MM-DD`, `MM/DD/YYYY`, `DD-MM-YYYY`, `YYYY/MM/DD`.

The server will not start properly without this file.
