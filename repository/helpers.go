package repository

import (
	ag "github.com/clubo-app/protobuf/auth"
	"github.com/clubo-app/protobuf/common"
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
