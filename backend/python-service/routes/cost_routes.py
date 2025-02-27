# FastAPI Service - cost_routes.py (API Routes for Cost Fetching)

from fastapi import APIRouter, Depends, Query
from models.cost_models import CostRequest, CostResponse
from services.cost_services import fetch_cost_data

router = APIRouter(prefix="/costs", tags=["Cost Data"])

@router.get("/", response_model=CostResponse)
def get_cost_data(
    cloud_provider: str = Query(..., description="Cloud provider (aws, gcp, azure, alibaba)"),
    start_date: str = Query(..., regex="^\d{4}-\d{2}-\d{2}$", description="Start date in YYYY-MM-DD format"),
    end_date: str = Query(..., regex="^\d{4}-\d{2}-\d{2}$", description="End date in YYYY-MM-DD format"),
):
    """Fetch cost data for a specific cloud provider between a given date range."""
    return fetch_cost_data(cloud_provider, start_date, end_date)
