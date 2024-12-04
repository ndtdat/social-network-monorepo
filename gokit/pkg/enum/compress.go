package enum

import (
	"database/sql/driver"
)

// nolint: revive
const (
	Compress_NONE    Compress = "none"
	Compress_ZSTD    Compress = "zstd"
	Compress_LZ4     Compress = "lz4"
	Compress_GZIP    Compress = "gzip"
	Compress_DEFLATE Compress = "deflate"
	Compress_BR      Compress = "br"
)

type Compress string

func (c Compress) String() string {
	return string(c)
}

func (c *Compress) Scan(value any) error {
	*c = Compress(value.([]byte)) //nolint: forcetypeassert

	return nil
}

func (c Compress) Value() (driver.Value, error) {
	return c.String(), nil
}
