package netsort

import (
	"ForumPublica/sde/static"
	"ForumPublica/server/models"
)

type CalcNode struct {
	Color    string
	Children []int32
}

type CalcNetwork map[int32]CalcNode

func ArrangeBPO(cnBpos models.CnBlueprints) []int32 {
	bpos := make([]int32, 0)
	for _, bpo := range cnBpos {
		bpos = append(bpos, bpo.Model.TypeId)
	}
	network := createNetwork(bpos)
	sorted := sortNetwork(network)
	return sorted
}

func createNetwork(bpos []int32) CalcNetwork {
	result := make(CalcNetwork)
	for _, bpoId := range bpos {
		deepAddNetworkElements(bpoId, result)
	}
	return result
}

func deepAddNetworkElements(key int32, result CalcNetwork) {
	_, exists := result[key]
	if !exists {
		bpo := static.Blueprints[key]

		keys := make([]int32, 0)
		for _, id := range bpo.L1BPOIds() {
			keys = append(keys, GetTreeKey(id, 10, 20))
		}

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

	for key, _ := range network {
		deepSortScan(&result, network, key)
	}

	return result
}

func deepSortScan(result *[]int32, network CalcNetwork, key int32) {
	node := network[key]
	if node.Color == "W" {
		node.Color = "G"
		network[key] = node

		for _, childBpoId := range node.Children {
			deepSortScan(result, network, childBpoId)
		}

		node.Color = "B"
		network[key] = node

		temp := append([]string{key}, *result...)
		*result = temp
	}
}
