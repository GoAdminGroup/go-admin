package auth

import (
	"github.com/chenhg5/go-admin/modules/connections"
	"github.com/chenhg5/go-admin/plugins/admin/modules"
	"strconv"
	"github.com/chenhg5/go-admin/context"
	"golang.org/x/crypto/bcrypt"
)

func Check(password string, username string) (user User, ok bool) {

	admin, _ := connections.GetConnection().Query("select * from goadmin_users where username = ?", username)

	if len(admin) < 1 {
		ok = false
	} else {
		if ComparePassword(password, admin[0]["password"].(string)) {
			ok = true

			roleModel, _ := connections.GetConnection().Query("select r.id, r.name, r.slug from goadmin_role_users " +
				"as u left join goadmin_roles as r on u.role_id = r.id where user_id = ?", admin[0]["id"])

			user.ID = strconv.FormatInt(admin[0]["id"].(int64), 10)
			user.Level = roleModel[0]["slug"].(string)
			user.LevelName = roleModel[0]["name"].(string)
			user.Name = admin[0]["name"].(string)
			user.CreateAt = admin[0]["created_at"].(string)
			user.Avatar = admin[0]["avatar"].(string)

			newPwd := EncodePassword([]byte(password))
			connections.GetConnection().Exec("update goadmin_users set password = ? where id = ?", newPwd, user.ID)

		} else {
			ok = false
		}
	}
	return
}

func ComparePassword(comPwd, pwdHash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(pwdHash), []byte(comPwd))
	if err != nil {
		return false
	} else {
		return true
	}
}

func EncodePassword(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hash[:])
}

func SetCookie(ctx *context.Context, user User) bool {
	InitSession(ctx).Set("user_id", user.ID)
	return true
}

func DelCookie(ctx *context.Context) bool {
	InitSession(ctx).Clear()
	return true
}

type CSRFToken []string

var TokenHelper = new(CSRFToken)

func (token *CSRFToken) AddToken() string {
	tokenStr := modules.Uuid(35)
	if len(*token) == 1 && (*token)[0] == "" {
		(*token)[0] = tokenStr
	} else {
		*token = append(*token, tokenStr)
	}
	return tokenStr
}

func (token *CSRFToken) CheckToken(tocheck string) bool {
	for i := 0; i < len(*token); i++ {
		if (*token)[i] == tocheck {
			*token = append((*token)[0:i], (*token)[i:len(*token)]...)
			return true
		}
	}
	return false
}
