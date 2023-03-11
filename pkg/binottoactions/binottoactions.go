package binottoactions

import (
	"context"

	"github.com/smallcase/go-be-template/pkg/store"
)

func IsCreated(ctx context.Context, binottoStore store.BinottoStore, venue string) bool {
	_, err := binottoStore.Create(ctx, venue)
	return err == nil
}
