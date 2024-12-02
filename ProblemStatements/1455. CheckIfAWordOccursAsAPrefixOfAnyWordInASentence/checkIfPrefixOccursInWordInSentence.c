int isPrefixOfWord(char* sentence, char* searchWord) {
    int count = 1, i = 0, len = strlen(sentence), lenSW = strlen(searchWord);
    while (i < len) {
        if (sentence[i] == ' ') {
            i++;
            count++;
            continue;
        }
        int j = 0, start = i;
        while (i < len && j < lenSW && sentence[i] == searchWord[j]) {
            i++;
            j++;
        }
        if (j == lenSW) {
            return count;
        }
        i = start;
        while (i < len && sentence[i] != ' ') {
            i++;
        }
    }
    return -1;
}
