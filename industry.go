package random

// Industry 行业类别
type Industry struct {
	Level1 IndustryElement // 门类
	Level2 IndustryElement // 大类
	Level3 IndustryElement // 中类
	Level4 IndustryElement // 小类
}

// IndustryElement 行业类别基础分类
type IndustryElement struct {
	Code string
	Name string
}

type industryClassNode struct {
	Code     string
	Name     string
	children []*industryClassNode
}

type industryRoot struct { // 固定4层
	Category []*industryClassNode // 门类 大 中 小
}

var industryTree *industryRoot

func init() {
	var inds []*industryClassNode

	registers[DataIndustryChina] = func() {
		CheckAndDecompress(dconf.dataUrls["industry"], &inds)
		industryTree = createFrom(inds)
	}
}

func (inds *industryClassNode) String() string {
	return inds.Code + " " + inds.Name
}

func createFrom(data []*industryClassNode) *industryRoot { // 固定4层
	industryTree = &industryRoot{}

	var c1, c2, c3 *industryClassNode
	for _, c := range data {
		// log.Println(c.Code, c.Name)
		if len(c.Code) == 1 {
			industryTree.Category = append(industryTree.Category, c)
			c1 = c
			continue
		}

		if len(c.Code) == 2 {
			c1.children = append(c1.children, c)
			c2 = c
			continue
		}

		if len(c.Code) == 3 {
			c2.children = append(c2.children, c)
			c3 = c
			continue
		}

		if len(c.Code) == 4 {
			c3.children = append(c3.children, c)
			continue
		}
	}

	return industryTree
}

func (inds *industryRoot) get(selects ...int) *industryClassNode {
	var sel *industryClassNode
	cur := inds.Category
	for _, s := range selects {
		sel = cur[s]
		cur = cur[s].children
	}
	return sel
}

func (ext *Extend) Industry() *Industry {

	var ind *Industry = &Industry{}
	cur := industryTree.Category

	sel := cur[ext.Intn(len(cur))]
	ind.Level1.Code = sel.Code
	ind.Level1.Name = sel.Name
	cur = sel.children

	sel = cur[ext.Intn(len(cur))]
	ind.Level2.Code = sel.Code
	ind.Level2.Name = sel.Name
	cur = sel.children

	sel = cur[ext.Intn(len(cur))]
	ind.Level3.Code = sel.Code
	ind.Level3.Name = sel.Name
	cur = sel.children

	sel = cur[ext.Intn(len(cur))]
	ind.Level4.Code = sel.Code
	ind.Level4.Name = sel.Name

	return ind
}
