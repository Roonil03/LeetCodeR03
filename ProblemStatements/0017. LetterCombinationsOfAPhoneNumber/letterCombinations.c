/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char* mapping[] = {
    "", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"
};

void backtrack(char* digits, int index, char* current, char** result, int* returnSize) {
    if (index == strlen(digits)) {
        result[*returnSize] = strdup(current);
        (*returnSize)++;
        return;
    }
    
    int digit = digits[index] - '0';
    char* letters = mapping[digit];
    
    for (int i = 0; letters[i] != '\0'; i++) {
        current[index] = letters[i];
        backtrack(digits, index + 1, current, result, returnSize);
    }
}

char** letterCombinations(char* digits, int* returnSize) {
    if (digits == NULL || strlen(digits) == 0) {
        *returnSize = 0;
        return NULL;
    }

    int maxCombinations = 1;
    for (int i = 0; digits[i] != '\0'; i++) {
        maxCombinations *= strlen(mapping[digits[i] - '0']);
    }

    char** result = (char**)malloc(maxCombinations * sizeof(char*));
    char* current = (char*)malloc((strlen(digits) + 1) * sizeof(char));
    current[strlen(digits)] = '\0';  // Null-terminate the string
    *returnSize = 0;

    backtrack(digits, 0, current, result, returnSize);

    free(current);
    return result;
}