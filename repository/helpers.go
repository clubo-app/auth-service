package repository

import (
	"fmt"
	"net/url"

	ag "github.com/clubo-app/protobuf/auth"
	"github.com/golang-migrate/migrate/v4"
	g "github.com/golang-migrate/migrate/v4/source/github"
)

func (p Provider) ToGRPCProvider() ag.Provider {
	switch p {
	case ProviderAPPLE:
		return ag.Provider_APPLE
	case ProviderFACEBOOK:
		return ag.Provider_FACEBOOK
	case ProviderGOOGLE:
		return ag.Provider_GOOGLE
	default:
		// TODO: make this return null
		return 4
	}
}

func (t Type) ToGRPCAccountType() ag.Type {
	switch t {
	case TypeCOMPANY:
		return ag.Type_COMPANY
	case TypeUSER:
		return ag.Type_USER
	case TypeADMIN:
		return ag.Type_ADMIN
	case TypeDEV:
		return ag.Type_DEV
	default:
		return ag.Type_USER
	}
}

func (a Account) ToGRPCAccount() *ag.Account {
	return &ag.Account{
		Id:            a.ID,
		Email:         a.Email,
		EmailVerified: a.EmailVerified,
		EmailCode:     a.EmailCode.String,
		Provider:      a.Provider.ToGRPCProvider(),
		Type:          a.Type.ToGRPCAccountType(),
	}
}

const version = 1

func validateSchema(url url.URL) error {
	url.Scheme = "pgx"
	url2 := fmt.Sprintf("%v%v", url.String(), "?sslmode=disable")
	g := g.Github{}
	d, err := g.Open("github://clubo-app/auth-service/repository/migrations")
	if err != nil {
		return err
	}
	defer d.Close()

	m, err := migrate.NewWithSourceInstance("github", d, url2)

	if err != nil {
		return err
	}
	err = m.Migrate(version) // current version
	if err != nil && err != migrate.ErrNoChange {
		return err
	}
	defer m.Close()
	return nil
}
