import asyncio
import os
import sys

from src.core.pipeline import run_pipeline


def main():
    if len(sys.argv) < 2:
        print("Usage: python main.py <input-data>")
        sys.exit(1)

    input_data = " ".join(sys.argv[1:])
    result = asyncio.run(run_pipeline(input_data))
    print("|-------------RESULTS--------------|")
    print(result)
    print("|------------------------------------|")

if __name__ == "__main__":
    main()