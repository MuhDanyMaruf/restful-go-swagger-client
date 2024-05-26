import { login, getAllBajus } from "./swaggerfunction.js";

window.login = async function () {
  const username = document.getElementById("username").value;
  const password = document.getElementById("password").value;

  try {
    const token = await login(username, password);
    document.getElementById("result").innerText =
      "Login successful. Token: " + token;
    const bajus = await getAllBajus(token);
    document.getElementById("result").innerText +=
      "\n\nBajus:\n" + JSON.stringify(bajus, null, 2);
  } catch (error) {
    document.getElementById("result").innerText = "Error: " + error.message;
  }
};
