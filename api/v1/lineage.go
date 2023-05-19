package v1

import (
	"encoding/json"
	"others-part/model"
)

func analysisLineage(s []byte) []model.LineageInfo {
	var lineage model.LineageType
	var lineageInfos []model.LineageInfo
	allGuids := make(map[string]int)
	//relationid:toid->fromid
	guidRelation := make(map[string]map[string]string)
	guidLevels := make(map[string]int)
	guidSite := make(map[int]int)
	_ = json.Unmarshal(s, &lineage)
	//1为根节点 2为非根节点
	for _, relation := range lineage.Relations {
		if guidRelation[relation.RelationshipID] == nil {
			guidRelation[relation.RelationshipID] = make(map[string]string)
		}
		guidRelation[relation.RelationshipID] = map[string]string{relation.ToEntityID: relation.FromEntityID}
		allGuids[relation.FromEntityID] = 2
		if allGuids[relation.ToEntityID] != 2 {
			allGuids[relation.ToEntityID] = 1
		}
	}
	for len(guidRelation) != 0 {
		for relationshipId, ids := range guidRelation {
			for toid, fromid := range ids {
				if allGuids[toid] == 1 {
					guidLevels[toid] = 1
					guidLevels[fromid] = 2
					delete(guidRelation, relationshipId)
				} else if guidLevels[toid] != 0 {
					guidLevels[fromid] = guidLevels[toid] + 1
					delete(guidRelation, relationshipId)
				}
			}

		}
	}
	for guid, level := range guidLevels {
		guidSite[level]++
		lineageInfos = append(lineageInfos, model.LineageInfo{
			BaseGuid: lineage.BaseEntityGUID,
			Guid:     guid,
			Level:    level,
			Site:     guidSite[level],
		})
	}
	flagUp, flagDown := 1, 1
	for i, info := range lineageInfos {
		y := 300
		site := info.Site
		level := info.Level
		total := guidSite[level]
		lineageInfos[i].X = 300 + (lineage.LineageDepth-level)*200
		if total >= 2 {
			if total%2 == 0 {
				if site <= total/2 {
					y = 300 + 50*flagUp
					flagUp++
				} else {
					y = 300 - 50*flagDown
					flagDown++
				}
			} else {
				if site <= total/2 {
					y = 300 + 50*flagUp
					flagUp++
				} else if site >= (total/2)+1 {
					y = 300 - 50*flagDown
					flagDown++
				}
			}
		}
		lineageInfos[i].Y = y
		if site == total {
			flagUp, flagDown = 1, 1
		}
	}
	return lineageInfos
	//infoJson, _ := json.Marshal(lineageInfos)
	//return infoJson
}
