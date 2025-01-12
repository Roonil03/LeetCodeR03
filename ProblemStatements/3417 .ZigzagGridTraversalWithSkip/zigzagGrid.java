class Solution {
    public List<Integer> zigzagTraversal(int[][] grid) {
        List<Integer> res = new ArrayList<>();
        boolean gr = true;
        boolean sc = false;
        for(int i = 0; i < grid.length; i++)
        {
            if(gr){
                for(int j = 0; j < grid[0].length; j++)
                {
                    if(!sc){
                        res.add(grid[i][j]);
                    }
                    sc = !sc;
                }
            }
            else{
                for(int j = grid[0].length - 1; j >=0; j--)
                {
                    if(!sc){
                        res.add(grid[i][j]);
                    }
                    sc = !sc;
                }
            }
            gr = !gr;
        }
        return res;
    }
}