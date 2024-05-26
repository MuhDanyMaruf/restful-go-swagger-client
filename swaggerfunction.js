export async function login(username, password) {
  const response = await fetch("http://localhost:8081/login", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ username, password }),
  });

  if (response.ok) {
    const data = await response.json();
    return data.token;
  } else {
    throw new Error("Login failed. Status: " + response.status);
  }
}

export async function getAllBajus(token) {
  const response = await fetch("http://localhost:8081/baju", {
    method: "GET",
    headers: {
      Authorization: "Bearer " + token,
    },
  });

  if (response.ok) {
    return await response.json();
  } else {
    throw new Error("Failed to get bajus. Status: " + response.status);
  }
}
