class Solution {
    public String convert(String s, int numRows) {
        if (numRows == 1 || numRows >= s.length()) {
            return s;
        }
        
        StringBuilder sb = new StringBuilder();
        int cycleLength = 2 * numRows - 2; // Length of one full zigzag cycle
        
        for (int row = 0; row < numRows; row++) {
            for (int j = 0; j + row < s.length(); j += cycleLength) {
                // Append the character for the vertical part of the zigzag
                sb.append(s.charAt(j + row));
                // Append the character for the diagonal part of the zigzag
                if (row > 0 && row < numRows - 1 && j + cycleLength - row < s.length()) {
                    sb.append(s.charAt(j + cycleLength - row));
                }
            }
        }
        
        return sb.toString();
    }
}