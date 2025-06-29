class Solution {
public:
    bool lemonadeChange(vector<int>& bills) {
        int i5{0}, i10{0};
        for(int i : bills){
            if ( i == 5){
                i5++;
            } else if(i == 10){
                if(i5 == 0){
                    return false;
                }
                i5--;
                i10++;
            } else{
                if(i10 > 0 && i5 > 0){
                    i10--;
                    i5--;
                } else if(i5 >= 3){
                    i5 -= 3;
                } else{
                    return false;
                }
            }
        }
        return true;
    }
};