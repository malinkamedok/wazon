from . import log


logger = log.setup_logging(__name__)

NEXT_STATUS_MAP = {
    "created": "prepare",
    "prepare": "delivery",
    "delivery": "await",
    "await": "received",
    "received": "received",
}


def get_new_status(prev_status: str) -> str:
    return NEXT_STATUS_MAP[prev_status]


def update_order(order: dict) -> dict:
    res_order = order.copy()
    res_order["order_status"] = get_new_status(order["order_status"])

    if res_order == order:
        return None

    logger.info(
        "Updating status of order %s: was %s, new %s",
        order["uuid"],
        order["order_status"],
        res_order["order_status"],
    )

    return res_order
