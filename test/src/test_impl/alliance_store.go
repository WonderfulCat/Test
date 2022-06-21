package test_impl

import (
	"fmt"
	"math"
	"sort"
	"test/src/test_constant"
	"test/src/test_model"
)

type AllianceStoreInfo struct {
	IncreaseCapacityCount int32   //己扩容次数
	CapacityLeft          int32   //当前仓库剩余容量 (未使用格子数*每格最大堆叠数)
	Items                 []*Item //仓库物品信息 (index: 0 ~ CapacityCur-1)
}

//堆叠物品返回信息
type getStoreItemPosInfo struct {
	pos  int32 //堆叠位置
	item *Item //堆叠物品信息
}

func InitAllianceStoreInfo() *AllianceStoreInfo {
	return &AllianceStoreInfo{
		IncreaseCapacityCount: 0,
		CapacityLeft:          test_constant.STORE_INIT_CAPACITY * test_constant.STORE_ITEM_STACK_MAX_NUM,
		//初始化仓库容量 = 基本大小+可扩容次数*单次可扩容空间.  为减少slice后续扩容消耗, cap空间可以分配更大一些, 以备后续扩展.
		Items: make([]*Item, test_constant.STORE_INIT_CAPACITY, test_constant.STORE_INIT_CAPACITY+test_constant.STORE_INCREASE_CAPACITY_COUNT*test_constant.STORE_INCREASE_CAPACITY),
	}
}

func (c *AllianceStoreInfo) IncreaseCapacity() *test_model.ResponseInfo {
	if c.IncreaseCapacityCount >= test_constant.STORE_INCREASE_CAPACITY_COUNT {
		return &test_model.ResponseInfo{Code: test_constant.RES_ERR, Msg: test_constant.RES_ERR_MSG_1}
	}

	//adding len for slice
	c.IncreaseCapacityCount++
	c.CapacityLeft += test_constant.STORE_INCREASE_CAPACITY * test_constant.STORE_ITEM_STACK_MAX_NUM
	c.Items = append(c.Items, make([]*Item, test_constant.STORE_INCREASE_CAPACITY)...)

	return &test_model.ResponseInfo{Code: test_constant.RES_OK}
}

func (c *AllianceStoreInfo) StoreItem(itemId, itemNum, index int32) *test_model.ResponseInfo {
	if index <= 0 || index > int32(len(c.Items)) {
		return &test_model.ResponseInfo{Code: test_constant.RES_ERR, Msg: test_constant.RES_ERR_MSG_3}
	}

	pos := c.GetStoreItemPos(c.CapacityLeft, itemId, itemNum, index)
	if pos == nil {
		return &test_model.ResponseInfo{Code: test_constant.RES_ERR, Msg: test_constant.RES_ERR_MSG_6}
	}

	//己放入仓库道具数量
	usedNum := 0

	for _, v := range pos {
		if c.Items[v.pos] == nil {
			c.Items[v.pos] = &Item{ID: v.item.ID, Number: v.item.Number}
		} else {
			c.Items[v.pos].Number += v.item.Number
		}

		//记录己放入仓库道具数量
		usedNum += int(v.item.Number)
		//更新仓库剩余容量
		c.CapacityLeft -= v.item.Number
	}

	return &test_model.ResponseInfo{Code: test_constant.RES_OK, Msg: fmt.Sprintf(test_constant.RES_ERR_MSG_7, usedNum, itemNum-int32(usedNum))}
}

