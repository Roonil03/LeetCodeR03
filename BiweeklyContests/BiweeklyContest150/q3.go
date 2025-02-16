// Could not solve during the contest time, has been solved post contest time

type Event struct {
	y, x1, x2 int64
	typ       int
}

type Slice struct {
	yStart, yEnd, coveredX int64
}

type SegmentTree struct {
	size       int
	coverCount []int64
	coveredLen []int64
	coords     []int64
}

func NewSegmentTree(xvals []int64) *SegmentTree {
	st := &SegmentTree{
		coords: xvals,
		size:   len(xvals) - 1,
	}
	st.coverCount = make([]int64, st.size*4)
	st.coveredLen = make([]int64, st.size*4)
	return st
}

func (st *SegmentTree) update(x1, x2 int64, val int) {
	if x1 >= x2 {
		return
	}
	L := st.indexOf(x1)
	R := st.indexOf(x2) - 1
	if L <= R {
		st.updateRange(1, 0, st.size-1, L, R, val)
	}
}

func (st *SegmentTree) indexOf(x int64) int {
	lo, hi := 0, len(st.coords)-1
	for lo < hi {
		mid := (lo + hi) / 2
		if st.coords[mid] < x {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

func (st *SegmentTree) updateRange(node, start, end, L, R int, val int) {
	if R < start || L > end {
		return
	}
	if L <= start && end <= R {
		st.coverCount[node] += int64(val)
	} else {
		mid := (start + end) / 2
		st.updateRange(node*2, start, mid, L, R, val)
		st.updateRange(node*2+1, mid+1, end, L, R, val)
	}
	if st.coverCount[node] > 0 {
		st.coveredLen[node] = st.coords[end+1] - st.coords[start]
	} else {
		if start == end {
			st.coveredLen[node] = 0
		} else {
			st.coveredLen[node] = st.coveredLen[node*2] + st.coveredLen[node*2+1]
		}
	}
}

func (st *SegmentTree) getCoveredLength() int64 {
	return st.coveredLen[1]
}

func uniqueLongs(arr []int64) []int64 {
	if len(arr) == 0 {
		return arr
	}
	res := []int64{arr[0]}
	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[i-1] {
			res = append(res, arr[i])
		}
	}
	return res
}

func separateSquares(squares [][]int) float64 {
	luntrivexi := squares
	var events []Event
	var xvals []int64
	for _, sq := range luntrivexi {
		x := int64(sq[0])
		y := int64(sq[1])
		l := int64(sq[2])
		x1 := x
		x2 := x + l
		y1 := y
		y2 := y + l
		events = append(events, Event{y: y1, x1: x1, x2: x2, typ: 1})
		events = append(events, Event{y: y2, x1: x1, x2: x2, typ: -1})
		xvals = append(xvals, x1, x2)
	}
	sort.Slice(xvals, func(i, j int) bool {
		return xvals[i] < xvals[j]
	})
	xvals = uniqueLongs(xvals)
	sort.Slice(events, func(i, j int) bool {
		return events[i].y < events[j].y
	})
	st := NewSegmentTree(xvals)
	var slices []Slice
	totalArea := 0.0
	if len(events) == 0 {
		return 0.0
	}
	currentY := events[0].y
	st.update(events[0].x1, events[0].x2, events[0].typ)
	for i := 0; i < len(events)-1; i++ {
		nextY := events[i+1].y
		coverageX := st.getCoveredLength()
		deltaY := float64(nextY - currentY)
		areaStrip := float64(coverageX) * deltaY
		slices = append(slices, Slice{yStart: currentY, yEnd: nextY, coveredX: coverageX})
		totalArea += areaStrip
		currentY = nextY
		st.update(events[i+1].x1, events[i+1].x2, events[i+1].typ)
	}
	half := totalArea / 2.0
	partial := 0.0
	for _, sl := range slices {
		stripHeight := float64(sl.yEnd - sl.yStart)
		stripArea := float64(sl.coveredX) * stripHeight
		if partial+stripArea >= half && stripArea > 0 {
			needed := half - partial
			someHeight := needed / float64(sl.coveredX)
			boundaryY := float64(sl.yStart) + someHeight
			return boundaryY
		}
		partial += stripArea
	}
	if len(slices) > 0 {
		return float64(slices[len(slices)-1].yEnd)
	}
	return 0.0
}

/*
This code uses a sweep line algorithm combined with a segment tree and coordinate compression to
efficiently calculate the union area of overlapping squares. The events (representing the start
and end boundaries of squares along the y-axis) are processed in sorted order, while the segment
tree dynamically maintains the total covered length along the x-axis as the sweep line moves. By
updating the tree with interval start and end events, the algorithm computes the union length of
x-intervals quickly, allowing it to integrate the area strip-by-strip and ultimately find the
y-coordinate that splits the total area in half.
*/

