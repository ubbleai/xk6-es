package es

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.k6.io/k6/output"
)

func TestConfig(t *testing.T) {
	t.Parallel()
	// TODO: add more cases
	testCases := map[string]struct {
		jsonRaw json.RawMessage
		env     map[string]string
		arg     string
		config  Config
		err     string
	}{
		"default": {
			config: Config{
				Address:      "http://127.0.0.1:9200",
				PushInterval: 1 * time.Second,
				MaxBulkSize:  2048,
			},
		},

		"overwrite": {
			env: map[string]string{"K6_ES_ADDRESS": "else", "K6_ES_PUSH_INTERVAL": "4ms"},
			config: Config{
				Address:      "else",
				PushInterval: 4 * time.Millisecond,
				MaxBulkSize:  2048,
			},
		},

		"early error": {
			env: map[string]string{"K6_ES_ADDRESS": "else", "K6_ES_PUSH_INTERVAL": "4something"},
			config: Config{
				Address:      "else",
				PushInterval: 1 * time.Second,
				MaxBulkSize:  2048,
			},
			err: `time: unknown unit "something" in duration "4something"`,
		},
	}

	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			config, err := NewConfig(output.Params{Environment: testCase.env})
			if testCase.err != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), testCase.err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, testCase.config, config)
		})
	}
}
