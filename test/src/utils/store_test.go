package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"test/src/test_common"
	"test/src/test_impl"
	"test/src/test_pb"
	"testing"

	"google.golang.org/protobuf/proto"
)

func TestStore(t *testing.T) {
	all := test_impl.InitAllianceStoreInfo(test_common.ItemData)
	all.StoreItem(2, 5, 1)
	all.StoreItem(2, 100, 1)
	all.StoreItem(1, 5, 2)
	all.StoreItem(1, 3, 3)

	all.ClearUp()

	fmt.Println(all.Items)
}

func TestPb(t *testing.T) {
	in, err := ioutil.ReadFile("testItem.data")
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}

	items := &test_pb.TestItem_Array{}

	if err := proto.Unmarshal(in, items); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}

}
