function loginToAccount(loginName, passwordName) {
  if (loginName == "bizon" && passwordName == "dupa") {
    return true;
  }
  return false;
}

module.exports = loginToAccount;
