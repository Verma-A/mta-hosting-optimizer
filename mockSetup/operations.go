package mocksetup

import "log"

func GetData(value int) []string {

	var result []string

	log.Println(">>>>>>data from get data: ", Data)

	countMap := make(map[string]int)
	for _, val := range Data {
		if _, ok := countMap[val.HostName]; !ok {
			countMap[val.HostName] = 1
		} else {
			countMap[val.HostName] = countMap[val.HostName] + 1
		}
	}

	log.Println(">>>>>>>count map: ", countMap)
	for key, val := range countMap {
		if val <= value {
			result = append(result, key)
		}
	}

	return result
}
