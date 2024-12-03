char* addSpaces(char* s, int* spaces, int spacesSize) {    
    int j = 0;
    int k = 0;
    int i = 0;
    int n = strlen(s);
    char* str = (char*)malloc(sizeof(char)*(spacesSize + n + 1));
    while(i < n){
        if(k < spacesSize && i == spaces[k])
        {
            k++;
            str[j++] = ' ';
        }
        str[j++] = s[i++];
    }
    str[j] = '\0';
    return str;
}