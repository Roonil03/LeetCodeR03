class Solution {
    public boolean isNumber(String s) {
        int len = s.length();
        if (len == 0) {
            return false;
        }
        boolean hasNum = false;
        boolean hasDot = false;
        boolean hasExpo = false;
        for (int i = 0; i < len; i++) {
            char c = s.charAt(i);
            if (Character.isDigit(c)) {
                hasNum = true;
            }
            else if (c == '+' || c == '-') {
                if (i > 0 && (s.charAt(i - 1) != 'e' && s.charAt(i - 1) != 'E')) {
                    return false;
                }
            } 
            else if (c == '.') {
                if (hasDot || hasExpo) {
                    return false;
                }
                hasDot = true;
            } 
            else if (c == 'e' || c == 'E') {
                if (hasExpo || !hasNum) {
                    return false;
                }
                hasExpo = true;
                hasNum = false; 
                } 
                else {
                return false;
            }
        }
        return hasNum;
    }
}