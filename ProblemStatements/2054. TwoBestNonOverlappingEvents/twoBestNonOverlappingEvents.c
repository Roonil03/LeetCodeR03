typedef struct {
    int start, end, value;
} Event;

int compare(const void *a, const void *b) {
    Event *eventA = (Event *)a;
    Event *eventB = (Event *)b;
    return eventA->end - eventB->end;
}

int binarySearch(Event *events, int n, int target) {
    int low = 0, high = n - 1, result = -1;
    while (low <= high) {
        int mid = low + (high - low) / 2;
        if (events[mid].end < target) {
            result = mid;
            low = mid + 1;
        } else {
            high = mid - 1;
        }
    }
    return result;
}

int maxTwoEvents(int** events, int eventsSize, int* eventsColSize) {
    Event *eventList = (Event *)malloc(eventsSize * sizeof(Event));
    for (int i = 0; i < eventsSize; i++) {
        eventList[i].start = events[i][0];
        eventList[i].end = events[i][1];
        eventList[i].value = events[i][2];
    }
    qsort(eventList, eventsSize, sizeof(Event), compare);
    int *maxValues = (int *)malloc((eventsSize + 1) * sizeof(int));
    maxValues[0] = 0;
    for (int i = 0; i < eventsSize; i++)
        maxValues[i + 1] = (eventList[i].value > maxValues[i]) ? eventList[i].value : maxValues[i];
    int result = 0;
    for (int i = 0; i < eventsSize; i++) {
        int prevIndex = binarySearch(eventList, eventsSize, eventList[i].start);
        int currentMax = eventList[i].value + (prevIndex >= 0 ? maxValues[prevIndex + 1] : 0);
        if (currentMax > result)
            result = currentMax;
    }
    free(eventList);
    free(maxValues);
    return result;
}
