package myTeams

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeOAuth(t *testing.T) {
	subtests := []struct {
		name         string
		secret       string
		id           string
		expectedData *OAuth
		expectedErr  bool
	}{
		{
			name:         "happy path",
			secret:       "test_secret",
			id:           "test_id",
			expectedData: &OAuth{secret: "test_secret", id: "test_id"},
			expectedErr:  false,
		},
		{
			name:         "secret not set",
			id:           "test_id",
			secret:       "",
			expectedData: nil,
			expectedErr:  true,
		},
		{
			name:         "id not set",
			secret:       "test_secret",
			expectedData: nil,
			expectedErr:  true,
		},
	}

	for _, subtest := range subtests {
		t.Run(subtest.name, func(t *testing.T) {
			t.Setenv("GITHUB_CLIENT_SECRET", subtest.secret)
			t.Setenv("GITHUB_CLIENT_ID", subtest.id)
			data, err := MakeOAuth()

			assert.Equal(t, subtest.expectedData, data, "they should be equal")

			if subtest.expectedErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
