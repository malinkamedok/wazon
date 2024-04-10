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

#### Получение подробной информации о конкретном продукте

<details>
 <summary><code>GET</code> <code><b>/</b></code> <code>delivery</code> <code><b>/</b></code> <code>UUID</code></summary>

##### Parameters

> | name | type     | data type | example                                  | description                 |
> |------|----------|-----------|------------------------------------------|-----------------------------|
> | UUID | required | string    | `8d7b317b-db39-4751-b0a4-8efeafc0cacd`   | Version 4 UUID              |


##### Example output

```json
{
    "8d7b317b-db39-4751-b0a4-8efeafc0cacd": {
      "status": "delivery",
      "created_at": 1712783228,
      "updated_at": 1712784228
    },
    "service": "delivery"
}
```

</details>

#### Обновление статуса заказа

<details>
 <summary><code>POST</code> <code><b>/</b></code> <code>delivery</code> <code><b>/</b></code> <code>UUID</code> <code><b>/</b></code> <code>edit</code></summary>

##### Parameters

> | name | type     | data type | example                                  | description                 |
> |------|----------|-----------|------------------------------------------|-----------------------------|
> | UUID | required | string    | `8d7b317b-db39-4751-b0a4-8efeafc0cacd`   | Version 4 UUID              |
> | Status | optional | string    | `await`   | Status of order              |



##### Example output

```json
{
    "8d7b317b-db39-4751-b0a4-8efeafc0cacd": {
      "status": "await",
      "created_at": 1712783228,
      "updated_at": 1712784437
    },
    "service": "delivery"
}
```

</details>