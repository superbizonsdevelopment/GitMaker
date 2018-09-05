const electron = require('electron');
const url = require('url');
const path = require('path');
const fileManager = require(__dirname + "/script/fileManager.js");
const repositoryManager = require(__dirname + "/script/repositoryManager.js");
const repositoryViewManager = require(__dirname + "/script/repositoryViewManager.js");

const {app, BrowserWindow, Menu} = electron;

let windowToLoad;

app.on('ready', function() {

  let mainWindow = new BrowserWindow({});

  mainWindow.loadURL(url.format({
    pathname: path.join(__dirname + "/window", 'mainWindow.html'),
    protocol: 'file:',
    slashes: true
  }));

  const mainMenu = Menu.buildFromTemplate(mainMenuTemplate);
  Menu.setApplicationMenu(mainMenu);

  fileManager.createAppDirectory();
  fileManager.createRepositoriesFile();

  fileManager.getRepositoriesFromFile();

  var repositoriesInArray = new Array('januszowe', 'zajebiste');

  repositoryViewManager.showSavedRepositories(repositoriesInArray);
});

app.on('window-all-closed', () => {
    // On macOS it is common for applications and their menu bar
    // to stay active until the user quits explicitly with Cmd + Q
    if (process.platform !== 'darwin') {
      app.quit()
    }
})

function loadWindow(name) {
  windowToLoad = new BrowserWindow({
    width: 400,
    height: 400
  });

  windowToLoad.loadURL(url.format({
    pathname: path.join(__dirname + "/window", name + '.html'),
    protocol: 'file:',
    slashes: true
  }));
}

const mainMenuTemplate =  [
  {
    label: 'GitMaker',
    submenu: [
      {
        label: 'About GitMaker',
        click() { loadWindow('aboutWindow') }
      },
      {
        label: 'View License',
        click() { loadWindow('licenseWindow') }
      },
      {
        label: 'Version 1.0.0'
      },
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
        click() { loadWindow('createFieldForRepositoryWindow'); }
      },
      {
        label: 'Delete',
        click() { }
      }
    ]
  }
];

module.exports = loadWindow()
