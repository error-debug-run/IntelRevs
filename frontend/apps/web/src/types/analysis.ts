export interface AnalysisResult {
  input_type: "url" | "text" | "image";
  analysis: Record<string, any>;
  review_count?: number;
}
