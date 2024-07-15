// Package factories 生成 user 数据
package factories

import (
	"blog/internal/server/models/user"
	"blog/pkg/helps"
	"github.com/bxcodec/faker/v3"
)

func MakeUsers(time int) []user.User {

	var objs []user.User

	// 设定唯一值
	faker.SetGenerateUniqueValues(true)

	for i := 0; i < time; i++ {
		model := user.User{
			LoginName: faker.Username(),
			TrueName:  faker.Name(),
			PassWord:  "$2a$14$oPzVkIdwJ8KqY0erYAYQxOuAAlbI/sFIsH0C0R4MPc.3JbWWSuaUe",
			Phone:     helps.RandomNumber(11),
		}

		objs = append(objs, model)
	}
	return objs
}
