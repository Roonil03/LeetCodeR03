#include <limits.h>

#define min(x, y) (((x) < (y)) ? (x) : (y))

long long maxMatrixSum(int** matrix, int matrixSize, int* matrixColSize) {
    long long sum = 0;
    int neg = 0;
    int minAbsVal = INT_MAX;
    for(int i = 0; i< matrixSize; i++)
    {
        for(int j = 0; j< matrixColSize[i]; j++)
        {
            sum += abs(matrix[i][j]);
            if(matrix[i][j]<0)
            {
                neg++;
            }
            minAbsVal = min(minAbsVal, abs(matrix[i][j]));
        }
    }
    if(neg % 2!=0){
        sum -= 2*minAbsVal;
    }
    return sum;
}