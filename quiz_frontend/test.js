import './style.css'

var button = document.querySelector('#change-heading');
button.addEventListener('click', changeText);

function changeText() {
  let heading = document.querySelector("#change-target")
  heading.innerHTML = ("Hello, Mom!!!")
}
