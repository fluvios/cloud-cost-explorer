# FastAPI Service - redis_cache.py (Redis Caching Logic)

import os
import redis
import json
from typing import Dict, Optional

# Initialize Redis connection
redis_client = redis.Redis(
    host=os.getenv("REDIS_HOST", "localhost"),
    port=int(os.getenv("REDIS_PORT", 6379)),
    db=0,
    decode_responses=True
)

def get_cached_data(cache_key: str) -> Optional[Dict[str, float]]:
    """Retrieve cached data from Redis."""
    cached_data = redis_client.get(cache_key)
    if cached_data:
        return json.loads(cached_data)
    return None

def set_cached_data(cache_key: str, data: Dict[str, float], expiration: int = 3600):
    """Store data in Redis cache with expiration time (default: 1 hour)."""
    redis_client.setex(cache_key, expiration, json.dumps(data))