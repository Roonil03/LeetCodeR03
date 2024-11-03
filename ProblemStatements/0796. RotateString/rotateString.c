#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

char* shift(char* str)
{
    char* res = (char*)malloc(sizeof(char)*(strlen(str)+1));
    for(int i = 0; i<strlen(str)-1;i++)
    {
        res[i] = str[i+1];
    }
    res[strlen(str)-1] = str[0];
    res[strlen(str)] = '\0';
    printf("%s\n",res);
    return res;
}
bool rotateString(char* s, char* goal) {
    char* temp = s;
    if (strcmp(s,goal) == 0) return true;
    s = shift(s);
    while(strcmp(s,temp) != 0)
    {
        if(strcmp(s,goal) == 0) return true;
        s = shift(s);
    }
    return false;
}