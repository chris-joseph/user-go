package binottoactions_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/smallcase/go-be-template/pkg/binotto"
	"github.com/smallcase/go-be-template/pkg/binottoactions"
	mock_store "github.com/smallcase/go-be-template/pkg/store/mock"
)

func TestIsCreated(t *testing.T) {
	type test struct {
		name             string
		venue            string
		expectedReturn   bool
		createStubReturn *binotto.Binotto
		createStubError  error
	}
	testCases := []test{
		{
			name:             "Binotto should be created",
			venue:            "mock venue",
			expectedReturn:   true,
			createStubReturn: &binotto.Binotto{Venue: "mock venue"},
			createStubError:  nil,
		},

		{
			name:             "Binotto should not be created",
			venue:            "mock venue",
			expectedReturn:   false,
			createStubReturn: nil,
			createStubError:  errors.New("Binotto could not be created"),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			binottoStore := mock_store.NewMockBinottoStore(ctrl)
			ctx := context.Background()
			binottoStore.EXPECT().Create(ctx, testCase.venue).Return(testCase.createStubReturn, testCase.createStubError) // stubbing the function
			result := binottoactions.IsCreated(context.TODO(), binottoStore, testCase.venue)
			if testCase.expectedReturn && !result || !testCase.expectedReturn && result {
				t.Errorf("Expected IsCreated to return `%v` but it returned `%v` instead.", testCase.expectedReturn, result)
			}
		})
	}
}
