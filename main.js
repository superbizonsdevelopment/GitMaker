const electron = require('electron');
const url = require('url');
const path = require('path');
//var loginScripts = require("script/login.js");
var loginButton = document.getElementById('loginButton');

const {app, BrowserWindow, Menu} = electron;

let mainWindow = new BrowserWindow({});


app.on('ready', function(){
  mainWindow.loadURL(url.format({
    pathname: path.join(__dirname + "/window", 'loginWindow.html'),
    protocol: 'file:',
    slashes: true
  }));
});

loginButton.addEventListener('click', function(){
  var loginField = document.getElementById('loginField').value;

  if (loginField == "bizon") {
    mainWindow.loadURL(url.format({
      pathname: path.join(__dirname + "/window", 'mainWindow.html'),
      protocol: 'file:',
      slashes: true
    }));
  }
})
