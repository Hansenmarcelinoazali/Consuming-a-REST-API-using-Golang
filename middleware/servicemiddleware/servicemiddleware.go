package servicemiddleware

import (
	"errors"
	"external_api/middleware/repositorymiddleware"
)

func ServiceMiddlewareIsLogin(refreshsToken string) (string, error) {

	result, _ := repositorymiddleware.RepoMiddleware(refreshsToken)
	// if err != nil {
	// 	return "", nil
	// }
	if result.IsLogin == true {
		return "sudah login", nil
	} else {
		return " ", errors.New("belum login")
	}

}
