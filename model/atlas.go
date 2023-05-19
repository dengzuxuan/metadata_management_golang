package model

type AtlasLineage struct {
	BaseEntityGUID   string `json:"baseEntityGuid"`
	LineageDirection string `json:"lineageDirection"`
	LineageDepth     int    `json:"lineageDepth"`
	GUIDEntityMap    struct {
	} `json:"-"`
	Relations []struct {
		FromEntityID   string `json:"fromEntityId"`
		ToEntityID     string `json:"toEntityId"`
		RelationshipID string `json:"relationshipId"`
	} `json:"relations"`
}

type AtlasSearchPreType struct {
	SearchResults struct {
		QueryType        string `json:"queryType"`
		SearchParameters struct {
		} `json:"searchParameters"`
		QueryText string `json:"queryText"`
		Entities  []struct {
			TypeName   string `json:"typeName"`
			Attributes struct {
				Owner         string `json:"owner"`
				QualifiedName string `json:"qualifiedName"`
				Name          string `json:"name"`
			} `json:"attributes"`
			GUID        string `json:"guid"`
			DisplayText string `json:"displayText"`
		} `json:"entities"`
		ApproximateCount int `json:"approximateCount"`
	} `json:"searchResults"`
	AggregationMetrics struct {
		TypeName []struct {
		} `json:"aggregationMetrics"`
	}
}

type AtlasOtherInfo struct {
}
type AtlasType []struct {
	GUID        string `json:"guid"`
	Name        string `json:"name"`
	ServiceType string `json:"serviceType,omitempty"`
	Category    string `json:"category"`
}

type AtlasFindType struct {
	QueryType        string `json:"queryType"`
	SearchParameters struct {
		TypeName                        string `json:"typeName"`
		ExcludeDeletedEntities          bool   `json:"excludeDeletedEntities"`
		IncludeClassificationAttributes bool   `json:"includeClassificationAttributes"`
		IncludeSubTypes                 bool   `json:"includeSubTypes"`
		IncludeSubClassifications       bool   `json:"includeSubClassifications"`
		Limit                           int    `json:"limit"`
		Offset                          int    `json:"offset"`
	} `json:"searchParameters"`
	Entities []struct {
		TypeName   string `json:"typeName"`
		Attributes struct {
			CreateTime    int64  `json:"createTime"`
			QualifiedName string `json:"qualifiedName"`
			Name          string `json:"name"`
		} `json:"attributes"`
		GUID                string        `json:"guid"`
		Status              string        `json:"status"`
		DisplayText         string        `json:"displayText"`
		ClassificationNames []interface{} `json:"classificationNames"`
		Classifications     []interface{} `json:"classifications"`
		MeaningNames        []interface{} `json:"meaningNames"`
		Meanings            []interface{} `json:"meanings"`
		IsIncomplete        bool          `json:"isIncomplete"`
		Labels              []interface{} `json:"labels"`
	} `json:"entities"`
	ApproximateCount int `json:"approximateCount"`
}

