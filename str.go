package deercoder

/*max upload file size*/
var MaxUploadMemory int64

//page thing
var ClientPageStr = GetConfigValue("clientPage")
var EveryPageStr = GetConfigValue("everyPage")

//struct value
type Value struct {
	Value string `json:"value"`
}