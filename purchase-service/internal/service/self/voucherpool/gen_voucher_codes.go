package voucherpool

import (
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/random"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/set"
)

var (
	codeCharsets = []rune("ABCDEFGHJKLMNPQRSTUVWXYZ")
)

func (s *Service) genVoucherCodes(length int, qty int64) ([]string, error) {
	var (
		voucherCodeSet = set.New[string]()
	)

	// loop to generate gift code until sufficient qty
	counter := int64(0)
	for {
		code := s.genVoucherCode(length)

		if !voucherCodeSet.Contains(code) {
			voucherCodeSet.Add(code)
			counter++
		}

		if counter == qty {
			break
		}
	}

	return voucherCodeSet.ItemArray(), nil
}

func (s *Service) genVoucherCode(codeLength int) string {
	code := generateRandomChars(codeCharsets, codeLength)

	return string(code)
}

// nolint intrange
func generateRandomChars(charsets []rune, resultCharsLen int) []rune {
	randomChars := make([]rune, 0, resultCharsLen)
	randomNumberMax := int64(len(charsets)) - 1

	for i := 0; i < resultCharsLen; i++ {
		index := random.Int64(0, randomNumberMax)
		randomChars = append(randomChars, charsets[index])
	}

	return randomChars
}

func (s *Service) getNotExistedCodes(totalCodes, existedCodes []string) []string {
	existedCodeSet := set.New[string]()
	for _, ec := range existedCodes {
		existedCodeSet.Add(ec)
	}

	// Remove allocated gift codes
	var notExistedCodes []string
	for _, codes := range totalCodes {
		if !existedCodeSet.Contains(codes) {
			notExistedCodes = append(notExistedCodes, codes)
		}
	}

	return notExistedCodes
}
