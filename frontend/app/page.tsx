"use client";

import { useState } from "react";
import { ConnectError } from "@connectrpc/connect";
import { calculatorClient } from "./client";

export default function Home() {
  const [a, setA] = useState("");
  const [b, setB] = useState("");
  const [operation, setOperation] = useState("add");
  const [result, setResult] = useState("");
  const [error, setError] = useState("");

  async function Calculate(e: React.SubmitEvent) {
    e.preventDefault();
    setResult("");
    setError("");

    try {
      const res = await calculatorClient.calculate({a: Number(a), b: Number(b), operation });
      setResult(res.expression);
    } catch (err) {
      setError(err instanceof ConnectError ? err.message : "somethings wrong");
    }
  }

  return (
    <main>
      <h2>Calculator</h2>

      <form onSubmit={Calculate}>
        <input type="number" value={a} onChange={(e)=> setA(e.target.value)}  required />

        <select value={operation} onChange={(e) => setOperation(e.target.value)}>
          <option value="add">+</option>
          <option value="subtract">−</option>
          <option value="multiply">×</option>
          <option value="divide">÷</option>
        </select>

        <input type="number" value={b} onChange={(e) => setB(e.target.value)} required />

        <button type="submit">=</button>
      </form>

      {result && <p>{result}</p>}
      {error && <p className="error">{error}</p>}
    </main>
  );
}
