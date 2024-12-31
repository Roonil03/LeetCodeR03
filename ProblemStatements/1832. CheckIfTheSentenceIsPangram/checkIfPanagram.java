class Solution {
    public boolean checkIfPangram(String sentence) {
        boolean h[] = new boolean[26];
        int len = sentence.length();
        // int count = 0;
        for(int i = 0; i< len; i++){
            ///h[sentence.charAt(i) - 'a'] = true;
            int id = sentence.charAt(i) - 'a';
            if(!h[id]){
                h[id] = true;
                count++;
                if(count == 26){
                    return true;
                }
            }
        }
        // for(int i = 0; i < 26; i++){
        //     if(!h[i]){
        //         return false;
        //     }
        // }
        return false;
    }
}