func numTilePossibilities(tiles string) int {
	id := 0
	arr := make([]int, 26)
	for i := 0; i < len(tiles); i++ {
		arr[tiles[i]-'A']++
	}
	bt(arr, &id)
	return id
}

func bt(arr []int, id *int) {
	for i := 0; i < 26; i++ {
		if arr[i] > 0 {
			arr[i]--
			(*id)++
			bt(arr, id)
			arr[i]++
		}
	}
}