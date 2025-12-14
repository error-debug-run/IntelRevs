# api/server.py
from fastapi import FastAPI

from src.core.models.requests import AnalyzeRequest
from src.core.models.response import AnalyzeResponse
from src.core.pipeline import run_pipeline

app = FastAPI(title="Product Review Assistant API")

@app.post("/v1/analyze/", response_model=AnalyzeResponse)


async def analyze(request: AnalyzeRequest):

    return await run_pipeline(request.input)