var fs = require('fs');

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
  try {
    
  } catch (err) {

  }
}

module.exports = {
  createAppDirectory,
  createRepositoriesFile
}
