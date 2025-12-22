type ResultPanelProps = {
  result: {
    input_type: string;
    summary: string;
    sentiment: string;
    review_count: number;
  };
};

export default function ResultPanel({ result }: ResultPanelProps) {
  return (
    <div style={{ marginTop: "2rem", padding: "1rem", border: "1px solid #ccc" }}>
      <h3>Analysis Result</h3>

      <p>
        <strong>Input Type:</strong> {result.input_type}
      </p>
      <p>
        <strong>Sentiment:</strong> {result.sentiment}
      </p>
      <p>
        <strong>Review Count:</strong> {result.review_count}
      </p>

      <p>
        <strong>Summary:</strong>
      </p>
      <p>{result.summary}</p>
    </div>
  );
}
