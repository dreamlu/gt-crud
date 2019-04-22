// @author  dreamlu
package deercoder

// max upload file size
var MaxUploadMemory int64

// page thing
var ClientPageStr = GetDevModeConfig("clientPage")
var EveryPageStr = GetDevModeConfig("everyPage")

// struct value
type Value struct {
	Value string `json:"value"`
}


// ID struct
type ID struct {
	ID int64 `json:"id"`
}