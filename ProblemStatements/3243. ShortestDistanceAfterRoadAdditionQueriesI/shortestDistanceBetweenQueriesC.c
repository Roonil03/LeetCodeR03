/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
#include <stdbool.h>

typedef struct Queue {
    int *data;
    int front, rear, size, capacity;
} Queue;

typedef struct List {
    int *data;
    int size, capacity;
} List;

typedef struct Graph {
    List *adjList;
    int nodeCount;
} Graph;

Queue* createQueue(int capacity) {
    Queue *queue = (Queue*)malloc(sizeof(Queue));
    queue->data = (int*)malloc(capacity * sizeof(int));
    queue->front = queue->size = 0;
    queue->rear = capacity - 1;
    queue->capacity = capacity;
    return queue;
}

bool isQueueEmpty(Queue *queue) {
    return queue->size == 0;
}

void enqueue(Queue *queue, int value) {
    queue->rear = (queue->rear + 1) % queue->capacity;
    queue->data[queue->rear] = value;
    queue->size++;
}

int dequeue(Queue *queue) {
    int value = queue->data[queue->front];
    queue->front = (queue->front + 1) % queue->capacity;
    queue->size--;
    return value;
}

List* createList(int capacity) {
    List *list = (List*)malloc(sizeof(List));
    list->data = (int*)malloc(capacity * sizeof(int));
    list->size = 0;
    list->capacity = capacity;
    return list;
}

void addToList(List *list, int value) {
    if (list->size == list->capacity) {
        list->capacity *= 2;
        list->data = (int*)realloc(list->data, list->capacity * sizeof(int));
    }
    list->data[list->size++] = value;
}

Graph* createGraph(int nodeCount) {
    Graph *graph = (Graph*)malloc(sizeof(Graph));
    graph->nodeCount = nodeCount;
    graph->adjList = (List*)malloc(nodeCount * sizeof(List));
    for (int i = 0; i < nodeCount; i++) {
        graph->adjList[i] = *createList(2);
    }
    return graph;
}

int bfs(Graph *graph, int n) {
    bool *visited = (bool*)calloc(n, sizeof(bool));
    Queue *queue = createQueue(n);
    int currLayer = 1, nextLayer = 0, layers = 0;
    enqueue(queue, 0);
    visited[0] = true;
    while (!isQueueEmpty(queue)) {
        for (int i = 0; i < currLayer; i++) {
            int currNode = dequeue(queue);
            if (currNode == n - 1) {
                free(queue->data);
                free(queue);
                free(visited);
                return layers;
            }
            List neighbors = graph->adjList[currNode];
            for (int j = 0; j < neighbors.size; j++) {
                int neighbor = neighbors.data[j];
                if (visited[neighbor]) continue;
                enqueue(queue, neighbor);
                visited[neighbor] = true;
                nextLayer++;
            }
        }
        currLayer = nextLayer;
        nextLayer = 0;
        layers++;
    }
    free(queue->data);
    free(queue);
    free(visited);
    return -1;
}

int* shortestDistanceAfterQueries(int n, int** queries, int queriesSize, int* queriesColSize, int* returnSize) {
    Graph *graph = createGraph(n);
    int *answer = (int*)malloc(queriesSize * sizeof(int));
    *returnSize = queriesSize;
    for (int i = 0; i < n - 1; i++) {
        addToList(&graph->adjList[i], i + 1);
    }
    for (int i = 0; i < queriesSize; i++) {
        int u = queries[i][0];
        int v = queries[i][1];
        addToList(&graph->adjList[u], v);
        answer[i] = bfs(graph, n);
    }
    for (int i = 0; i < n; i++) {
        free(graph->adjList[i].data);
    }
    free(graph->adjList);
    free(graph);
    return answer;
}
