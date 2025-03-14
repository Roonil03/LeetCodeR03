#include <stdbool.h>

bool canChange(char* start, char* target) {
    int n = strlen(start);
    int i = 0, j = 0;

    while (i < n || j < n) {
        while (i < n && start[i] == '_')
        {
            i++;
        }
        while (j < n && target[j] == '_')
        {
            j++;
        }
        if (i == n && j == n)
        {
            return true;
        }
        if (i == n || j == n)
        {
            return false;
        }
        if (start[i] != target[j]) return false;
        if ((start[i] == 'L' && i < j) || (start[i] == 'R' && i > j))
        {
            return false;
        }
        i++;
        j++;
    }
    return true;
}