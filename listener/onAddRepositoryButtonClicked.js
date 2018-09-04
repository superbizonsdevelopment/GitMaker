document.getElementById("addRepositoryButton").addEventListener("click", function() {
  const load = require('./main.js');
  load.loadWindow('createFieldForRepositoryWindow');
});
