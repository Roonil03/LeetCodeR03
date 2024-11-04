#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char* compressedString(char* word) {
    int len = strlen(word);
    char* comp = (char*)malloc(2*len+1);
    int a = 0;    
    for (int i = 0; i < len;) {
        char c = word[i];
        int count = 0;        
        while (i < len && word[i] == c && count < 9) {
            count++;
            i++;
        }        
        comp[a++] = count + '0';
        comp[a++] = c;
    }    
    comp[a] = '\0';
    return comp;
}
