import { useState } from "react";

function App() {
  const [input, setInput] = useState("");
  const [result, setResult] = useState("");

  async function sendMessage() {
    const response = await fetch("http://localhost:8080/api/uppercase", {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify({ message: input })
    });

    const data = await response.json();
    setResult(data.result);
  }

  return (
    <div style={{ padding: "40px" }}>
      <h2>Uppercase Test</h2>

      <input
        value={input}
        onChange={(e) => setInput(e.target.value)}
        placeholder="Type something..."
      />

      <button onClick={sendMessage}>
        Send
      </button>

      <h3>Result:</h3>
      <p>{result}</p>
    </div>
  );
}

export default App;