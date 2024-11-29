#include <stdbool.h>


typedef struct {
    int row, col, time;
} Node;

typedef struct {
    Node* h;
    int size;
    int cap;
} MinHeap;

void swap(Node* a, Node* b) {
    Node temp = *a;
    *a = *b;
    *b = temp;
}

void heapifyUp(MinHeap* h, int index) {
    while (index > 0) {
        int parent = (index - 1) / 2;
        if (h->h[parent].time <= h->h[index].time)
            break;
        swap(&h->h[parent], &h->h[index]);
        index = parent;
    }
}

void heapifyDown(MinHeap* h, int index) {
    int left, right, smallest;
    while (true) {
        left = 2 * index + 1;
        right = 2 * index + 2;
        smallest = index;
        if (left < h->size && h->h[left].time < h->h[smallest].time)
            smallest = left;
        if (right < h->size && h->h[right].time < h->h[smallest].time)
            smallest = right;
        if (smallest == index)
            break;
        swap(&h->h[smallest], &h->h[index]);
        index = smallest;
    }
}

MinHeap* createMinHeap(int cap) {
    MinHeap* h = (MinHeap*)malloc(sizeof(MinHeap));
    h->h = (Node*)malloc(sizeof(Node) * cap);
    h->size = 0;
    h->cap = cap;
    return h;
}

void push(MinHeap* h, Node n) {
    if (h->size == h->cap)
        return;
    h->h[h->size++] = n;
    heapifyUp(h, h->size - 1);
}

Node pop(MinHeap* h) {
    Node top = h->h[0];
    h->h[0] = h->h[--h->size];
    heapifyDown(h, 0);
    return top;
}

bool isEmpty(MinHeap* h) {
    return h->size == 0;
}

int minimumTime(int** grid, int gridSize, int* gridColSize) {
    int dir[4][2] = {{-1, 0}, {1, 0}, {0, -1}, {0, 1}};
    int rows = gridSize, cols = gridColSize[0];
    if (grid[0][1] > 1 && grid[1][0] > 1)
        return -1;
    MinHeap* h = createMinHeap(rows * cols);
    bool** visited = (bool**)malloc(rows * sizeof(bool*));
    for (int i = 0; i < rows; i++) {
        visited[i] = (bool*)calloc(cols, sizeof(bool));
    }
    push(h, (Node){0, 0, 0});
    visited[0][0] = true;
    while (!isEmpty(h)) {
        Node n = pop(h);
        if (n.row == rows - 1 && n.col == cols - 1)
            return n.time;
        for (int i = 0; i < 4; i++) {
            int newRow = n.row + dir[i][0];
            int newCol = n.col + dir[i][1];

            if (newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols && !visited[newRow][newCol]) {
                int waitTime = (grid[newRow][newCol] - n.time) % 2 == 0 ? 1 : 0;
                int nextTime = (n.time >= grid[newRow][newCol]) ? n.time + 1 : grid[newRow][newCol] + waitTime;
                push(h, (Node){newRow, newCol, nextTime});
                visited[newRow][newCol] = true;
            }
        }
    }
    for (int i = 0; i < rows; i++) {
        free(visited[i]);
    }
    free(visited);
    free(h->h);
    free(h);
    return -1;
}
