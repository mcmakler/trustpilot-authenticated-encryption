package trustpilotLinkGen

import (
	"encoding/base64"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestNewTrustpilotLinkGenerator(t *testing.T) {
	t.Run("Should: return error"+errorDecoding.Error(), func(t *testing.T) {
		_, err := NewTrustpilotLinkGenerator("XXXXXaGVsbG8=", "", "")
		assert.Equal(t, errorDecoding, err)
	})
	t.Run("Should: return error"+errorDecoding.Error(), func(t *testing.T) {
		_, err := NewTrustpilotLinkGenerator("", "XXXXXaGVsbG8=", "")
		assert.Equal(t, errorDecoding, err)
	})
	t.Run("Should: return nil error", func(t *testing.T) {
		_, err := NewTrustpilotLinkGenerator("", "", "")
		assert.NoError(t, err)
	})
}

func TestTrustpilotLinkGenerator_GenerateBusinessLink(t *testing.T) {
	t.Run("Should: return error"+errorMarshalling.Error(), func(t *testing.T) {
		tgen, _ := NewTrustpilotLinkGenerator("", "", "")
		_, err := tgen.GenerateBusinessLink(math.Inf(1))
		assert.Equal(t, errorMarshalling, err)
	})
	t.Run("Should: return error"+errorMarshalling.Error(), func(t *testing.T) {
		tgen, _ := NewTrustpilotLinkGenerator("", "", "")
		_, err := tgen.GenerateBusinessLink([]byte{1, 2, 3})
		assert.Equal(t, errorLinkGenerating, err)
	})
	t.Run("Should: return nil error", func(t *testing.T) {
		//32-bit "key"
		tmp := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1}
		key := base64.StdEncoding.EncodeToString(tmp)
		tgen, _ := NewTrustpilotLinkGenerator(
			key,
			key,
			"")
		res, err := tgen.GenerateBusinessLink(
			&TrustpilotUserData{
				Email: "email@email.email",
				Name:  "name",
				Ref:   "ref",
				Skus:  []string{"sku1", "sku2"},
				Tags:  []string{"tag1", "tag2"},
			})
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})
}
