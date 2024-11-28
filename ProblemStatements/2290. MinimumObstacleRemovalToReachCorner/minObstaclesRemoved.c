#include <stdlib.h>
#include <limits.h>
#include <stdbool.h>

typedef struct {
    int x, y, cost;
} Node;

typedef struct {
    Node* data;
    int size;
    int capacity;
} MinHeap;

MinHeap* createHeap(int capacity) {
    MinHeap* heap = (MinHeap*)malloc(sizeof(MinHeap));
    heap->data = (Node*)malloc(sizeof(Node) * capacity);
    heap->size = 0;
    heap->capacity = capacity;
    return heap;
}

void swap(Node* a, Node* b) {
    Node temp = *a;
    *a = *b;
    *b = temp;
}

void push(MinHeap* heap, Node node) {
    heap->data[heap->size++] = node;
    int i = heap->size - 1;
    while (i > 0 && heap->data[(i - 1) / 2].cost > heap->data[i].cost) {
        swap(&heap->data[i], &heap->data[(i - 1) / 2]);
        i = (i - 1) / 2;
    }
}

Node pop(MinHeap* heap) {
    Node root = heap->data[0];
    heap->data[0] = heap->data[--heap->size];
    int i = 0;
    while (true) {
        int smallest = i;
        int left = 2 * i + 1, right = 2 * i + 2;
        if (left < heap->size && heap->data[left].cost < heap->data[smallest].cost) smallest = left;
        if (right < heap->size && heap->data[right].cost < heap->data[smallest].cost) smallest = right;
        if (smallest == i) break;
        swap(&heap->data[i], &heap->data[smallest]);
        i = smallest;
    }
    return root;
}

bool isEmpty(MinHeap* heap) {
    return heap->size == 0;
}

int minimumObstacles(int** grid, int gridSize, int* gridColSize) {
    int m = gridSize, n = gridColSize[0];
    int directions[4][2] = {{0, 1}, {1, 0}, {0, -1}, {-1, 0}};
    int** cost = (int**)malloc(sizeof(int*) * m);
    for (int i = 0; i < m; i++) {
        cost[i] = (int*)malloc(sizeof(int) * n);
        for (int j = 0; j < n; j++) cost[i][j] = INT_MAX;
    }
    MinHeap* heap = createHeap(m * n);
    push(heap, (Node){0, 0, 0});
    cost[0][0] = 0;
    while (!isEmpty(heap)) {
        Node current = pop(heap);
        if (current.x == m - 1 && current.y == n - 1) {
            int result = current.cost;
            for (int i = 0; i < m; i++) free(cost[i]);
            free(cost);
            free(heap->data);
            free(heap);
            return result;
        }
        for (int i = 0; i < 4; i++) {
            int nx = current.x + directions[i][0];
            int ny = current.y + directions[i][1];
            if (nx >= 0 && nx < m && ny >= 0 && ny < n) {
                int newCost = current.cost + grid[nx][ny];
                if (newCost < cost[nx][ny]) {
                    cost[nx][ny] = newCost;
                    push(heap, (Node){nx, ny, newCost});
                }
            }
        }
    }
    for (int i = 0; i < m; i++) free(cost[i]);
    free(cost);
    free(heap->data);
    free(heap);
    return -1;
}