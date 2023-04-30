package project

type Api struct {
	Name        string `json:"name"`
	NameEn      string `json:"nameEn"`
	Method      string `json:"method"`
	Uri         string `json:"uri"`
	Description string `json:"description"`
}
type Functions struct {
	Name   string `json:"name"`
	NameEn string `json:"nameEn"`
	Apis   []*Api `json:"apis"`
}
type Field struct {
	Name string
	Type string
}
type Model struct {
	Name   string   `json:"name"`
	NameEn string   `json:"nameEn"`
	Fields []*Field `json:"fields"`
}

type Project struct {
	NameEn        string       `json:"name_en" comment:"项目英文名称"`
	Name          string       `json:"name" comment:"项目名称"`
	FunctionsList []*Functions `json:"functions" comment:"功能列表"`
	Models        []*Model     `json:"models" comment:"数据库表"`
}