//获取堆叠物品放置信息
func (c *AllianceStoreInfo) GetStoreItemPos(capacityLeft, itemId, itemNum, index int32) (ret []*getStoreItemPosInfo) {
	//最大容量,物品数量及索引判断
	if capacityLeft <= 0 || itemNum <= 0 || index > int32(len(c.Items)) {
		return ret
	}

	item := c.Items[index-1]

	//空=当前格子没有道具
	if item == nil {
		//计算放入当前格子内道具数量 = min(剩余容量,min(道具数量,最大堆叠数))
		num := int32(math.Min(float64(capacityLeft), math.Min(float64(itemNum), float64(test_constant.STORE_ITEM_STACK_MAX_NUM))))
		//存储放入此格子道具数量
		ret = append(ret, &getStoreItemPosInfo{pos: index - 1, item: &Item{ID: itemId, Number: num}})
		//继续计算下一个存放位置
		ret = append(ret, c.GetStoreItemPos(capacityLeft-num, itemId, itemNum-num, index+1)...)
	} else {
		//当前位置道具ID不同则继续判断下一个格子
		if item.ID != itemId {
			ret = append(ret, c.GetStoreItemPos(capacityLeft, itemId, itemNum, index+1)...)
		}

		//当前位置道具ID相等&&未达到堆叠上限
		if item.ID == itemId && item.Number < test_constant.STORE_ITEM_STACK_MAX_NUM {
			//计算放入此格子内道具数量 = min(剩余容量,最大堆叠数量-己堆叠数量)
			canStoreNum := int32(math.Min(float64(capacityLeft), float64(test_constant.STORE_ITEM_STACK_MAX_NUM-item.Number)))
			//存储放入此格子道具数量
			ret = append(ret, &getStoreItemPosInfo{pos: index - 1, item: &Item{ID: itemId, Number: canStoreNum}})
			//继续计算下一个存放位置
			ret = append(ret, c.GetStoreItemPos(capacityLeft-canStoreNum, itemId, itemNum-canStoreNum, index+1)...)
		}
	}

	return ret
}

func (c *AllianceStoreInfo) DestoryItem(index int32) *test_model.ResponseInfo {
	//index check
	if index <= 0 || index > int32(len(c.Items)) {
		return &test_model.ResponseInfo{Code: test_constant.RES_ERR, Msg: test_constant.RES_ERR_MSG_3}
	}

	//item check
	if i := c.Items[index-1]; i == nil {
		return &test_model.ResponseInfo{Code: test_constant.RES_ERR, Msg: test_constant.RES_ERR_MSG_2}
	}

	//deal
	c.CapacityLeft += c.Items[index-1].Number
	c.Items[index-1] = nil
	return &test_model.ResponseInfo{Code: test_constant.RES_OK}
}

func (c *AllianceStoreInfo) ClearUp() {
	tempMap := make(map[int]int32) // {key: itemID  value: itemNum} 临时存放所有物品
	tempKeySlice := []int{}        //{itemID} 使用itemID有序遍历tempMap

	//将仓库内所有相同 itemID 的物品的数量总和放入 tempMap
	for k, v := range c.Items {
		if v == nil {
			continue
		}

		if _, ok := tempMap[int(v.ID)]; ok {
			tempMap[int(v.ID)] += v.Number
		} else {
			tempMap[int(v.ID)] = v.Number
			tempKeySlice = append(tempKeySlice, int(v.ID))
		}

		//清空仓库
		c.Items[k] = nil
	}

	//使用itemID 排序
	sort.Ints(tempKeySlice)

	//需要放入物品的格子索引
	index := 0
	//遍历所有物品,依次放入指定格子内,己满则顺延
	for _, key := range tempKeySlice {
		index = int(c.ClearUpPutItems(int32(key), int32(tempMap[key]), int32(index)))
	}
}

//将格子内物品按最大堆叠数依次放入指定格子内, 如果指定格子己满则顺延到下一个格子, 直到物品全部放置
func (c *AllianceStoreInfo) ClearUpPutItems(itemId, itemNum, index int32) int32 {
	if itemNum <= 0 {
		return index
	}

	setNum := int32(math.Min(float64(itemNum), test_constant.STORE_ITEM_STACK_MAX_NUM))
	c.Items[index] = &Item{ID: itemId, Number: setNum}

	return c.ClearUpPutItems(itemId, itemNum-setNum, index+1)
}
