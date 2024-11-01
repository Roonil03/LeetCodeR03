#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char* makeFancyString(char* s) {
    char* res = (char*)malloc(strlen(s) + 1);
    int a = 1, b = 1; 
    res[0] = s[0];
    while (s[b] != '\0') {
        if (!(a > 1 && s[b] == res[a - 1] && s[b] == res[a - 2])) {
            res[a++] = s[b];
        }
        b++;
    }
    res[a] = '\0';    
    return res;
}