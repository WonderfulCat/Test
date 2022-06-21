package utils

import (
	"fmt"
	"test/src/test_impl"
	"testing"
)

func TestStore(t *testing.T) {
	all := test_impl.InitAllianceStoreInfo()
	all.StoreItem(2, 5, 1)
	all.StoreItem(1, 5, 2)
	all.StoreItem(1, 3, 3)

	all.ClearUp()

	fmt.Println(all.Items)
}
