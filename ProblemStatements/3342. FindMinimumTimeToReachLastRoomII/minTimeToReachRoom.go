// func minTimeToReach(moveTime [][]int) int {
//     n, m := len(moveTime), len(moveTime[0])
// 	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
// 	visited := make([][]bool, n)
// 	for i := range visited {
// 		visited[i] = make([]bool, m)
// 	}
//     pq := &PriorityQueue{}
// 	heap.Init(pq)
// 	heap.Push(pq, Node{0, 0, 0})
// 	for pq.Len() > 0 {
// 		node := heap.Pop(pq).(Node)
// 		x, y, currTime := node.x, node.y, node.time

// 		if x == n-1 && y == m-1 {
// 			return currTime
// 		}
// 		if visited[x][y] {
// 			continue
// 		}
// 		visited[x][y] = true
// 		for _, dir := range directions {
// 			newX, newY := x+dir[0], y+dir[1]
// 			if newX >= 0 && newX < n && newY >= 0 && newY < m && !visited[newX][newY] {
// 				waitTime := 0
//                 if currTime < moveTime[newX][newY] {
// 					waitTime = moveTime[newX][newY] - currTime
// 				}
//                 nextMoveTime := 1
// 				if (currTime+waitTime)%2 == 1 {
// 					nextMoveTime = 2
// 				}
// 				heap.Push(pq, Node{newX, newY, currTime + waitTime + nextMoveTime})
// 			}
// 		}
// 	}
//     return -1
// }

// type PriorityQueue []Node

// type Node struct {
// 	x, y, time int
// }

// func (pq PriorityQueue) Len() int { return len(pq) }

// func (pq PriorityQueue) Less(i, j int) bool {
// 	return pq[i].time < pq[j].time
// }

// func (pq PriorityQueue) Swap(i, j int) {
// 	pq[i], pq[j] = pq[j], pq[i]
// }

// func (pq *PriorityQueue) Push(x interface{}) {
// 	*pq = append(*pq, x.(Node))
// }

// func (pq *PriorityQueue) Pop() interface{} {
// 	old := *pq
// 	n := len(old)
// 	x := old[n-1]
// 	*pq = old[:n-1]
// 	return x
// }

var UP = [2]int{-1, 0}
var DOWN = [2]int{1, 0}
var LEFT = [2]int{0, -1}
var RIGHT = [2]int{0, 1}
var MOVES = [4][2]int{UP, DOWN, LEFT, RIGHT}

var rows int
var columns int

var startRow int
var startColumn int

var targetRow int
var targetColumn int

func minTimeToReach(moveTime [][]int) int {
	rows = len(moveTime)
	columns = len(moveTime[0])

	startRow = 0
	startColumn = 0

	targetRow = rows - 1
	targetColumn = columns - 1

	return dijkstraSearchForPathWithMinTime(moveTime)
}

func dijkstraSearchForPathWithMinTime(moveTime [][]int) int {
	minHeapForTime := make(PriorityQueue, 0)
	step := NewStep(startRow, startColumn, 0, 0)
	heap.Push(&minHeapForTime, step)

	minTimeMatrix := make([][]int, rows)
	for row := 0; row < rows; row++ {
		minTimeMatrix[row] = make([]int, columns)
		for column := 0; column < columns; column++ {
			minTimeMatrix[row][column] = math.MaxInt
		}
	}
	minTimeMatrix[startRow][startColumn] = 0

	for len(minHeapForTime) > 0 {
		current := heap.Pop(&minHeapForTime).(*Step)
		if current.row == targetRow && current.column == targetColumn {
			break
		}

		for _, move := range MOVES {
			nextRow := current.row + move[0]
			nextColumn := current.column + move[1]
			if !isInMatrix(nextRow, nextColumn) {
				continue
			}

			timeToMoveToNextPoint := getNewTimeToMoveBetweenTwoPoints(current.timeToMoveBetweenTwoPoints)
			nextValueForTime := max(timeToMoveToNextPoint+current.timeFromStart,
				timeToMoveToNextPoint+moveTime[nextRow][nextColumn])

			if minTimeMatrix[nextRow][nextColumn] > nextValueForTime {
				minTimeMatrix[nextRow][nextColumn] = nextValueForTime
				step := NewStep(nextRow, nextColumn, timeToMoveToNextPoint, nextValueForTime)
				heap.Push(&minHeapForTime, step)
			}
		}
	}

	return minTimeMatrix[targetRow][targetColumn]
}

func isInMatrix(row int, column int) bool {
	return row >= 0 && row < rows && column >= 0 && column < columns
}

func getNewTimeToMoveBetweenTwoPoints(prviousTimeToMoveBetweenTwoPoints int) int {
	return 1 + (prviousTimeToMoveBetweenTwoPoints % 2)
}

type Step struct {
	row                        int
	column                     int
	timeToMoveBetweenTwoPoints int
	timeFromStart              int
}

func NewStep(row int, column int, timeToMoveBetweenTwoPoints int, timeFromStart int) *Step {
	step := &Step{
		row:                        row,
		column:                     column,
		timeToMoveBetweenTwoPoints: timeToMoveBetweenTwoPoints,
		timeFromStart:              timeFromStart,
	}
	return step
}

type PriorityQueue []*Step

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(first int, second int) bool {
	return pq[first].timeFromStart < pq[second].timeFromStart
}

func (pq PriorityQueue) Swap(first int, second int) {
	pq[first], pq[second] = pq[second], pq[first]
}

func (pq *PriorityQueue) Push(object any) {
	step := object.(*Step)
	*pq = append(*pq, step)
}

func (pq *PriorityQueue) Pop() any {
	step := (*pq)[len(*pq)-1]
	(*pq)[len(*pq)-1] = nil
	*pq = (*pq)[0 : len(*pq)-1]
	return step
}