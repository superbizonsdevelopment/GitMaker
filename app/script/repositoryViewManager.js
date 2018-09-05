function showSavedRepositories(elmnt, repositories) {

    for (var i = 0; i < repositories.size(); i++) {
      const repository = elmnt.createElement('div');
      repository.classList.add('klasa');
    }
}

module.exports = {
  showSavedRepositories
}
