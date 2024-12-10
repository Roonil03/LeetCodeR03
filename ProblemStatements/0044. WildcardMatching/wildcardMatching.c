# include <stdbool.h>

bool isMatch(char *s, char *p) {
    int sLen = strlen(s), pLen = strlen(p);
    int i = 0, j = 0, star = -1, match = 0;
    while (i < sLen) {
        if (j < pLen && (p[j] == '?' || p[j] == s[i])) {
            i++;
            j++;
        } else if (j < pLen && p[j] == '*') {
            star = j++;
            match = i;
        } else if (star != -1) {
            j = star + 1;
            i = ++match;
        } else {
            return false;
        }
    }
    while (j < pLen && p[j] == '*') {
        j++;
    }
    return j == pLen;
}