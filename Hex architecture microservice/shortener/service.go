package shortener

type RedirectService interface {
	//Find method will find the appropriate url and send back to the user
	Find(code string) (*Redirect, error)
	//Store method will store the url in the repository
	Store(redirect *Redirect) error
}
