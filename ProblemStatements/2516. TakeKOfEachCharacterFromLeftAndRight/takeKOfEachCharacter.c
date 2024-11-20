int min(int a, int b) {
    return a < b ? a : b;
}

int takeCharacters(char* s, int k) {
    int count[3] = {0};
    int n = strlen(s);
    for (int i = 0; i < n; i++) {
        count[s[i] - 'a']++;
    }
    if (min(count[0], min(count[1], count[2])) < k) {
        return -1;
    }
    int res = n + 1, l = 0;
    for (int r = 0; r < n; r++) {
        count[s[r] - 'a']--;
        
        while (min(count[0], min(count[1], count[2])) < k) {
            count[s[l] - 'a']++;
            l++;
        }        
        res = min(res, n - (r - l + 1));
    }
    return res;
}
