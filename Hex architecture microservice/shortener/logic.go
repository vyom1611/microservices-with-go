package shortener

import (
	"errors"
	errs "github.com/pkg/errors"
	"github.com/teris-io/shortid"
	"gopkg.in/dealancer/validate.v2"
	"time"
)

var (
	ErrRedirectNotFound = errors.New("Redirect not found")
	ErrRedirectInvalid  = errors.New("Redirect invalid")
)

type redirectService struct {
	redirectRepo RedirectRepository
}

func (r redirectService) Find(code string) (*Redirect, error) {
	//Finding redirect from the code
	return r.redirectRepo.Find(code)
}

func (r redirectService) Store(redirect *Redirect) error {
	//Validating URL of the redirect
	//The validate package uses a bunch of stuff like regular expressions to validate
	if err := validate.Validate(redirect); err != nil {
		return errs.Wrap(ErrRedirectInvalid, "service.Redirect.Store")
	}
	//Creating IDs with shortid
	redirect.Code = shortid.MustGenerate()
	//Timestamp
	redirect.CreatedAt = time.Now().UTC().Unix()

	//Calling to the redirect repo
	return r.redirectRepo.Store(redirect)
}

func NewRedirectService(redirectRepo RedirectRepository) RedirectService {
	return &redirectService{
		redirectRepo,
	}

}
