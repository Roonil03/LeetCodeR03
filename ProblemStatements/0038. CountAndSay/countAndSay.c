#include <stdlib.h>
#include <string.h>


char* helper(char* c)
{
    if(c == NULL)
    {
        char* t = (char*)malloc(sizeof(char)*2);
        t[0] = '1';
        t[1] = '\0';
        return t;
    }
    int i = 0;
    int j = 0;
    int n = strlen(c);
    char* s = (char*)malloc((n * 2 + 1) * sizeof(char));
    while(i<n){
        int count = 1;
        while(i<n-1 &&  c[i]== c[i+1])
        {
            count++;
            i++;
        }

        s[j++] = count + '0';
        s[j++] = c[i++];
    }
    s[j] = '\0';
    //printf("1. %s",s);
    return s;
}
char* countAndSay(int n){
    char* str = helper(NULL);
    for(int i = 0; i < n-1; i++)
    {
        char* newStr = helper(str);
        free(str);
        str = newStr;
    }
    return str;
}