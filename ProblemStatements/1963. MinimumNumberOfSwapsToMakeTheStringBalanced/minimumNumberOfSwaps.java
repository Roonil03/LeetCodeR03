class Solution {
    public int minSwaps(String s) {
    //     int count = 0;
    //     int left = 0, right  = s.length() - 1;
    //     StringBuilder sb = new StringBuilder(s);
    //     while(!isBalanced(s))
    //     {
    //         left = s.indexOf(']',left);
    //         right = s.indexOf('[',right);
    //         if(left != -1 && right != -1){
    //         sb.setCharAt(left, '[');
    //         sb.setCharAt(right, ']');            
    //         count++;
    //         }
    //         s = sb.toString();
    //         System.out.println(s);
    //     }
    //     return count;
    // }
    // boolean isBalanced(String s){
    //     Stack<Character> stack = new Stack<>();
    //     for(int i = 0; i<s.length();i++)
    //     {
    //         if(s.charAt(i)  == '[')
    //         {
    //             stack.push('[');
    //         }
    //         else if(!stack.isEmpty() && s.charAt(i) == ']'){
    //             stack.pop();
    //         }
    //         else{
    //             return false;
    //         }
    //     }
    //     if(stack.isEmpty()){
    //         return true;
    //     }
    //     return false;
    // }
        int open = 0;
        for(int i = 0; i< s.length();i++)
        {
            if(s.charAt(i) == '[')
            {
                open++;
            }
            else{
                if(open>0){
                open--;
                }
            }
        }
        return (open+1)/2;
    }
}