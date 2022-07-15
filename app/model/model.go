package model

type TemplateTest struct {
	Id       int32   `json:"Id"`
	RndStr   string  `json:"RandomString"`
	RndFloat float32 `json:"RandomFloat"`
	RndInt   int32   `json:"RandomInteger"`
}
