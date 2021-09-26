package random

import (
	"log"
	"testing"
)

func TestName(t *testing.T) {
	r := New()

	Use(DataNameChina)
	Use(DataIndustryChina)
	Use(DataCityChina)
	Use(DataPoetryChina)

	log.Println(r.Extend().FirstName()) // 宇文
	log.Println(r.Extend().LastName())  // 毅皿
	log.Println(r.Extend().FullName())  // 费梦骞

	ind := r.Extend().Industry()
	log.Println(ind.Level1.Code, ind.Level1.Name) // 门类: B 采矿业
	log.Println(ind.Level2.Code, ind.Level2.Name) // 大类: 07 石油和天然气开采业
	log.Println(ind.Level3.Code, ind.Level3.Name) // 中类: 072 天然气开采
	log.Println(ind.Level4.Code, ind.Level4.Name) // 小类: 0722 海洋天然气及可燃冰开采

	city := r.Extend().City()
	log.Println(city) // 泰州市

	poetry := r.Extend().Poetry()
	log.Println(poetry.Title, poetry.Writer, poetry.Dynasty, poetry.Content)
	// 标题: 读山海经·其一
	// 作者: 陶渊明
	// 朝代: 魏晋
	// 诗内容:
	// 孟夏草木长，绕屋树扶疏。
	// 众鸟欣有托，吾亦爱吾庐。
	// 既耕亦已种，时还读我书。
	// 穷巷隔深辙，颇回故人车。
	// 欢言酌春酒，摘我园中蔬。
	// 微雨从东来，好风与之俱。
	// 泛览《周王传》，流观《山海》图。
	// 俯仰终宇宙，不乐复何如
}
