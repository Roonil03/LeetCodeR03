class Solution {
public:
    int computeArea(int ax1, int ay1, int ax2, int ay2, int bx1, int by1, int bx2, int by2) {
        int a1 = abs((ax2 - ax1) * (ay2 - ay1));
        int a2 = abs((bx2 - bx1) * (by2 - by1));
        int l = max(ax1, bx1);
        int r = min(ax2, bx2);
        int b = max(ay1, by1);
        int t = min(ay2, by2);
        int temp {0};
        if(r > l && t > b){
            temp = (r - l) * (t - b);
        }
        return a1 + a2 - temp;
    }
};