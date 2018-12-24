package netsort

import (
	"ForumPublica/sde/static"
	"ForumPublica/server/models"
)

//CalcNode sorting node
type CalcNode struct {
	Color    string
	Children []int32
}

//CalcNetwork sorting nodes hash
type CalcNetwork map[int32]CalcNode

//ArrangeBPO sort BPO by manufacturing dependens
func ArrangeBPO(cnBpos *models.ConstructionBpos) []int32 {
	bpoIds := make([]int32, 0)
	for _, bpo := range *cnBpos {
		bpoIds = append(bpoIds, bpo.TypeID)
	}
	network := createNetwork(bpoIds)
	sorted := sortNetwork(network)
	return sorted
}

func createNetwork(bpos []int32) CalcNetwork {
	result := make(CalcNetwork)
	for _, bpoID := range bpos {
		deepAddNetworkElements(bpoID, result)
	}
	return result
}

func deepAddNetworkElements(key int32, result CalcNetwork) {
	_, exists := result[key]
	if !exists {
		keys := static.Level1BPOIds(key)

		result[key] = CalcNode{
			Color:    "W",
			Children: keys,
		}
		for _, childKey := range keys {
			deepAddNetworkElements(childKey, result)
		}
	}
}

func sortNetwork(network CalcNetwork) []int32 {
	result := make([]int32, 0)

	for key := range network {
		deepSortScan(&result, network, key)
	}

	return result
}

func deepSortScan(result *[]int32, network CalcNetwork, key int32) {
	node := network[key]
	if node.Color == "W" {
		node.Color = "G"
		network[key] = node

		for _, childBpoID := range node.Children {
			deepSortScan(result, network, childBpoID)
		}

		node.Color = "B"
		network[key] = node

		temp := append([]int32{key}, *result...)
		*result = temp
	}
}
