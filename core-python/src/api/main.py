from fastapi import FastAPI
from go.go_client import fetch_anything


app = FastAPI()

GO_BASE_URL = "https://localhost:8080"
@app.get("v1/analyze")

async def analyze(request: dict):

    meta = request["meta"]
    url = request["url"]

    go_resp = await fetch_anything(url)
    raw_bytes = go_resp.content
    content_type = go_resp.headers.get("content-type", "")
    if meta["site"] == "reddit":
        payload = go_resp.json()
    else :
        payload = raw_bytes.decode(errors="ignore")

    return {
        "site": meta["site"],
        "content_type": content_type,
        "size": len(raw_bytes),
    }