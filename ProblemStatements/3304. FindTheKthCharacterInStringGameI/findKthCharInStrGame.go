func kthCharacter(k int) byte {
	pos, lvl, l := k, 0, 1
	for l < k {
		l *= 2
		lvl++
	}
	for lvl > 0 {
		h := l / 2
		if pos > h {
			pos = pos - h
		}
		l = h
		lvl--
	}
	// ch := byte('a')
	pos, lvl, l = k, 0, 1
	for l < k {
		l *= 2
		lvl++
	}
	res, l1, p, ln := 0, lvl, pos, l
	for l1 > 0 {
		h := ln / 2
		if p > h {
			res++
			p = p - h
		}
		ln = h
		l1--
	}
	return byte('a' + (res % 26))
}