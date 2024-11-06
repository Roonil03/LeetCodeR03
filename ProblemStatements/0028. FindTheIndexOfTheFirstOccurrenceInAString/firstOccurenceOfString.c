#include <stdio.h>
#include <string.h>

int strStr(char* haystack, char* needle) {
    int haystackLen = strlen(haystack);
    int needleLen = strlen(needle);
    if (needleLen == 0) {
        return 0;
    }
    if (needleLen > haystackLen) {
        return -1;
    }
    for (int i = 0; i <= haystackLen - needleLen; i++) {
        int j = 0;
        while (j < needleLen && haystack[i + j] == needle[j]) {
            j++;
        }
        if (j == needleLen) {
            return i;
        }
    }
    return -1;
}