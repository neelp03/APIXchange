<script>
  let status = "Not Checked";

  async function checkHealth() {
    try {
      const response = await fetch("http://localhost:8080/health");
      if (response.ok) {
        status = await response.text();
      } else {
        status = "Failed to connect.";
      }
    } catch (error) {
      status = "Error: " + (error instanceof Error ? error.message : String(error));
    }
  }
</script>

<main>
  <h1>APIXchange Frontend</h1>
  <p>Status: {status}</p>
  <button on:click={checkHealth}>Check Backend Health</button>
</main>

<style>
  main {
    font-family: Arial, sans-serif;
    text-align: center;
    margin-top: 50px;
  }
  button {
    padding: 10px 20px;
    font-size: 16px;
    cursor: pointer;
  }
</style>
