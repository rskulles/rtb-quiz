import './style.css'

let questionTime = 10000;//10000 ms is 10 seconds
const questionInterval = 1000;
let questionId = "???";
let submitButton = document.querySelector("#submit-answer");
submitButton.addEventListener('click', submitAnswer);

document.addEventListener('DOMContentLoaded', () => {
  //TODO: This jsut loads the static content that is already on the page.
  //This should do the following in the future.
  //Probably just call showQuestion()
  //Show question should look for a query string to the question
  //Make a server request to get the question.
  //Fill out the question and aswers from the response.
  //Hide loading and show the question.
  //use a setTimeout for timing answering the question.
  setTimeout(showQuestion, 3000);
});

function tickTime() {
  //subtract the interval from the time remaining
  questionTime -= questionInterval;
  //TODO: if there is still time left, call the interval again.
  //else move on to the next question.

}

//This should be called at some point in the 'DOMContentLoaded' listener above
function setQuestionId(id) {
  questionId = id;
}

function showQuestion() {
  let loadingSection = document.querySelector("#loading-section");
  let questionSection = document.querySelector("#question-section");
  //TODO: will need to get the quesion/answer elements to fill them in
  loadingSection.style.display = "none";
  questionSection.style.display = "initial";

}
function submitAnswer() {
  alert("You are trying to submit an answer!");
  //TODO: How submitting an answer should work
  //Get all the radio inputs
  //Find the one that is selected
  //Submit the answer to the server useing the fetch api
  //The server should point you to the next question in the response
  //Redirect to the correct page
  //
}
