const BASE_URL = "http://localhost:8080";

let token = "";

// REGISTER
async function register() {
  const res = await fetch(BASE_URL + "/register", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      username: regUsername.value,
      email: regEmail.value,
      password: regPassword.value
    })
  });

  document.getElementById("status").innerText = "Account created ✅";
}

// LOGIN
async function login() {
  const res = await fetch(BASE_URL + "/login", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      email: loginEmail.value,
      password: loginPassword.value
    })
  });

  const data = await res.json();

  token = data.token;

  document.getElementById("status").innerText = "Logged in 🔥";
}

// UPLOAD
async function uploadPhoto() {
  const file = fileInput.files[0];

  if (!file) {
    alert("Choose file!");
    return;
  }

  let formData = new FormData();
  formData.append("file", file);

  await fetch(BASE_URL + "/photos/upload", {
    method: "POST",
    headers: {
      "Authorization": "Bearer " + token
    },
    body: formData
  });

  loadPhotos();
}

// LOAD PHOTOS
async function loadPhotos() {
  const res = await fetch(BASE_URL + "/photos");
  const data = await res.json();

  photos.innerHTML = "";

  data.forEach(p => {
    photos.innerHTML += `
      <img src="${BASE_URL}/${p.url}" />
    `;
  });
}