# core/pipeline.py

from typing import Dict, Any, List

from src.core.services.scraper_client import fetch_reviews
from src.core.ocr.main import process_image
from src.core.nlp.main import analyze_review_text


def is_url(text: str) -> bool:
    return text.startswith(("http://", "https://"))


def is_image_path(text: str) -> bool:
    return text.lower().endswith((".png", ".jpg", ".jpeg"))


async def run_pipeline(input_data: str) -> Dict[str, Any]:
    """
    Unified pipeline entry point.

    Input types supported:
    - URL -> Go scraper
    - Image path -> OCR -> NLP
    - Raw text -> NLP
    """

    result: Dict[str, Any] = {
        "input_type": None,
        "raw_text": None,
        "analysis": None,
        "review_count": 0,
    }

    # -------- URL input --------
    if is_url(input_data):
        result["input_type"] = "url"

        scrape_data = fetch_reviews(input_data)
        reviews: List[str] = scrape_data.get("reviews", [])

        combined_text = "\n".join(reviews)

        result["raw_text"] = combined_text
        result["review_count"] = len(reviews)

    # -------- Image input --------
    elif is_image_path(input_data):
        result["input_type"] = "image"

        extracted_text = process_image(input_data)
        result["raw_text"] = extracted_text

    # -------- Plain text --------
    else:
        result["input_type"] = "text"
        result["raw_text"] = input_data

    # -------- NLP Analysis --------
    if result["raw_text"]:
        result["analysis"] = analyze_review_text(result["raw_text"])
    else:
        result["analysis"] = {}

    return result
