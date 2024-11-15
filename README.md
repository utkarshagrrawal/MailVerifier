<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
</head>
<body>

  <h1>Mail Verifier</h1>
  <p>A comprehensive Go-based email verification tool that checks if an email address is valid and deliverable. The tool verifies emails by checking MX records, filtering disposable domains, validating format with regular expressions, and testing deliverability via SMTP.</p>
  <blockquote>Note: Some SMTP servers may restrict deliverability checks, which can affect verification results.</blockquote>

  <h2>Features</h2>
  <ul>
    <li><strong>Email Syntax Validation:</strong> Ensures the email address follows proper syntax using regex.</li>
    <li><strong>MX Record Check:</strong> Verifies that the email’s domain has valid MX records, indicating an email server is available.</li>
    <li><strong>Disposable Domain Check:</strong> Filters out email addresses from known disposable domains.</li>
    <li><strong>Deliverability Check via SMTP:</strong> Attempts to connect to the SMTP server to verify if the email address is deliverable.</li>
  </ul>

  <h2>Requirements</h2>
  <ul>
    <li>Go 1.23.3 or compatible version.</li>
    <li>Internet connection to reach DNS servers and SMTP servers.</li>
  </ul>

  <h2>Getting Started</h2>

  <h3>Installation</h3>
<ol>
  <li>Clone this repository:
    <div class="code-block">
      <pre>git clone https://github.com/utkarshagrrawal/mailverifier.git</pre>
    </div>
  </li>
  <li>Navigate to the backend directory and build the Go application:
    <div class="code-block">
      <pre>
cd backend
go build .
./mailverifier</pre>
    </div>
  </li>
  <li>Navigate to the frontend directory and start the development server:
    <div class="code-block">
      <pre>
cd ../frontend
npm run dev</pre>
    </div>
  </li>
</ol>

  <h2>Limitations</h2>
  <p>Some limitations to consider:</p>
  <ul>
    <li>SMTP deliverability checks may fail for servers that disable the <code>VRFY</code> command.</li>
  </ul>

  <h2>Author</h2>
  <p>This project was created by <a href="https://github.com/utkarshagrrawal">me</a>.</p>

  <h2>License</h2>
  <p>This project is licensed under the MIT License - see the <a href="LICENSE">LICENSE</a> file for details.</p>

</body>
</html>
