package repository

var(
	Repositories = make(map[string]Repository)
)

type Repository struct{
	Name string
	Author string
	URL string
}

func makeNewRepository(name string, author string, url string) {
	repo := Repository{name, author, url}
	
	Repositories[repo.Name] = repo
}

func removeRepository(repository Repository) {
	delete(repository.Name, repository)
}