import './style.css'

let button = document.querySelector("#create-user")
button.addEventListener('click', () => {
  const name = document.querySelector("#user-name")
  const password = document.querySelector("#user-name")
  const verify = document.querySelector("#user-name")

  let data = { name: name.value, password: password.value, verify: verify.value }
  fetch("http://localhost:3000/users", {
    method: "POST",
    body: JSON.stringify(data),
    headers: {
      "Content-type": "application/json; charset=UTF-8"
    }
  }).then((resp) => {
    if (resp.ok) {
      name.value = ""
      password.value = ""
      verify.verify = ""
    }
    return resp.json()
  })
    .then((json) => console.log(json));
})
