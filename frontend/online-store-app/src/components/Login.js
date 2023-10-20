import React, { useState } from 'react';

const Login = ({ onLogin }) => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');

  const handleSubmit = event => {
    event.preventDefault();

    // Prepare the request body
    const data = {
      username: username,
      password: password,
    };

    // Make a POST request to the login endpoint
    fetch('http://localhost:8080/v1/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
    })
      .then(response => {
        if (response.status === 200) {
          // Handle successful login, e.g., store the token in localStorage
          console.log('Login successful!');
          return response.json();
        } else if (response.status === 401) {
          // Handle unauthorized access
          console.error('Unauthorized access');
        } else {
          // Handle other errors
          console.error('Login failed');
        }
      })
      .then(data => {
        // Handle the response data if needed
        console.log(data);

        const token = data.token;
        onLogin(token)

        // Store the token in local storage
        localStorage.setItem('token', token);
      })
      .catch(error => {
        // Handle errors
        console.error('Error:', error);
      });

      onLogin();
  };

  return (
    <div className="login-form">
      <h2>Login</h2>
      <form onSubmit={handleSubmit}>
        <div className="form-group">
          <label>Username:</label>
          <input
            type="text"
            value={username}
            onChange={e => setUsername(e.target.value)}
            required
          />
        </div>
        <div className="form-group">
          <label>Password:</label>
          <input
            type="password"
            value={password}
            onChange={e => setPassword(e.target.value)}
            required
          />
        </div>
        <button type="submit">Login</button>
      </form>
    </div>
  );
};

export default Login;
