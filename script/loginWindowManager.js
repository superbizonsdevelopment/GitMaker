const { remote, BrowserWindow } = require('electron')
const currentWindow = remote.getCurrentWindow()

var loginButton = document.getElementById('loginButton');

loginButton.addEventListener('click', function(){
  var login = document.getElementById('loginField').value;
  var password = document.getElementById('passwordField').value;

  if (login == "bizon" && password == "dupa") {
    console.log("done");
    currentWindow.loadURL(url.format({
      pathname: path.join(__dirname + "/window", 'mainWindow.html'),
      protocol: 'file:',
      slashes: true
    }));
  }

})
