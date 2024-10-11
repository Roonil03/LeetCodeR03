#include <stdlib.h>

struct Guest {
    int start, end;
    short index;
};

int compareGuests(const void* a, const void* b) {
    return ((struct Guest*)a)->start - ((struct Guest*)b)->start;
}

int smallestChair(int** timeData, int dataSize, int* timeColSize, int targetGuest) {
    struct Guest *guests = (struct Guest*)malloc(sizeof(struct Guest) * dataSize);
    short k;

    for(k = 0; k < dataSize; k++) {
        guests[k].index = k;
        guests[k].start = timeData[k][0];
        guests[k].end = timeData[k][1];
    }

    qsort(guests, dataSize, sizeof(struct Guest), compareGuests);

    int l, currentGuest, arrivalTime, totalSeats = 0;
    int *seatStatus = (int*)malloc(sizeof(int) * dataSize);
    
    for(k = 0; k < dataSize; k++) {
        arrivalTime = guests[k].start;

        for(l = 0; l < totalSeats; l++) {
            if(seatStatus[l] <= arrivalTime) {
                seatStatus[l] = guests[k].end;
                break;
            }
        }

        if(l == totalSeats) seatStatus[totalSeats++] = guests[k].end;
        if(guests[k].index == targetGuest) break;
    }

    free(guests);
    free(seatStatus);
    return l;
}
