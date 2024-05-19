package pokeapi

// JSON To go 工具有 时它不能准确知道你想要的字段类型
/*
如, Previous 字段应该是 string 和 Next 一样
  "next": "https://pokeapi.co/api/v2/location-area/?offset=40&limit=20",
  "previous": "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
*/
// 可空string用  *string  指针

/* json:"next" 等标签用于 JSON 序列化和反序列化操作，
指定 JSON 编码/解码时使用的字段名。*/

type LocationAreasResp struct {
	Count    int    `json:"count"`
	Next     *string `json:"next"`
	Previous *string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}