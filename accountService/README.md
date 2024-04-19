## Микросервис витрины магазина

Работает с базой данных, хранящей сведения о продукции.

По Rest запросу отдает как список всех имеющихся продуктов, так и информацию о конкретном.

#### Получение информации о всех продуктах

<details>
 <summary><code>GET</code> <code><b>/</b></code> <code>storefront</code></summary>

##### Example output

```json
{
    "products": {
      "8d7b317b-db39-4751-b0a4-8efeafc0cacd": "Holographic sticker Miyagi Endspiel logo",
      "5acbbf8a-fec9-48d1-b326-08c71ff64b59": "The Book of Statham's Quotes",
      ...
    },
    "service": "storefront"
}
```

</details>

#### Получение подробной информации о конкретном продукте

<details>
 <summary><code>GET</code> <code><b>/</b></code> <code>storefront</code><code><b>/</b></code> <code>UUID</code></summary>

##### Parameters

> | name | type     | data type | example                                  | description                 |
> |------|----------|-----------|------------------------------------------|-----------------------------|
> | UUID | required | string    | `8d7b317b-db39-4751-b0a4-8efeafc0cacd`   | Version 4 UUID              |


##### Example output

```json
{
    "8d7b317b-db39-4751-b0a4-8efeafc0cacd": {
      "name": "Holographic sticker Miyagi Endspiel logo",
      "description": "Type: Bank card sticker. Color: Holographic metallic grey",
      "price": 228
    },
    "service": "storefront"
}
```

</details>