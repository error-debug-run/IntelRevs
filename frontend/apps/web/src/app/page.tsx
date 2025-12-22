"use client";

import { useState } from "react";
import AnalyzeForm from "@/components/AnalyzeForm";
import ResultPanel from "@/components/ResultPanel";
import { analyzeReal, fakeAnalyze } from "@/lib/api";


export default function Home() {
  const [result, setResult] = useState<any>(null);
  const [loading, setLoading] = useState(false);

  async function handleAnalyze(input: string) {
    setLoading(true);
    setResult(null);

    try {
        const data =
            process.env.NODE_ENV === "development"
            ? await fakeAnalyze(input)
            : await analyzeReal(input);
        setResult(data);
      } finally {
        setLoading(false);
      }
  }

  return (
    <main style={{ padding: "2rem" }}>
      <h1>IntelRevs</h1>

      <AnalyzeForm onAnalyze={handleAnalyze} />

      {loading && <p>Analyzing…</p>}

      {result && <ResultPanel result={result} />}
    </main>
  );
}
