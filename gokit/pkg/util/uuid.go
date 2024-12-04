package util

import "github.com/google/uuid"

var NullUUIDStr = uuid.Nil.String()

func UUIDFromStr(idStr string) *uuid.UUID {
	id, err := uuid.Parse(idStr)
	if err != nil {
		return nil
	}

	return &id
}

func UUIDToStr(id *uuid.UUID) string {
	str := ""
	if id != nil {
		str = id.String()
	}

	return str
}

func UUIDsToStrings(ids []*uuid.UUID) []string {
	var results []string
	for _, i := range ids {
		results = append(results, UUIDToStr(i))
	}

	return results
}

func UUIDsFromStrings(idStrs []string) []*uuid.UUID {
	var ids []*uuid.UUID
	for _, idStr := range idStrs {
		id := UUIDFromStr(idStr)
		if id != nil {
			ids = append(ids, id)
		}
	}

	return ids
}

func MustParseUUIDsFromStrings(idStrs []string) []uuid.UUID {
	var ids []uuid.UUID
	for _, idStr := range idStrs {
		id := uuid.MustParse(idStr)
		ids = append(ids, id)
	}

	return ids
}

func UUIDInstancesToStrings(ids []uuid.UUID) []string {
	var results []string
	for _, id := range ids {
		results = append(results, id.String())
	}

	return results
}

func UUIDInstancesFromStrings(ids []string) []uuid.UUID {
	var results []uuid.UUID
	for _, id := range ids {
		results = append(results, uuid.MustParse(id))
	}

	return results
}
