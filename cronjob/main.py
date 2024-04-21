import time
import os
import random

import lib.log as log
import lib.delivery as delivery
import lib.status as status


logger = log.setup_logging(__name__)

HOST_DEFAULT_VALUE = "delivery"
WAIT_BETWEEN_UPDATES_SEC = int(os.environ.get("WAIT_BETWEEN_UPDATES_SEC", 1))


def _resolve_delivery():
    port = os.environ.get("DELIVERY_PORT", None)
    if port is None:
        raise Exception("No port provided")

    host = os.environ.get("DELIVERY_HOST", None)
    if host is None:
        logger.warning(f"No host provided, defaulting to {HOST_DEFAULT_VALUE}")
        host = HOST_DEFAULT_VALUE

    address = "http://{}:{}/delivery".format(host, port)

    return delivery.Delivery(address)


if __name__ == "__main__":
    service = _resolve_delivery()

    while True:
        orders = service.get_orders()

        if orders is not None:
            order = random.choice(orders)

            new_order = status.update_order(order)

            if new_order is not None:
                service.update_order(**new_order)

        time.sleep(WAIT_BETWEEN_UPDATES_SEC)
