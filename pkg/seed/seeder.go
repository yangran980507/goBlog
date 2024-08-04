// Package seed 数据库填充
package seed

import (
	"blog/pkg/console"
	"blog/pkg/mysql"
	"gorm.io/gorm"
)

var seeders []Seeder

var orderedSeederNames []string

type SeederFunc func(db *gorm.DB)

type Seeder struct {
	Func SeederFunc
	Name string
}

// Add 注册到 seeders
func Add(name string, fn SeederFunc) {
	seeders = append(seeders, Seeder{
		Name: name,
		Func: fn,
	})
}

func SetRunOrder(names []string) {
	orderedSeederNames = names
}

func GetSeeder(name string) Seeder {

	for _, sdr := range seeders {
		if name == sdr.Name {
			return sdr
		}
	}
	return Seeder{}
}

// RunAll 运行所有 Seeder
func RunAll() {

	// 先运行 ordered 的
	executed := make(map[string]string)
	for _, name := range orderedSeederNames {
		sdr := GetSeeder(name)
		if len(sdr.Name) > 0 {
			console.Warn("Running Ordered Seeder: " + sdr.Name)
			sdr.Func(mysql.DB)
			executed[name] = name

		}
	}

	// 再运行剩下的
	for _, sdr := range seeders {
		// 过滤已运行
		if _, ok := executed[sdr.Name]; !ok {
			console.Warn("Running Seeder: " + sdr.Name)
			sdr.Func(mysql.DB)
		}
	}
}

// RunSeeder 运行单个 Seeder
func RunSeeder(name string) {
	for _, sdr := range seeders {
		if name == sdr.Name {
			sdr.Func(mysql.DB)
			break
		}
	}
}
