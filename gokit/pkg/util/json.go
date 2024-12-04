package util

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
)

func MustMarshalJSON(o any) []byte {
	data, err := json.Marshal(o)
	if err != nil {
		panic(fmt.Sprintf("cannot marshal json for %v due to %v", o, err))
	}

	return data
}

func MustUnmarshalJSON(data []byte, o any) {
	err := json.Unmarshal(data, &o)
	if err != nil {
		panic(fmt.Sprintf("cannot unmarshal json for %v due to %v", o, err))
	}
}

func PrettyPrintJSON(logger *zap.Logger, title string, object any) {
	logger.Debug(fmt.Sprintf("%s: %s", title, PrettyJSON(object)))
}

func PrettyJSON(object any) string {
	prettyJSON, _ := json.MarshalIndent(object, "", "    ")

	return string(prettyJSON)
}

func PrettyPrintJSONWithInfo(logger *zap.Logger, title string, object any, ignoreJSONFields []string) {
	defer func() {
		if rec := recover(); rec != nil {
			logger.Error(fmt.Sprintf("Failed to pretty print json due to: %v", rec))
		}
	}()

	// Convert struct to map
	objectMap, err := StructToMap(object)
	if err != nil {
		logger.Error(fmt.Sprintf("Cannot parse struct to map due to: %v", err))

		return
	}

	// Remove field in map
	for _, f := range ignoreJSONFields {
		delete(objectMap, f)
	}

	// Convert map to struct
	results, err := MapToStruct(objectMap)
	if err != nil {
		logger.Error(fmt.Sprintf("Cannot parse map to struct due to: %v", err))

		return
	}

	logger.Info(fmt.Sprintf("%s: %s", title, PrettyJSON(results)))
}
