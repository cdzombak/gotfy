package gotfy

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestUnixTime_UnmarshalJSON(t *testing.T) {
	r := require.New(t)

	target := UnixTime{}
	r.NoError(target.UnmarshalJSON([]byte("1685150791")))
	r.Equal("2023-05-27T01:26:31Z", target.In(time.UTC).Format(time.RFC3339))
}
