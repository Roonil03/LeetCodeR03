class Solution {
    public String minimumString(String a, String b, String c) {
        String[] candidates = {
            mergeStrings(mergeStrings(a, b), c),
            mergeStrings(mergeStrings(a, c), b),
            mergeStrings(mergeStrings(b, a), c),
            mergeStrings(mergeStrings(b, c), a),
            mergeStrings(mergeStrings(c, a), b),
            mergeStrings(mergeStrings(c, b), a)
        };
        Arrays.sort(candidates, new Comparator<String>() {
            public int compare(String s1, String s2) {
                if (s1.length() != s2.length()) {
                    return Integer.compare(s1.length(), s2.length());
                }
                return s1.compareTo(s2);
            }
        });
        return candidates[0];
    }
    static String mergeStrings(String a, String b){
        if (a.contains(b)){
            return a;
        }
        if (b.contains(a)){
            return b;
        }
        for (int i = Math.min(a.length(), b.length()); i > 0; i--){
            if (a.endsWith(b.substring(0, i))){
                return a + b.substring(i);
            }
        }
        return a + b;
    }
}