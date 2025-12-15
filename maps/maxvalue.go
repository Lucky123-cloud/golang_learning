//get the max value of a map
func MaxValueKeys(m map[string]int) []string {
    // If map is empty, return empty slice
    if len(m) == 0 {
        return []string{}
    }

    maxVal := 0
    first := true

    // 1. Find the maximum value
    for _, v := range m {
        if first || v > maxVal {
            maxVal = v
            first = false
        }
    }

    // 2. Collect all keys that have the maximum value
    result := []string{}
    for k, v := range m {
        if v == maxVal {
            result = append(result, k)
        }
    }

    return result
}

