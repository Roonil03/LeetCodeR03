#include <stdlib.h>
#include <string.h>
#include <stdbool.h>

int* findSubstring(char* s, char** words, int wordsSize, int* returnSize) {
    int sLen = strlen(s);
    int wordLen = strlen(words[0]);
    int concatLen = wordLen * wordsSize;
    int* result = (int*)malloc(sizeof(int) * sLen);
    *returnSize = 0;
    if (sLen < concatLen) {
        return result;
    }
    int* wordCount = (int*)calloc(5000, sizeof(int));
    int* slidingWindowCount = (int*)calloc(5000, sizeof(int));
    int uniqueWordsCount = 0;
    char* uniqueWords[wordsSize];    
    for (int i = 0; i < wordsSize; i++) {
        bool found = false;
        for (int j = 0; j < uniqueWordsCount; j++) {
            if (strcmp(uniqueWords[j], words[i]) == 0) {
                wordCount[j]++;
                found = true;
                break;
            }
        }
        if (!found) {
            uniqueWords[uniqueWordsCount] = words[i];
            wordCount[uniqueWordsCount]++;
            uniqueWordsCount++;
        }
    }
    for (int i = 0; i < wordLen; i++) {
        memset(slidingWindowCount, 0, sizeof(int) * uniqueWordsCount);
        
        int start = i, matchedWords = 0;
        
        for (int j = i; j <= sLen - wordLen; j += wordLen) {
            char* sub = strndup(s + j, wordLen);
            
            int index = -1;
            for (int k = 0; k < uniqueWordsCount; k++) {
                if (strcmp(uniqueWords[k], sub) == 0) {
                    index = k;
                    break;
                }
            }
            
            if (index >= 0) { 
                slidingWindowCount[index]++;                
                if (slidingWindowCount[index] <= wordCount[index]) {
                    matchedWords++;
                } else {
                    while (slidingWindowCount[index] > wordCount[index]) {
                        char* leftSub = strndup(s + start, wordLen);
                        int leftIndex = -1;
                        for (int k = 0; k < uniqueWordsCount; k++) {
                            if (strcmp(uniqueWords[k], leftSub) == 0) {
                                leftIndex = k;
                                break;
                            }
                        }
                        free(leftSub);                        
                        slidingWindowCount[leftIndex]--;
                        if (slidingWindowCount[leftIndex] < wordCount[leftIndex]) {
                            matchedWords--;
                        }
                        start += wordLen;
                    }
                }                
                if (matchedWords == wordsSize) {
                    result[(*returnSize)++] = start;
                }
            } else {
                memset(slidingWindowCount, 0, sizeof(int) * uniqueWordsCount);
                matchedWords = 0;
                start = j + wordLen;
            }
            free(sub);
        }
    }    
    free(wordCount);
    free(slidingWindowCount);
    return result;
}
