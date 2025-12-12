# core/pipeline.py

from typing import Optional, Dict

from src.core.ocr.main import process_image
from src.core.scraping.main import fetch_reviews_from_url
from src.core.nlp.main import analyze_review_text

def is_url(input_text: str) -> bool:
    return input_text.startswith("http://") or input_text.startswith("https://")


def is_image_path(input_text: str) -> bool:
    return input_text.lower().endswith((".png", ".jpg", ".jpeg"))


async def run_pipeline(input_data: str) -> Dict:
    """
    Decides the workflow based on the input:
    - URL -> scraping
    - Image -> OCR then NLP
    - Plain text -> NLP only
    """

    result = {"input_type": None, "analysis": None, "raw_text": None}

    # URL → Scrape reviews
    if is_url(input_data):
        result["input_type"] = "url"
        text = await fetch_reviews_from_url(input_data)
        result["raw_text"] = text

    # Image → OCR + NLP
    elif is_image_path(input_data):
        result["input_type"] = "image"
        text = process_image(input_data)
        result["raw_text"] = text

    # Text → NLP directly
    else:
        result["input_type"] = "text"
        text = input_data
        result["raw_text"] = input_data

    # Send text to NLP processor
    analysis = analyze_review_text(text)
    result["analysis"] = analysis

    return result