type AtlasAudit []struct {
	EntityID  string      `json:"entityId"`
	Timestamp int64       `json:"timestamp"`
	User      string      `json:"user"`
	Action    string      `json:"action"`
	Details   string      `json:"details"`
	EventKey  string      `json:"eventKey"`
	Entity    interface{} `json:"entity"`
	Type      interface{} `json:"type"`
}
type AtlasClassification struct {
	EnumDefs           []interface{} `json:"enumDefs"`
	StructDefs         []interface{} `json:"structDefs"`
	ClassificationDefs []struct {
		Category      string        `json:"category"`
		GUID          string        `json:"guid"`
		CreatedBy     string        `json:"createdBy"`
		UpdatedBy     string        `json:"updatedBy"`
		CreateTime    int64         `json:"createTime"`
		UpdateTime    int64         `json:"updateTime"`
		Version       int           `json:"version"`
		Name          string        `json:"name"`
		Description   string        `json:"description"`
		TypeVersion   string        `json:"typeVersion"`
		AttributeDefs []interface{} `json:"attributeDefs"`
		SuperTypes    []interface{} `json:"superTypes"`
		EntityTypes   []interface{} `json:"entityTypes"`
		SubTypes      []interface{} `json:"subTypes"`
	} `json:"classificationDefs"`
	EntityDefs           []interface{} `json:"entityDefs"`
	RelationshipDefs     []interface{} `json:"relationshipDefs"`
	BusinessMetadataDefs []interface{} `json:"businessMetadataDefs"`
}
type AtlasGlossary []struct {
	GUID             string `json:"guid"`
	QualifiedName    string `json:"qualifiedName"`
	Name             string `json:"name"`
	ShortDescription string `json:"shortDescription"`
	LongDescription  string `json:"longDescription"`
	Terms            []struct {
		TermGUID     string `json:"termGuid"`
		RelationGUID string `json:"relationGuid"`
		DisplayText  string `json:"displayText"`
	} `json:"terms,omitempty"`
}
type AtlasBusinessMeta struct {
	EnumDefs             []interface{} `json:"enumDefs"`
	StructDefs           []interface{} `json:"structDefs"`
	ClassificationDefs   []interface{} `json:"classificationDefs"`
	EntityDefs           []interface{} `json:"entityDefs"`
	RelationshipDefs     []interface{} `json:"relationshipDefs"`
	BusinessMetadataDefs []struct {
		Category      string `json:"category"`
		GUID          string `json:"guid"`
		CreatedBy     string `json:"createdBy"`
		UpdatedBy     string `json:"updatedBy"`
		CreateTime    int64  `json:"createTime"`
		UpdateTime    int64  `json:"updateTime"`
		Version       int    `json:"version"`
		Name          string `json:"name"`
		Description   string `json:"description"`
		TypeVersion   string `json:"typeVersion"`
		AttributeDefs []struct {
			Name                  string `json:"name"`
			TypeName              string `json:"typeName"`
			IsOptional            bool   `json:"isOptional"`
			Cardinality           string `json:"cardinality"`
			ValuesMinCount        int    `json:"valuesMinCount"`
			ValuesMaxCount        int    `json:"valuesMaxCount"`
			IsUnique              bool   `json:"isUnique"`
			IsIndexable           bool   `json:"isIndexable"`
			IncludeInNotification bool   `json:"includeInNotification"`
			SearchWeight          int    `json:"searchWeight"`
			Options               struct {
				ApplicableEntityTypes string `json:"applicableEntityTypes"`
				MaxStrLength          string `json:"maxStrLength"`
			} `json:"options"`
		} `json:"attributeDefs"`
	} `json:"businessMetadataDefs"`
}
type AtlasEntityTypes struct {
	Data struct {
		General struct {
			CollectionTime int64 `json:"collectionTime"`
			EntityCount    int   `json:"entityCount"`
			Stats          struct {
				NotificationLastMessageProcessedTime int `json:"Notification:lastMessageProcessedTime"`
				NotificationCurrentDayEntityUpdates  int `json:"Notification:currentDayEntityUpdates"`
				NotificationCurrentDayFailed         int `json:"Notification:currentDayFailed"`
				NotificationTopicDetails             struct {
				} `json:"Notification:topicDetails"`
				NotificationCurrentHourStartTime      int64  `json:"Notification:currentHourStartTime"`
				NotificationPreviousDayEntityCreates  int    `json:"Notification:previousDayEntityCreates"`
				NotificationCurrentHourAvgTime        int    `json:"Notification:currentHourAvgTime"`
				NotificationPreviousHour              int    `json:"Notification:previousHour"`
				NotificationTotalUpdates              int    `json:"Notification:totalUpdates"`
				NotificationPreviousHourEntityUpdates int    `json:"Notification:previousHourEntityUpdates"`
				ServerStatusIndexStore                string `json:"Server:statusIndexStore"`
				NotificationTotalAvgTime              int    `json:"Notification:totalAvgTime"`
				NotificationCurrentDayEntityDeletes   int    `json:"Notification:currentDayEntityDeletes"`
				NotificationCurrentHourEntityCreates  int    `json:"Notification:currentHourEntityCreates"`
				NotificationTotalDeletes              int    `json:"Notification:totalDeletes"`
				NotificationPreviousHourEntityDeletes int    `json:"Notification:previousHourEntityDeletes"`
				ServerStartTimeStamp                  int64  `json:"Server:startTimeStamp"`
				NotificationPreviousHourEntityCreates int    `json:"Notification:previousHourEntityCreates"`
				NotificationCurrentDayStartTime       int64  `json:"Notification:currentDayStartTime"`
				ServerUpTime                          string `json:"Server:upTime"`
				NotificationCurrentDay                int    `json:"Notification:currentDay"`
				NotificationCurrentHourEntityUpdates  int    `json:"Notification:currentHourEntityUpdates"`
				NotificationCurrentHour               int    `json:"Notification:currentHour"`
				NotificationTotalFailed               int    `json:"Notification:totalFailed"`
				NotificationCurrentDayEntityCreates   int    `json:"Notification:currentDayEntityCreates"`
				NotificationCurrentHourEntityDeletes  int    `json:"Notification:currentHourEntityDeletes"`
				ServerStatusBackendStore              string `json:"Server:statusBackendStore"`
				NotificationTotalCreates              int    `json:"Notification:totalCreates"`
				NotificationPreviousDayEntityUpdates  int    `json:"Notification:previousDayEntityUpdates"`
				NotificationCurrentHourFailed         int    `json:"Notification:currentHourFailed"`
				NotificationCurrentDayAvgTime         int    `json:"Notification:currentDayAvgTime"`
				NotificationPreviousHourFailed        int    `json:"Notification:previousHourFailed"`
				NotificationTotal                     int    `json:"Notification:total"`
				NotificationPreviousDayEntityDeletes  int    `json:"Notification:previousDayEntityDeletes"`
				ServerActiveTimeStamp                 int64  `json:"Server:activeTimeStamp"`
				NotificationPreviousDay               int    `json:"Notification:previousDay"`
				NotificationPreviousHourAvgTime       int    `json:"Notification:previousHourAvgTime"`
				NotificationPreviousDayFailed         int    `json:"Notification:previousDayFailed"`
				NotificationPreviousDayAvgTime        int    `json:"Notification:previousDayAvgTime"`
			} `json:"stats"`
			TagCount        int `json:"tagCount"`
			TypeUnusedCount int `json:"typeUnusedCount"`
			TypeCount       int `json:"typeCount"`
		} `json:"general"`
		System struct {
			Memory struct {
				HeapInit         string `json:"heapInit"`
				HeapMax          string `json:"heapMax"`
				HeapCommitted    string `json:"heapCommitted"`
				HeapUsed         string `json:"heapUsed"`
				NonHeapInit      string `json:"nonHeapInit"`
				NonHeapMax       string `json:"nonHeapMax"`
				NonHeapCommitted string `json:"nonHeapCommitted"`
				NonHeapUsed      string `json:"nonHeapUsed"`
				MemoryPoolUsages struct {
					PSEdenSpace struct {
						Init      int `json:"init"`
						Used      int `json:"used"`
						Committed int `json:"committed"`
						Max       int `json:"max"`
					} `json:"PS Eden Space"`
					PSSurvivorSpace struct {
						Init      int `json:"init"`
						Used      int `json:"used"`
						Committed int `json:"committed"`
						Max       int `json:"max"`
					} `json:"PS Survivor Space"`
					PSOldGen struct {
						Init      int `json:"init"`
						Used      int `json:"used"`
						Committed int `json:"committed"`
						Max       int `json:"max"`
					} `json:"PS Old Gen"`
				} `json:"memory_pool_usages"`
			} `json:"memory"`
			Os struct {
				OsSpec  string `json:"os.spec"`
				OsVcpus string `json:"os.vcpus"`
			} `json:"os"`
			Runtime struct {
				Name    string `json:"name"`
				Version string `json:"version"`
			} `json:"runtime"`
		} `json:"system"`
		Tag struct {
			TagEntities struct {
				ClassificationtTest02  int `json:"classificationtTest02"`
				ClassififcationsTest01 int `json:"classififcationsTest01"`
			} `json:"tagEntities"`
		} `json:"tag"`
		Entity struct {
			EntityDeletedTypeAndSubTypes struct {
			} `json:"entityDeleted-typeAndSubTypes"`
			EntityActiveTypeAndSubTypes struct {
				DataSet              int `json:"DataSet"`
				HiveDbDdl            int `json:"hive_db_ddl"`
				Process              int `json:"Process"`
				HiveTable            int `json:"hive_table"`
				Ddl                  int `json:"ddl"`
				HiveDb               int `json:"hive_db"`
				HiveProcess          int `json:"hive_process"`
				HiveStoragedesc      int `json:"hive_storagedesc"`
				AtlasGlossary        int `json:"AtlasGlossary"`
				Referenceable        int `json:"Referenceable"`
				ProcessExecution     int `json:"ProcessExecution"`
				AtlasGlossaryTerm    int `json:"AtlasGlossaryTerm"`
				HiveColumnLineage    int `json:"hive_column_lineage"`
				Internal             int `json:"__internal"`
				Asset                int `json:"Asset"`
				HiveColumn           int `json:"hive_column"`
				HiveProcessExecution int `json:"hive_process_execution"`
				HiveTableDdl         int `json:"hive_table_ddl"`
			} `json:"entityActive-typeAndSubTypes"`
			EntityActive struct {
				HiveStoragedesc      int `json:"hive_storagedesc"`
				AtlasGlossary        int `json:"AtlasGlossary"`
				AtlasGlossaryTerm    int `json:"AtlasGlossaryTerm"`
				HiveColumnLineage    int `json:"hive_column_lineage"`
				HiveDbDdl            int `json:"hive_db_ddl"`
				HiveTable            int `json:"hive_table"`
				HiveColumn           int `json:"hive_column"`
				HiveProcessExecution int `json:"hive_process_execution"`
				HiveDb               int `json:"hive_db"`
				HiveProcess          int `json:"hive_process"`
				HiveTableDdl         int `json:"hive_table_ddl"`
			} `json:"entityActive"`
			EntityShell struct {
			} `json:"entityShell"`
			EntityShellTypeAndSubTypes struct {
			} `json:"entityShell-typeAndSubTypes"`
			EntityDeleted struct {
			} `json:"entityDeleted"`
		} `json:"entity"`
	} `json:"data"`
}

type AtlasSearch struct {
	Entities []struct {
		Attributes struct {
			QualifiedName string `json:"qualifiedName"`
		} `json:"attributes"`
		ClassificationNames []string `json:"classificationNames"`
		Classifications     []struct {
			Attributes struct {
			} `json:"attributes"`
			EntityGUID string `json:"entityGuid"`
			TypeName   string `json:"typeName"`
		} `json:"classifications"`
		GUID         string   `json:"guid"`
		MeaningNames []string `json:"meaningNames"`
		Status       string   `json:"status"`
	} `json:"entities"`
	Info struct {
		ID          int    `json:"id"`
		Typename    string `json:"typename"`
		Userid      int    `json:"userid"`
		Username    string `json:"username"`
		Avatar      string `json:"avatar"`
		Createtime  string `json:"createtime"`
		Description string `json:"description"`
	} `json:"info"`
}
