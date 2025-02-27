# FastAPI Service - cost_models.py (Pydantic Models for Request/Response)

from pydantic import BaseModel, Field
from typing import Dict, Union

class CostRequest(BaseModel):
    cloud_provider: str = Field(..., description="Cloud provider (aws, gcp, azure, alibaba)")
    start_date: str = Field(..., regex="^\d{4}-\d{2}-\d{2}$", description="Start date in YYYY-MM-DD format")
    end_date: str = Field(..., regex="^\d{4}-\d{2}-\d{2}$", description="End date in YYYY-MM-DD format")

class CostResponse(BaseModel):
    cached: bool = Field(..., description="Indicates if data was retrieved from cache")
    data: Dict[str, Union[str, float]] = Field(..., description="Cloud cost data by date")
