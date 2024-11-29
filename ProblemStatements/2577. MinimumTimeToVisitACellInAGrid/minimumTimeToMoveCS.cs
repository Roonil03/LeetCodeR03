public class Solution {
    public int MinimumTime(int[][] grid) {
        int rows = grid.Length, cols = grid[0].Length;
        int[][] dir = new int[][] { new int[] { -1, 0 }, new int[] { 1, 0 }, new int[] { 0, -1 }, new int[] { 0, 1 } };
        if (grid[0][1] > 1 && grid[1][0] > 1)
            return -1;
        var pq = new PriorityQueue<Node>();
        bool[,] visited = new bool[rows, cols];
        pq.Enqueue(new Node(0, 0, 0));
        visited[0, 0] = true;
        while (pq.Count > 0) {
            var n = pq.Dequeue();
            if (n.Row == rows - 1 && n.Col == cols - 1)
                return n.Time;
            foreach (var d in dir) {
                int newRow = n.Row + d[0];
                int newCol = n.Col + d[1];
                if (newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols && !visited[newRow, newCol]) {
                    int waitTime = (grid[newRow][newCol] - n.Time) % 2 == 0 ? 1 : 0;
                    int nextTime = (n.Time >= grid[newRow][newCol]) ? n.Time + 1 : grid[newRow][newCol] + waitTime;
                    pq.Enqueue(new Node(newRow, newCol, nextTime));
                    visited[newRow, newCol] = true;
                }
            }
        }
        return -1;
    }
}

public class Node : IComparable<Node> {
    public int Row { get; }
    public int Col { get; }
    public int Time { get; }
    public Node(int row, int col, int time) {
        Row = row;
        Col = col;
        Time = time;
    }
    public int CompareTo(Node other) {
        return Time.CompareTo(other.Time);
    }
}

public class PriorityQueue<T> where T : IComparable<T> {
    private List<T> heap = new List<T>();
    public int Count => heap.Count;
    public void Enqueue(T item) {
        heap.Add(item);
        HeapifyUp(heap.Count - 1);
    }

    public T Dequeue() {
        if (Count == 0)
            throw new InvalidOperationException("The priority queue is empty.");
        T top = heap[0];
        heap[0] = heap[heap.Count - 1];
        heap.RemoveAt(heap.Count - 1);
        if (Count > 0)
            HeapifyDown(0);
        return top;
    }

    private void HeapifyUp(int index) {
        while (index > 0) {
            int parent = (index - 1) / 2;
            if (heap[index].CompareTo(heap[parent]) >= 0)
                break;
            Swap(index, parent);
            index = parent;
        }
    }

    private void HeapifyDown(int index) {
        while (true) {
            int left = 2 * index + 1;
            int right = 2 * index + 2;
            int smallest = index;

            if (left < Count && heap[left].CompareTo(heap[smallest]) < 0)
                smallest = left;
            if (right < Count && heap[right].CompareTo(heap[smallest]) < 0)
                smallest = right;

            if (smallest == index)
                break;

            Swap(index, smallest);
            index = smallest;
        }
    }

    private void Swap(int a, int b) {
        T temp = heap[a];
        heap[a] = heap[b];
        heap[b] = temp;
    }
}
