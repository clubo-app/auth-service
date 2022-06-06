package repository

import (
	"fmt"
	"net/url"

	ag "github.com/clubo-app/protobuf/auth"
	"github.com/clubo-app/protobuf/common"
	"github.com/golang-migrate/migrate/v4"
	g "github.com/golang-migrate/migrate/v4/source/github"
)

func (p Provider) ToGRPCProvider() common.Provider {
	switch p {
	case ProviderApple:
		return common.Provider_APPLE
	case ProviderFacebook:
		return common.Provider_FACEBOOK
	case ProviderGoogle:
		return common.Provider_GOOGLE
	default:
		// TODO: make this return null
		return 4
	}
}

func (r Role) ToGRPCRole() common.Role {
	switch r {
	case RoleAdmin:
		return common.Role_ADMIN
	case RoleDev:
		return common.Role_DEV
	case RoleUser:
		return common.Role_USER
	default:
		return common.Role_USER
	}
}

func (t Type) ToGRPCAccountType() ag.AccountType {
	switch t {
	case TypeCompany:
		return ag.AccountType_COMPANY
	case TypeUser:
		return ag.AccountType_USER
	default:
		return ag.AccountType_USER
	}
}

func (a Account) ToGRPCAccount() *ag.Account {
	return &ag.Account{
		Id:            a.ID,
		Email:         a.Email,
		EmailVerified: a.EmailVerified,
		EmailCode:     a.EmailCode.String,
		Provider:      a.Provider.ToGRPCProvider(),
		Role:          a.Role.ToGRPCRole(),
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
