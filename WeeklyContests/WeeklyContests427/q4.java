/* Could not solve it entirely during the time of the contest, 
* it has been solved post the time limit of the contest.
*/

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

/*
 * This code finds the maximum area of a rectangle that can be formed from a set of points on a 2D plane, where the rectangle's edges are parallel to the axes and no other points lie inside or on the border of the rectangle. Here's a breakdown of how it works:

 1. Coordinate Preprocessing (generateIndexMapping)
The generateIndexMapping function creates a sorted version of the yCoord array and then compresses the yCoord values by mapping them to their positions in the sorted array. This helps reduce the range of possible y-coordinates and speeds up the Fenwick Tree operations.

2. Sorting Points
The points array stores each point as a pair [xCoord[i], yCoord[i]]. It is then sorted primarily by the x-coordinate and secondarily by the y-coordinate. This helps in efficiently finding and processing vertical pairs of points later.

3. Fenwick Tree (Binary Indexed Tree)
A Fenwick Tree is used to efficiently manage and update the counts of y-coordinates. The fenwickTree array helps in determining how many points exist between two y-coordinates in logarithmic time. This is critical to avoid brute-force checks for every point within a rectangle.

4. Iterating Over Point Pairs
For each point, we consider it and the previous point with the same x-coordinate (i.e., points that are vertically aligned). If these two points form the diagonal of a rectangle, we check if the other two corners of the rectangle (the other two points) exist in the point set.
The rectangleCountMap stores how many times a specific pair of y-coordinates (representing the opposite corners of a rectangle) has been encountered, and xCoordinateMap stores the x-coordinate of the first occurrence of that pair.

5. Rectangle Area Calculation
If the opposite corners of a rectangle are found, the code calculates its area using the formula:
Area = ( 𝑥 2 − 𝑥 1 ) × ( 𝑦 2 − 𝑦 1 ) 
    Area=(x2−x1)×(y2−y1)
The area is updated to the maximum found during the iterations.

6. Returning the Result
The function returns the maximum area found, or -1 if no rectangle could be formed.

Key Concepts:
Efficient Rectangle Check: Instead of checking every point inside the potential rectangle, the code only checks whether the opposite corners exist, reducing the complexity.
Fenwick Tree: This data structure is used for efficient prefix sum queries and updates, helping to manage the vertical segments of rectangles.
 */