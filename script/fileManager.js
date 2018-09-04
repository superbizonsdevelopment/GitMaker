const fs = require('fs');
//const files = require('files.js');

function createAppDirectory() {
  try {
      fs.mkdirSync('gitmaker');
      console.log("GitMaker directory created!");
  } catch (err) {
    if (err.code == 'EEXIST') {
      console.log("The directory gitmaker exists.");
    }
  }
}

function createRepositoriesFile() {

  if (!fs.existsSync('gitmaker')) { //files.__AppDirectory
    createAppDirectory();
  }

  fs.appendFile('repositories.json', '{repositories: {}}', function (err) {//files.__RepositoryFile
  if (err) throw err;
    console.log('Saved!');
  });
}

function getRepositoriesFromFile() {
  fs.readFile('repositories.json', function(err, data) {
    console.log('Content: ');
    console.log(String.fromCharCode(data));
  });
}

module.exports = {
  createAppDirectory,
  createRepositoriesFile,
  getRepositoriesFromFile
}
