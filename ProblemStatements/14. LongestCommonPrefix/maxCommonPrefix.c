char* longestCommonPrefix(char** strs, int strsSize) {
    char* str = (char*)malloc(200 * sizeof(char));
    int count = 0, condition = 1;
    
    while (condition != 0) {
        char temp = strs[0][count];
        condition = 0;
        
        for (int i = 0; i < strsSize; i++) {
            if (count >= strlen(strs[i]) || strs[i][count] != temp) {
                condition = 0;
                str[count] = '\0';
                return str;
            }
            condition = 1;
        }
        str[count] = temp;
        count++;
    }
    str[count] = '\0';
    return str;
}