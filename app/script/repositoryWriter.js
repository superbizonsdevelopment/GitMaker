const fm = require('fileManager.js')
const files = require('files.js')

function writeRepository(name, author, url, stars, isStarred) {
  if (isRepositoryExists(name)) {
    console.warn("Repository: " + name + " exists! Program Can't make file with same name!");
    return;
  }

  fm.open(files.__RepositoryFile, 'w', function (err, file) {
    if (err) throw err;
    console.log(file);
  });
}

function getRepository(name) {

}

function getAllRepositories() {
  var repository = new Object();
  var repositories = new Array();

  

  return repositories;
}


function isRepositoryExists(name) {
  return false;
}

function deleteRepository(name) {

}

module.exports = {
  writeRepository,
  getRepository,
  getAllRepositories,
  isRepositoryExists,
  deleteRepository
}
