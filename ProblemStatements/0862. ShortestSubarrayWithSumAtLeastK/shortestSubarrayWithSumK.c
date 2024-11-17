typedef struct {
    long val;
    int idx;
} Node;

typedef struct {
    Node* arr;
    int front;
    int rear;
} MonoQ;

void initQ(MonoQ* q, int size) {
    q->arr = (Node*)malloc(size * sizeof(Node));
    q->front = 0;
    q->rear = 0;
}

void push(MonoQ* q, int size, Node n) {
    while (q->front != q->rear && q->arr[(q->rear - 1 + size) % size].val >= n.val) {
        q->rear = (q->rear - 1 + size) % size;
    }
    q->arr[q->rear] = n;
    q->rear = (q->rear + 1) % size;
}

void pop(MonoQ* q, int size) {
    q->front = (q->front + 1) % size;
}

Node front(MonoQ* q) {
    return q->arr[q->front];
}

int shortestSubarray(int* nums, int numsSize, int k) {
    long* pre = (long*)malloc((numsSize + 1) * sizeof(long));
    pre[0] = 0;

    for (int i = 0; i < numsSize; i++) {
        pre[i + 1] = pre[i] + nums[i];
        if (nums[i] >= k) {
            free(pre);
            return 1;
        }
    }    
    MonoQ q;
    initQ(&q, numsSize + 1);
    int res = numsSize + 1;
    for (int i = 0; i <= numsSize; i++) {
        while (q.front != q.rear && pre[i] - front(&q).val >= k) {
            res = (i - front(&q).idx) < res ? (i - front(&q).idx) : res;
            pop(&q, numsSize + 1);
        }
        push(&q, numsSize + 1, (Node){pre[i], i});
        while (q.front != q.rear && front(&q).idx < i - res + 1) {
            pop(&q, numsSize + 1);
        }
    }    
    free(pre);
    free(q.arr);
    return res == numsSize + 1 ? -1 : res;
}
