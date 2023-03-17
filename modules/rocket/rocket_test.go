package rocket

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRocketService(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	t.Run("tests get rocket by id", func(t *testing.T) {
		rocRepoMock := NewMockRepository(mockCtrl)
		id := "UUID-1"
		rocRepoMock.EXPECT().GetByID(id).Return(Rocket{ID: id}, nil)

		rocService := GetService(rocRepoMock)
		roc, err := rocService.GetRocketByID(context.Background(), id)
		assert.NoError(t, err)
		assert.Equal(t, "UUID-1", roc.ID)
	})

	t.Run("tests insert rocket", func(t *testing.T) {
		rocRepoMock := NewMockRepository(mockCtrl)
		id := "UUID-1"
		rkt := Rocket{ID: id}
		rocRepoMock.EXPECT().Insert(rkt).Return(Rocket{ID: id}, nil)

		rocService := GetService(rocRepoMock)
		roc, err := rocService.InsertRocket(context.Background(), rkt)
		assert.NoError(t, err)
		assert.Equal(t, "UUID-1", roc.ID)
	})

	t.Run("tests get rocket by id", func(t *testing.T) {
		rocRepoMock := NewMockRepository(mockCtrl)
		id := "UUID-1"
		rocRepoMock.EXPECT().Remove(id).Return(nil)

		rocService := GetService(rocRepoMock)
		err := rocService.RemoveRocket(context.Background(), id)
		assert.NoError(t, err)
	})
}
