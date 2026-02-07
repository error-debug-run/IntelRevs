import argparse
import asyncio
from dataclasses import field, dataclass

from api.go.go_client import fetch_anything


@dataclass
class FetchRequest:
    url: str
    method: str = "GET"
    meta: dict = field(default_factory=dict)


async def main():
    parser = argparse.ArgumentParser("intelrevs-raw")

    parser.add_argument("--url", required=True)
    parser.add_argument("--site", required=True)
    parser.add_argument("--dump", action="store_true")
    parser.add_argument("--out", default="raw_output.bin")

    args = parser.parse_args()

    request = FetchRequest(
        url=args.url,
        method="GET",
        meta={"site": args.site},
    )

    raw = await fetch_anything(request.__dict__)

    print("\n=== GO RESPONSE ===")
    print("Type:", type(raw))
    print("Size (bytes):", len(raw.encode()))

    if args.dump:
        with open(args.out, "wb") as f:
            f.write(raw.encode())
        print(f"Raw body saved to {args.out}")

    print("\n=== PREVIEW (first 500 chars) ===")
    print(raw[:500])


if __name__ == "__main__":
    asyncio.run(main())
