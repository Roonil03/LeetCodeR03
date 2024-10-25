/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
#include <stdio.h>
#include <stdlib.h>

char** fizzBuzz(int n, int* returnSize) {
    //char* n1 = (char*)malloc(sizeof(char)*8);
    char** returnStr = (char**)malloc(sizeof(char*)*n);
    *returnSize = n;
    for(int i = 1; i<=n;i++)
    {
        if(i%15==0)
        {
            returnStr[i - 1] = (char*)malloc(9);
            returnStr[i-1] = "FizzBuzz";
        }
        else if(i%5 == 0)
        {
            returnStr[i - 1] = (char*)malloc(5);
            returnStr[i-1] = "Buzz";
        }
        else if(i%3 == 0)
        {
            returnStr[i - 1] = (char*)malloc(5);
            returnStr[i-1] = "Fizz";
        }
        else{
            returnStr[i - 1] = (char*)malloc(12);
            sprintf(returnStr[i - 1], "%d", i);
        }
    }
    return returnStr;
}