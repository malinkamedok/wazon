import requests
import json


class Delivery:
    def __init__(self, address) -> None:
        self.address = address

    def get_orders(self):
        r = requests.get(self.address)
        assert r.ok

        return r.json()["orders"]

    def update_order(self, uuid, order_status):
        query_data = json.dumps(
            {
                "uuid": uuid,
                "status": order_status,
            }
        )

        r = requests.post(f"{self.address}/update", data=query_data)

        assert r.ok, r.request.body
        return r.json()["order"]
