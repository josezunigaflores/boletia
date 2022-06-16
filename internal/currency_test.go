package internal

import (
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTimeFilter(t *testing.T) {
	t.Parallel()
	t.Run("should return success time", func(t *testing.T) {
		t.Parallel()
		ft, err := NewTimeFilter("2022-11-28T20:15:00")
		assert.NoError(t, err)
		assert.NotNil(t, ft)
	})

	t.Run("should return error  time is empty", func(t *testing.T) {
		t.Parallel()
		_, err := NewTimeFilter("")
		assert.Error(t, err)
	})
	t.Run("should return error  time is bad", func(t *testing.T) {
		t.Parallel()
		_, err := NewTimeFilter(faker.Word())
		assert.Error(t, err)
	})
}

func TestNewCode(t *testing.T) {
	t.Parallel()
	t.Run("should get new code", func(t *testing.T) {
		t.Parallel()
		c, err := NewCode(faker.Word())
		assert.NoError(t, err)
		assert.NotNil(t, c)
	})
	t.Run("should get error tries to create new code", func(t *testing.T) {
		t.Parallel()
		_, err := NewCode("")
		assert.Error(t, err)
	})
	t.Run("should return is 'all' is the code for return flag", func(t *testing.T) {
		c, err := NewCode(faker.Word())
		assert.NoError(t, err)
		assert.Equal(t, false, c.IsAll())
	})
}
