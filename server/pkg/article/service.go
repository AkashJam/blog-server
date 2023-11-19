package article

//Necessary to import local package into another local package
type Service interface {
	GetArticles() ([]article, error)
}

// type ListRepository interface {
// 	GetArticles() []article
// }

type service struct {
	// supposed to be repository when using db
	s Service
}

// takes repository as input which is the db initialization
func NewService() Service {
	return &service{}
}