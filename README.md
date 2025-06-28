<h1>One Piece API + Backend</h1>

<h2>How to Use</h2>

<h3>Setting up the database:</h3>
<ol>
  <li>Connect to the MySQL CLI and create a database (you will need to reflect your choices in <code>setupDB()</code> in <code>db.go</code>).</li>
  <li>Run the <code>db/schema.sql</code> script in the MySQL CLI.</li>
  <li>Run the <code>db/inserts.sql</code> script in the MySQL CLI.</li>
</ol>

<h3>Starting the web service:</h3>
<ol>
  <li>Open a new terminal.</li>
  <li>Navigate to <code>/web-service</code> and run <code>go run .</code> in your terminal.</li>
</ol>

<h3>Optionally, you can use my pre-made tests:</h3>
<ol>
  <li>Create a virtual environment:
    <pre><code>python -m venv &lt;venv_name&gt;</code></pre>
  </li>
  <li>Activate the virtual environment.</li>
  <li>Install dependencies:
    <pre><code>pip install requests</code></pre>
  </li>
  <li>Run the tests:
    <pre><code>python test_the_api.py</code></pre>
  </li>
</ol>
