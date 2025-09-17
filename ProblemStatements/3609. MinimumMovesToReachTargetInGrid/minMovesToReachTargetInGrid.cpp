class Solution {
public:
    int minMoves(int sx, int sy, int tx, int ty) {
        int cnt = 0;
        for (; tx > sx || ty > sy; ++cnt) {
            if (ty > tx) {
                swap(tx, ty);
                swap(sx, sy);
            }
            if (tx == ty) {
                if (sx == 0)
                    tx = 0;
                else
                    ty = 0;                
            }
            else {
                if (tx < ty * 2)
                    tx -= ty;
                else {
                    if (tx % 2)
                        return -1;
                    tx /= 2;
                }
            }
        }
        return tx < sx || ty < sy ? -1 : cnt;
    }
};

// func minMoves(sx int, sy int, tx int, ty int) int {
//     if tx < sx || ty < sy {
//         return -1
//     }
//     moves := 0
//     for tx > sx || ty > sy {
// 		if tx < sx || ty < sy {
// 			return -1
// 		}
//         if tx > ty {
// 			if ty == 0 {
// 				return -1
// 			}
// 			if tx > 2*ty {
// 				if tx%2 != 0 {
// 					return -1
// 				}
// 				tx /= 2
// 			} else {
// 				tx -= ty
// 			}
// 		} else if ty > tx {
// 			if tx == 0 {
// 				return -1
// 			}
// 			if ty > 2*tx {
// 				if ty%2 != 0 {
// 					return -1
// 				}
// 				ty /= 2
// 			} else {
// 				ty -= tx
// 			}
// 		} else {
// 			return -1
// 		}
// 		moves++
// 	}
// 	return moves
// }

