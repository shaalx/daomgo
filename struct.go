package daomgo

import (
	"encoding/json"
	"github.com/shaalx/goutils"
)

func ConvStruct(i interface{}, ret interface{}) bool {
	b := I2JsonBytes(i)
	err := json.Unmarshal(b, &ret)
	return !goutils.CheckErr(err)
}
