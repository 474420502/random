# 中国相关的随机数据. 用于测试, 模拟数据.


## 数据分为内置数据和需要加载的数据. 具体用法如下:



* 添加依赖包, 不包含数据

``` bash
go get  github.com/474420502/random
```

* 例子

``` golang
package main

import (
	"log"

	"github.com/474420502/random"
)

func main() {
	r := random.New() // 可以定义种子. 方便测试查找问题. 默认继承go rand包. 加了几个函数.

 	log.Println(r.Probability(0.4)) // 40%的概率返回true
	log.Println(r.OneOf64n(4))      // 4份一概率返回true

	random.Use(DataNameChina)
	random.Use(DataIndustryChina)
	random.Use(DataCityChina)
	random.Use(DataPoetryChina)
	random.Use(DataEmailChina)
	random.Use(DataCountryChina)

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

	phone := r.Extend().Phone()
	log.Println(phone.Operators, phone.Number) // 中国电信 18030082158

	date1 := r.Extend().Date()
	date2 := r.Extend().DateRange("2021-10-21", "2021-10-25")
	log.Println(date1, date2)
	// 1993-11-22 16:55:37.264646 +0800 CST
	// 2021-10-22 19:16:17.424177 +0800 CST

	// 随机文字
	log.Println(r.Extend().TextN(10)) // 漡詟覷髋簑㖦䈉们䩈㥄
	// email
	log.Println(r.Extend().Email()) // vkwlfjbdd7@foot.com
	// 国家
	country := r.Extend().Country()
	log.Println(country.Code, country.Alpha2, country.Alpha3, country.FullName, country.LocalName) // vkwlfjbdd7@foot.com
	// 116 KH KHM the Kingdom of Cambodia 柬埔寨

	log.Println(r.Bytes(10, 100)) // 随机字节

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

```

* 需要还原上次测试

```
从日志输出 seed: 1632645723358105735 提取 1632645723358105735 种子
```

```go
r := random.New(1632645723358105735)
// 这样就可以还原上次出错的随机测试情景
```

* 所有需要下载的内容

```
	// DataNameChina 姓名数据
	DataNameChina
	// DataProvinceChina 省份
	DataProvinceChina
	// DataCityChina 城市
	DataCityChina
	// DataTownChina 县城
	DataAreaChina
	// DataTownStreetChina 街道(乡) 数据 Town() Street() 一样的数据
	DataTownStreetChina
	// DataCountryChina 全球国家 数据
	DataCountryChina
	// DataEmailChina email 数据
	DataEmailChina
	// DataIndustryChina 中国行业 数据
	DataIndustryChina
	// DataIdidomChina 中国成语
	DataIdidomChina
	// DataIdidomChina 中国诗文
	DataPoetryChina
```