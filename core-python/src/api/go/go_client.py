import httpx
from typing import Any, Dict

GO_BASE_URL = "http://localhost:8080"


async def fetch_anything(request: Dict[str, Any]) -> str:
    """
    Sends a generic fetch request to the Go server and returns raw response data.
    """

    url: str = request["url"]
    method: str = request.get("method", "GET")
    meta: dict = request.get("meta", {})

    headers = {
        "User-Agent": "IntelRevs/0.1",
        "Content-Type": "application/json",
    }

    async with httpx.AsyncClient(timeout=30.0) as client:
        resp = await client.request(
            method=method,
            url=f"{GO_BASE_URL}/v1/scraper",
            params={
                "url": url,
                **meta,
            },
            headers=headers,
        )

    resp.raise_for_status()
    return resp.text
