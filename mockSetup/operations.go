package mocksetup

func GetData(value int) []string {

	var result []string

	countMap := make(map[string]int)
	for _, val := range Data {
		if _, ok := countMap[val.HostName]; !ok && val.Active {
			countMap[val.HostName] = 1
		} else if _, ok := countMap[val.HostName]; !ok && !val.Active {
			countMap[val.HostName] = 0
		} else if val.Active {
			countMap[val.HostName] = countMap[val.HostName] + 1
		}
	}

	for key, val := range countMap {
		if val <= value {
			result = append(result, key)
		}
	}

	return result
}
