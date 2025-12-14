import requests


GO_SCRAPER_URL = "http://localhost:8080/v1/scraper"

def fetch_reviews(url: str) -> dict:

    response = requests.get(GO_SCRAPER_URL,
                            params={'url': url},
                            timeout=10,
                            )
    response.raise_for_status()
    return response.json()