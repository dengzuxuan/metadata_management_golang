package v1

import (
	"encoding/json"
	"others-part/model"
)

func AddBusiness(info []byte) {
	businessMeta := model.AtlasBusinessMeta{}
	json.Unmarshal(info, &businessMeta)
	model.AddBusinessMeta(businessMeta)
}
