const electron = require('electron');
const url = require('url');
const path = require('path');

const {app, BrowserWindow, Menu} = electron;

let mainWindow;



app.on('ready', function(){
  mainWindow = new BrowserWindow({});

  mainWindow.loadURL(url.format({
    pathname: path.join(__dirname + "/window", 'mainWindow.html'),
    protocol: 'file:',
    slashes: true
  }));

  const mainMenuTemplate =  [
    {
        label: 'File',
        submenu: [
            {
                label: 'Quit',
                click() {
                    app.quit();
                }
            }
        ]
    }
];

if (process.platform == 'darwin') {
    mainMenuTemplate.unshift({});
}

  const mainMenu = Menu.buildFromTemplate(mainMenuTemplate);

  Menu.setApplicationMenu(mainMenu);
});
