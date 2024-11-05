int minChanges(char* s) {
    int res = 0;
    for(int i = 0 ;i<strlen(s);)
    {
        char temp = s[i];
        int count = 0;
        while(i<strlen(s) && temp == s[i])
        {
            count++;
            i++;
        }
        if(count%2 != 0)
        {
            res++;
            i++;
        }
    }
    return res;
}