#include <limits.h>
#include <stdio.h>
#include <stdlib.h>

int myAtoi(char* s) {
    char sign = '+';
    int num = 0, i=0;
    while(s[i] == ' ')
    {
        i++;
    }
    if(s[i] == '-' || s[i] == '+')
    {
        sign = s[i]=='-'?'-':'+';
        i++;
    }
    while(isdigit(s[i])&&i<strlen(s))
    {
        if (num > (INT_MAX - (s[i] - '0')) / 10) {
            return (sign == '+') ? INT_MAX : INT_MIN;
        }
        num *= 10;
        num += (s[i] - '0');
        i++;
    }
    return sign == '-'?-num:num;
}