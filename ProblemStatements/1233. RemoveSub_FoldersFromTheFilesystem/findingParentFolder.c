/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int compare(const void* a, const void* b) {
    return strcmp(*(char**)a, *(char**)b);
}
char** removeSubfolders(char** folder, int folderSize, int* returnSize) {
    qsort(folder, folderSize, sizeof(char*), compare);
    char** result = (char**)malloc(folderSize * sizeof(char*));
    int index = 0;
    result[index++] = folder[0];
    for (int i = 1; i < folderSize; i++) {
        int len = strlen(result[index - 1]);
         if (strncmp(result[index - 1], folder[i], len) != 0 || folder[i][len] != '/') {
            result[index++] = folder[i]; 
        }
    }
    *returnSize = index;
    return result;
}     
