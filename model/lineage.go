package model

type LineageType struct {
	BaseEntityGUID string `json:"baseEntityGuid"`
	LineageDepth   int    `json:"lineageDepth"`
	Relations      []struct {
		FromEntityID   string `json:"fromEntityId"`
		ToEntityID     string `json:"toEntityId"`
		RelationshipID string `json:"relationshipId"`
	} `json:"relations"`
}
type LineageInfo struct {
	Id       int    `gorm:"id" json:"id"`
	BaseGuid string `gorm:"baseGuid" json:"baseGuid"`
	Guid     string `gorm:"guid" json:"guid"`
	Level    int    `gorm:"level" json:"level"`
	Site     int    `gorm:"site" json:"site"`
	X        int    `gorm:"x" json:"x"`
	Y        int    `gorm:"y" json:"y"`
}
