# FastAPI Service - cost_services.py (Business Logic for Cost Retrieval)

from cache.redis_cache import get_cached_data, set_cached_data
from repository.database import fetch_cloud_costs
from models.cost_models import CostResponse

def fetch_cost_data(cloud_provider: str, start_date: str, end_date: str) -> CostResponse:
    """Fetch cloud cost data from cache or database."""
    cache_key = f"{cloud_provider}_costs:{start_date}_{end_date}"
    cached_data = get_cached_data(cache_key)
    
    if cached_data:
        return CostResponse(cached=True, data=cached_data)
    
    # Fetch data from database or external API
    cost_data = fetch_cloud_costs(cloud_provider, start_date, end_date)
    
    if cost_data:
        set_cached_data(cache_key, cost_data)
        return CostResponse(cached=False, data=cost_data)
    
    return CostResponse(cached=False, data={})
