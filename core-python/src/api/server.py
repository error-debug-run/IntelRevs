# api/server.py
from fastapi import FastAPI
from src.core.pipeline import run_pipeline

app = FastAPI(title="Product Review Assistant API")

@app.post("/analyze/")
def analyze_review(data: dict):
    """
    data example:
    {
        "text": "I love this product, but the battery life is short",
        "language": "en"
    }
    """
    result = run_pipeline(data)
    return result
