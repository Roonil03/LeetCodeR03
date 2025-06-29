/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * int guess(int num);
 */

class Solution {
public:
    int guessNumber(int n) {
        int l {0}, r = n;
        while(l <= r){
            long long m = l + (r - l) / 2;
            int res = guess(m);
            switch(res){
                case 0:
                return (int)m;
                case 1:
                l = m + 1;
                break;
                case -1:
                r = m - 1;
                break;
                default:
                return -1;
            }
        }
        return -1;
    }
};