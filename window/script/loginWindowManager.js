var loginButton = document.getElementById('loginButton');

loginButton.addEventListener('click', function(){
  var loginField = document.getElementById('loginField').value;

  if (loginField == "bizon") {
    console.log("done");
  }
})
