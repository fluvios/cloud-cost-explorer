# FastAPI Service - database.py (Database Connection & Queries)

import os
import psycopg2
from psycopg2.extras import RealDictCursor
from typing import Dict

def get_db_connection():
    """Establish a connection to PostgreSQL database."""
    return psycopg2.connect(
        dbname=os.getenv("DB_NAME", "cloud_costs"),
        user=os.getenv("DB_USER", "postgres"),
        password=os.getenv("DB_PASSWORD", "password"),
        host=os.getenv("DB_HOST", "localhost"),
        port=os.getenv("DB_PORT", "5432"),
        cursor_factory=RealDictCursor
    )

def fetch_cloud_costs(cloud_provider: str, start_date: str, end_date: str) -> Dict[str, float]:
    """Fetch cloud cost data from the database."""
    query = """
        SELECT date, cost FROM cloud_costs
        WHERE provider = %s AND date BETWEEN %s AND %s
        ORDER BY date;
    """
    
    with get_db_connection() as conn:
        with conn.cursor() as cursor:
            cursor.execute(query, (cloud_provider, start_date, end_date))
            result = cursor.fetchall()
    
    return {row["date"]: row["cost"] for row in result}
