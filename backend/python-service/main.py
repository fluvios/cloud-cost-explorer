from fastapi import FastAPI
from routes import cost_routes

app = FastAPI()

# Register Routes
app.include_router(cost_routes.router)

@app.get("/health")
def health_check():
    return {"status": "ok"}

if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=8000)