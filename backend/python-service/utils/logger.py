import logging
import os

# Define log level
LOG_LEVEL = os.getenv("LOG_LEVEL", "INFO").upper()

# Configure logger
logging.basicConfig(
    format="%(asctime)s - %(levelname)s - %(message)s",
    level=getattr(logging, LOG_LEVEL, logging.INFO),
    handlers=[
        logging.StreamHandler(),  # Log to console
    ]
)

logger = logging.getLogger(__name__)

def log_info(message: str):
    """Log an info message."""
    logger.info(message)

def log_warning(message: str):
    """Log a warning message."""
    logger.warning(message)

def log_error(message: str):
    """Log an error message."""
    logger.error(message)
