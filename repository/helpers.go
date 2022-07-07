package repository

import (
	"embed"
	"fmt"
	"net/url"

	ag "github.com/clubo-app/protobuf/auth"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

func (p NullProvider) ToGRPCProvider() ag.Provider {
	if !p.Valid {
		return ag.Provider_UNKOWNPROVIDER
	}

	val, err := p.Value()
	if err != nil {
		return ag.Provider_UNKOWNPROVIDER
	}

	switch val {
	case ProviderAPPLE:
		return ag.Provider_APPLE
	case ProviderFACEBOOK:
		return ag.Provider_FACEBOOK
	case ProviderGOOGLE:
		return ag.Provider_GOOGLE
	default:
		return ag.Provider_UNKOWNPROVIDER
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
		return ag.Type_UNKOWNTYPE
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

//go:embed migrations/*.sql
var fs embed.FS

func migrateSchema(url url.URL) error {
	url.Scheme = "pgx"
	urlf := fmt.Sprintf("%v%v", url.String(), "?sslmode=disable")

	d, err := iofs.New(fs, "migrations")
	if err != nil {
		return err
	}

	m, err := migrate.NewWithSourceInstance("github", d, urlf)
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
