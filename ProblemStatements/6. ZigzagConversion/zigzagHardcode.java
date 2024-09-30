class Solution {
    public String convert(String s, int numRows) {
        char[][] mat = new char[numRows][s.length()];
        if(numRows==1 || numRows == s.length())
        {
            return s;
        }
        int count = 0;
        int i = 0;
        int j = 0;
        while(count < s.length())
        {
            if(i == numRows-1){
                while(i!=0 && count < s.length())
                {
                    mat[i--][j++] = s.charAt(count);
                    count++;
                }
                continue;
            }
            mat[i++][j] = s.charAt(count);
            count++;
        }
        String n = "";
        for(int a = 0; a<numRows;a++)
        {
            for(int b = 0; b<s.length();b++)
            {
                if(mat[a][b] != 0)
                {
                    n = n + Character.toString(mat[a][b]);
                }
            }
        }
        return n;
    }
}