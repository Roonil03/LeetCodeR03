#include <stdio.h>
#include <string.h>
#include <stdlib.h>

int valueOf(char ch)
{
    switch(ch)
    {
        case 'I':
        return 1;
        case 'V':
        return 5;
        case 'X':
        return 10;
        case 'L':
        return 50;
        case 'C':
        return 100;
        case 'D':
        return 500;
        case 'M':
        return 1000;
        case 'Q':
        return 4;
        case 'W':
        return 9;
        case 'E':
        return 40;
        case 'R':
        return 90;
        case 'T':
        return 400;
        case 'Y':
        return 900;
    }
}
int romanToInt(char* s) {
    int sum = 0;
    int len = strlen(s);
    for(int i = 0; i<len;i++)
    {
        if(s[i] == 'I' && (i+1)<len && s[i+1] == 'V')
        {
            i++;
            sum += valueOf('Q');
            continue;
        }
        if(s[i] == 'I' && (i+1)<len && s[i+1] == 'X')
        {
            i++;
            sum += valueOf('W');
            continue;
        }
        if(s[i] == 'X' && (i+1)<len && s[i+1] == 'L')
        {
            i++;
            sum += valueOf('E');
            continue;
        }
        if(s[i] == 'X' && (i+1)<len && s[i+1] == 'C')
        {
            i++;
            sum += valueOf('R');
            continue;
        }
        if(s[i] == 'C' && (i+1)<len && s[i+1] == 'D')
        {
            i++;
            sum += valueOf('T');
            continue;
        }
        if(s[i] == 'C' && (i+1)<len && s[i+1] == 'M')
        {
            i++;
            sum += valueOf('Y');
            continue;
        }
        sum += valueOf(s[i]);
    }
    return sum;
}

// int main()
// {
//     printf("%d", romanToInt("MCMXCIV"));
//     return 0;
// }