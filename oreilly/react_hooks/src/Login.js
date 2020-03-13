import React, { useState } from 'react' ;

export default function Login() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [user, setUser] = useState(null);

  const handleSubmit = event => {
    event.preventDefault();
    const userData = {
      username,
      password
    };
    setUser(userData);
    setUsername("");
    setPassword("");
  };

  return (
    <div
      style={{
        textAlign: "center"
      }}
    >
      <h2>Login</h2>
      <form
        style={{
          display: 'grid',
          alignItems: 'center',
          justifyItems: 'center'
        }}
        onSubmit={handleSubmit}
      >
        <input
          type="text"
          placeholder="Username"
          onChange={event => setUsername(event.target.value)}
          value={username}
        />
        <input
          type="password"
          placeholder="Password"
          onChange={event => setPassword(event.target.value)}
          value={password}
        />
        <button type="submit">Submit</button>
      </form>

      {user && JSON.stringify(user, null, 2)}
    </div>
  );
}
