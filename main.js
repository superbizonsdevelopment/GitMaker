const electron = require('electron');
const url = require('url');
const path = require('path');

const {app, BrowserWindow, Menu} = electron;

let mainWindow;



app.on('ready', function(){
  mainWindow = new BrowserWindow({});

  mainWindow.loadURL(url.format({
    pathname: path.join(__dirname + "/window", 'loginWindow.html'),
    protocol: 'file:',
    slashes: true
  }));
  const mainMenu = Menu.buildFromTemplate(mainMenuTemplate);

  Menu.setApplicationMenu(mainMenu);
});

//Create menu template
const mainMenuTemplate = [
  {
    label:'File'
  }
];
