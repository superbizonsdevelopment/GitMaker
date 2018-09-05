var GitHub = require('github-api');

const github = new GitHub();

function isRepositoryExists(nameOfRepository) {

  if (typeof nameOfRepository != "string") {
    console.error("Wrong format!!!");
    return;
  }

  if (github.search().forRepositories(nameOfRepository) != null) {
    return true;
  }
  return false;
}

module.exports = {
  isRepositoryExists
}
