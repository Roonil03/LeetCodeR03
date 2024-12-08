import java.util.*; 

class Solution {
    public int maxRectangleArea(int[][] points) {
        if (points.length < 4) 
            return -1;        
        HashSet<String> psHash = new HashSet<>();
        for (int[] point : points) {
            psHash.add(point[0] + "," + point[1]);
        }        
        int area = -1;
        for (int i = 0; i < points.length; i++) {
            for (int j = i + 1; j < points.length; j++) {
                int x1 = points[i][0], y1 = points[i][1];
                int x2 = points[j][0], y2 = points[j][1];                
                if (x1 != x2 && y1 != y2) {
                    String c1 = x1 + "," + y2;
                    String c2 = x2 + "," + y1;                    
                    if (psHash.contains(c1) && psHash.contains(c2)) {
                        if (!isValid(x1, y1, x2, y2, psHash)) {
                            int a = Math.abs(x2 - x1) * Math.abs(y2 - y1);
                            area = Math.max(area, a);
                        }
                    }
                }
            }
        }
        return area;
    }

    private boolean isValid(int x1, int y1, int x2, int y2, HashSet<String> psHash) {
        int minX = Math.min(x1, x2), maxX = Math.max(x1, x2);
        int minY = Math.min(y1, y2), maxY = Math.max(y1, y2);        
        for (int x = minX; x <= maxX; x++) {
            for (int y = minY; y <= maxY; y++) {
                String point = x + "," + y;
                if (psHash.contains(point) && !(x == x1 && y == y1) && !(x == x2 && y == y2) && !(x == x1 && y == y2) && !(x == x2 && y == y1)) 
                {
                    return true;
                }
            }
        }
        return false;
    }
}
