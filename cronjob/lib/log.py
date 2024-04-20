import logging


FORMAT = "%(asctime)s %(module)s %(levelname)s %(message)s"
logging.basicConfig(format=FORMAT, level=logging.INFO)


def setup_logging(name) -> logging.Logger:
    logger = logging.getLogger(name)
    return logger
