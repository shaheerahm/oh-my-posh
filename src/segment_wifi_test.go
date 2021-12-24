package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWiFiSegment(t *testing.T) {
	cases := []struct {
		Case            string
		ExpectedString  string
		ExpectedEnabled bool
		Network         *wifiInfo
		WifiError       error
		DisplayError    bool
	}{
		{
			Case: "No error and nil network",
		},
		{
			Case:      "Error and nil network",
			WifiError: errors.New("oh noes"),
		},
		{
			Case:            "Display error and nil network",
			WifiError:       errors.New("oh noes"),
			ExpectedString:  "oh noes",
			DisplayError:    true,
			ExpectedEnabled: true,
		},
	}

	for _, tc := range cases {
		env := new(MockedEnvironment)
		env.On("getPlatform", nil).Return(windowsPlatform)
		env.On("isWsl", nil).Return(false)
		env.On("getWifiNetwork", nil).Return(tc.Network, tc.WifiError)

		w := &wifi{
			env: env,
			props: map[Property]interface{}{
				DisplayError: tc.DisplayError,
			},
		}

		assert.Equal(t, tc.ExpectedEnabled, w.enabled(), tc.Case)
		if tc.WifiError != nil && tc.DisplayError {
			assert.Equal(t, tc.ExpectedString, w.string(), tc.Case)
		}
	}
}
