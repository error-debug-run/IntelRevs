"use client";

import { useState } from "react";

type AnalyzeFormProps = {
  onAnalyze: (input: string) => void;
};

export default function AnalyzeForm({ onAnalyze }: AnalyzeFormProps) {
  const [input, setInput] = useState("");

  return (
    <>
      <textarea
        placeholder="Paste a product URL or review text"
        style={{ width: "100%", height: "100px" }}
        value={input}
        onChange={(e) => setInput(e.target.value)}
      />

      <br />
      <br />

      <button
        style={{ cursor: "pointer" }}
        onClick={() => onAnalyze(input)}
      >
        Analyze
      </button>
    </>
  );
}
