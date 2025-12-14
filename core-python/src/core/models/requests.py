from typing import Optional, List
from pydantic import BaseModel

class AnalyzeRequest(BaseModel):
    input: str
    product_url: Optional[str] = None
    raw_reviews: Optional[List[str]] = None
    Language: Optional[str] = "en"
    preferencse: Optional[dict] = {}
    budget: Optional[float] = None