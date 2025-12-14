from typing import List
from pydantic import BaseModel


class ReviewInsight(BaseModel):
    pros: List[str]
    cons: List[str]
    sentiment_score: float
    confidence_score: float

class AnalyzeResponse(BaseModel):
    product: str
    verdict: str
    insights: ReviewInsight
    review_count: int