"use client";

import { useState } from "react";

export default function Home() {

    const [input, setInput] = useState("");
    const [submitted, setSubmitted] = useState("");

    return (
        <main style={{padding: "2rem"}}>
            <h1>IntelRevs</h1>

            <textarea
                placeholder="Paste a product URL or review text"
                style={{width: "100%", height: "100px"}}
                value={input}
                onChange={(e) => setInput(e.target.value)}
            />

            <br/>
            <br/>

            <button
                style={{ cursor:"pointer" }}
                onClick={() => setSubmitted(input)}
            >
                Analyze
            </button>

            {submitted && (
                <>
                    <hr style={{margin:"2rem 0"}} />
                    <h3>Submitted Input</h3>
                    <pre>{submitted}</pre>
                </>
            )}
        </main>
    );
}
