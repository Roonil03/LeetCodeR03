class Solution {
public:
    int numRookCaptures(vector<vector<char>>& board) {
        int x {-1}, y{-1};
        int res {0};
        for(int i {0}; i < 8; i++){
            for(int j {0}; j < 8; j++){
                if(board[i][j] == 'R'){
                    x = i;
                    y = j;
                    break;
                }
            }
            if(x != -1){
                break;
            }
        }
        int dirs[4][2] = {{-1, 0}, {1, 0}, {0, -1}, {0, 1}};
        for(int d {0}; d < 4; d++){
            int nx = x + dirs[d][0];
            int ny = y + dirs[d][1];
            while(nx >= 0 && nx < 8 && ny >= 0 && ny < 8){
                if(board[nx][ny] == 'B'){
                    break;
                }
                if(board[nx][ny] == 'p'){
                    res++;
                    break;
                }
                nx += dirs[d][0];
                ny += dirs[d][1];
            }
        }
        return res;
    }
};