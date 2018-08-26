const electron = require('electron');
const url = require('url');
const path = require('path');
var fileManager = require(__dirname + "/script/fileManager.js");
var repositoryManager = require(__dirname + "/script/repositoryManager.js");

const {app, BrowserWindow, Menu} = electron;

let createRepositoryWindow;

app.on('ready', function() {

  let mainWindow = new BrowserWindow({});

  mainWindow.loadURL(url.format({
    pathname: path.join(__dirname + "/window", 'mainWindow.html'),
    protocol: 'file:',
    slashes: true
  }));

  const mainMenu = Menu.buildFromTemplate(mainMenuTemplate);
  Menu.setApplicationMenu(mainMenu);

  //repositoryManager.isRepositoryExists('exdns');
  //fileManager.createAppDirectory();
  //fileManager.createRepositoriesFile();
});

app.on('window-all-closed', () => {
    // On macOS it is common for applications and their menu bar
    // to stay active until the user quits explicitly with Cmd + Q
    if (process.platform !== 'darwin') {
      app.quit()
    }
})

function loadCreateRepositoryWindow() {
  createRepositoryWindow = new BrowserWindow({
    width: 400,
    height: 400
  });

  createRepositoryWindow.loadURL(url.format({
    pathname: path.join(__dirname + "/window", 'createFieldForRepositoryWindow.html'),
    protocol: 'file:',
    slashes: true
  }));
}

const mainMenuTemplate =  [
  {
    label: 'File',
    submenu: [
      {
        label: 'Quit',
        accelerator: 'CmdOrCtrl+Q',
        click() { app.quit(); }
      }
    ]
  },
  {
    label: 'Repository',
    submenu: [
      {
        label: 'Create',
        click() { loadCreateRepositoryWindow(); }
      },
      {
        label: 'Delete',
        click() { }
      }
    ]
  }
];
