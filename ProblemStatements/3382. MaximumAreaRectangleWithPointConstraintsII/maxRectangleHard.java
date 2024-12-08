import java.util.*;

class Solution {
    public long maxRectangleArea(int[] xCoord, int[] yCoord) {
        int numPoints = xCoord.length;
        int[][] points = new int[numPoints][];
        int[] sortedY = generateIndexMapping(yCoord);
        for (int i = 0; i < numPoints; i++) {
            points[i] = new int[]{xCoord[i], Arrays.binarySearch(sortedY, yCoord[i])};
        }
        Arrays.sort(points, (a, b) -> {
            if (a[0] != b[0]) return a[0] - b[0];
            return a[1] - b[1];
        });
        Map<Long, Integer> rectangleCountMap = new HashMap<>();
        Map<Long, Integer> xCoordinateMap = new HashMap<>();
        long maxArea = -1;
        int[] fenwickTree = new int[sortedY.length + 1];
        for (int i = 0; i < points.length; i++) {
            updateFenwickTree(fenwickTree, points[i][1], 1);
            if (i - 1 >= 0 && points[i][0] == points[i - 1][0]) {
                long yCoordinatesPair = (long) points[i][1] << 32 | points[i - 1][1];
                int areaFrequency = sumFenwickTree(fenwickTree, points[i][1]) - sumFenwickTree(fenwickTree, points[i - 1][1] - 1);
                if (rectangleCountMap.containsKey(yCoordinatesPair)) {
                    int previousFrequency = rectangleCountMap.get(yCoordinatesPair);
                    if (areaFrequency == previousFrequency + 2) {
                        int xCoordinate = xCoordinateMap.get(yCoordinatesPair);
                        long area = (long) (points[i][0] - xCoordinate) * (sortedY[points[i][1]] - sortedY[points[i - 1][1]]);
                        maxArea = Math.max(maxArea, area);
                    }
                }
                rectangleCountMap.put(yCoordinatesPair, areaFrequency);
                xCoordinateMap.put(yCoordinatesPair, points[i][0]);
            }
        }
        return maxArea;
    }

    public static int sumFenwickTree(int[] fenwickTree, int index) {
        int sum = 0;
        for (index++; index > 0; index -= index & -index) sum += fenwickTree[index];
        return sum;
    }

    public static void updateFenwickTree(int[] fenwickTree, int index, int value) {
        if (value == 0 || index < 0) return;
        int size = fenwickTree.length;
        for (index++; index < size; index += index & -index) fenwickTree[index] += value;
    }

    public static int[] generateIndexMapping(int[] array) {
        int[] indexMapping = Arrays.copyOf(array, array.length);
        Arrays.sort(indexMapping);
        int position = 1;
        for (int i = 1; i < indexMapping.length; i++) {
            if (indexMapping[i] != indexMapping[position - 1]) indexMapping[position++] = indexMapping[i];
        }
        return Arrays.copyOf(indexMapping, position);
    }
}
