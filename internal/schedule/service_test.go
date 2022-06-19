package schedule

import (
	"boletia/internal"
	"boletia/internal/mocks"
	"boletia/kit/event/eventmocks"
	"context"
	"testing"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/mock"
)

func TestService_Do(t *testing.T) {
	t.Run("Should do infinite loop ", func(t *testing.T) {
		mockhttp := &mocks.RepositoryHttp{}
		mockEvnet := &eventmocks.Event{}
		c := make(internal.Currencies, 0)
		c = append(c, internal.Currency{
			Code:          faker.Word(),
			Value:         0,
			LastUpdatedAt: time.Now(),
		})
		mockhttp.On("GetCurrencies").
			Return(c, &internal.MetaData{}, mockEvnet, nil)
		mockCurrency := &mocks.RepositoryCurrency{}
		mockCurrency.On("CreateCurrencies", mock.Anything, mock.Anything).Return(nil)
		mocksBus := &eventmocks.Bus{}
		mocksBus.On("Publish", mock.Anything, mock.Anything).Return(nil)
		s := NewServiceSchedule(mockhttp, mockCurrency, 5, 5, mocksBus)
		background, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		go s.Do()
		<-background.Done()
	})
}
