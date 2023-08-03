package parser

import "github.com/finexblock-dev/gofinexblock/pkg/entity"

func AdminToPartial(data []*entity.Admin) (result []*entity.PartialAdmin) {
	for _, v := range data {
		result = append(result, &entity.PartialAdmin{
			ID:           v.ID,
			Email:        v.Email,
			Grade:        v.Grade,
			IsBlocked:    v.IsBlocked,
			InitialLogin: v.InitialLogin,
			CreatedAt:    v.CreatedAt,
			UpdatedAt:    v.UpdatedAt,
			PwdUpdatedAt: v.PwdUpdatedAt,
		})
	}
	return result
}