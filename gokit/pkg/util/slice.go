package util

func SplitSlice[T any](slice []T, batchSize int) [][]T {
	if batchSize == 0 {
		batchSize = len(slice)
	}

	var chunks [][]T
	for {
		if len(slice) == 0 {
			break
		}

		if len(slice) < batchSize {
			batchSize = len(slice)
		}

		chunks = append(chunks, slice[0:batchSize])
		slice = slice[batchSize:]
	}

	return chunks
}
