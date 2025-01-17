bool doesValidArrayExist(int* derived, int n) {
    unsigned ans = 0;
    for (int i = 0; i < n; i++) {
        ans ^= derived[i];
    }
    return ans == 0;
}