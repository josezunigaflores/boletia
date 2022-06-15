package schedule

import (
	"testing"
)

func TestService_Do(t *testing.T) {
	t.Run("Should do infinite loop ", func(t *testing.T) {
		s := Service{}
		s.Do()
	})
}
