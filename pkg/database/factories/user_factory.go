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
			TrueName:  faker.ChineseName(),
			PassWord:  faker.Password(),
			Phone:     helps.RandomNumber(11),
			Address:   faker.Sentence(),
		}

		objs = append(objs, model)
	}
	return objs
}
