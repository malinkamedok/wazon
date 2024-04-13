## Микросервис витрины магазина

Работает с базой данных, хранящей сведения о заказах.

#### Получение информации о статусе всех заказов

<details>
 <summary><code>GET</code> <code><b>/</b></code> <code>delivery</code></summary>

##### Example output

```json
{
    "orders": {
      "8d7b317b-db39-4751-b0a4-8efeafc0cacd": "created",
      "5acbbf8a-fec9-48d1-b326-08c71ff64b59": "delivery",
      ...
    },
    "service": "delivery"
}
```

</details>

#### Получение подробной информации о конкретной доставке

<details>
 <summary><code>GET</code> <code><b>/</b></code> <code>delivery</code> <code><b>/</b></code> <code>UUID</code></summary>

##### Parameters

> | name | type     | data type | example                                  | description                 |
> |------|----------|-----------|------------------------------------------|-----------------------------|
> | UUID | required | string    | `0d39d5a6-ca6d-4ef2-b477-621e1b37e526`   | Version 4 UUID              |


##### Example output

```json
{
  "order": {
    "uuid": "0d39d5a6-ca6d-4ef2-b477-621e1b37e526",
    "order_status": "created",
    "created_at": "2024-04-12T22:00:46.464337Z",
    "updated_at": "2024-04-12T22:00:46.464337Z"
  },
  "service": "delivery"
}
```

</details>

#### Создание заказа

<details>
 <summary><code>POST</code> <code><b>/</b></code> <code>delivery</code> <code><b>/</b></code> <code>create</code> </summary>

##### Parameters

> | name | type     | data type | example                                  | description                 |
> |------|----------|-----------|------------------------------------------|-----------------------------|
> | UUID | required | string    | `0d39d5a6-ca6d-4ef2-b477-621e1b37e526`   | Version 4 UUID              |


##### Example output

```json
{
  "order": {
    "uuid": "0d39d5a6-ca6d-4ef2-b477-621e1b37e526",
    "order_status": "created",
    "created_at": "2024-04-12T22:00:46.464337Z",
    "updated_at": "2024-04-12T22:00:46.464337Z"
  },
  "service": "delivery"
}
```

</details>

#### Обновление статуса заказа

<details>
 <summary><code>POST</code> <code><b>/</b></code> <code>delivery</code> <code><b>/</b></code> <code>update</code></summary>

##### Parameters

> | name | type     | data type | example                                  | description                 |
> |------|----------|-----------|------------------------------------------|-----------------------------|
> | UUID | required | string    | `0d39d5a6-ca6d-4ef2-b477-621e1b37e526`   | Version 4 UUID              |
> | Status | optional | string    | `await`   | Status of order              |



##### Example output

```json
{
  "order": {
    "uuid": "0d39d5a6-ca6d-4ef2-b477-621e1b37e526",
    "order_status": "created",
    "created_at": "2024-04-12T22:00:46.464337Z",
    "updated_at": "2024-04-12T22:21:21.32669Z"
  },
  "service": "delivery"
}
```

</details>