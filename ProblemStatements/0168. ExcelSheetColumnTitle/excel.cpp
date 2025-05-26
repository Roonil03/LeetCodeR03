class Solution {
public:
    string convertToTitle(int columnNumber) {
        // string res = ""
        // string hash[] = {"Z","A","B","C","D","E","F","G","H","I","J","K","L","M","N","O","P", "Q","R","S","T","U","V","W","X","Y"};
        // while(columnNumber > 0){
        //     res = hash[columnNumber % 26]

        // }
        return columnNumber == 0 ? "" : convertToTitle((columnNumber - 1) / 26) + (char) ((columnNumber - 1) % 26 + 'A');
    }
};