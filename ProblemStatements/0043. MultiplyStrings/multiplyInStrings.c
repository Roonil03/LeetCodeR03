char* multiply(char* num1, char* num2) {
    int len1 = strlen(num1);
    int len2 = strlen(num2);
    if (len1 == 0 || len2 == 0 || (len1 == 1 && num1[0] == '0') || (len2 == 1 && num2[0] == '0')) {
        char* zero = (char*)malloc(2);
        zero[0] = '0';
        zero[1] = '\0';
        return zero;
    }
    int* result = (int*)calloc(len1 + len2, sizeof(int));
    for (int i = len1 - 1; i >= 0; i--) {
        for (int j = len2 - 1; j >= 0; j--) {
            int mul = (num1[i] - '0') * (num2[j] - '0');
            int sum = mul + result[i + j + 1];
            result[i + j + 1] = sum % 10;
            result[i + j] += sum / 10;
        }
    }
    int index = 0;
    while (index < len1 + len2 && result[index] == 0) {
        index++;
    }
    char* res = (char*)malloc(len1 + len2 - index + 1);
    int k = 0;
    while (index < len1 + len2) {
        res[k++] = result[index++] + '0';
    }
    res[k] = '\0';
    free(result);
    return res;
}