class Solution {
    public List<Integer> getRow(int rowIndex) {
        List<Integer> ls = new ArrayList<>();
        long cur = 1;
        ls.add((int)cur);
        for(int i = 1; i<= rowIndex; i++)
        {
            cur = cur * (rowIndex - i + 1) / i;
            ls.add((int) cur);
        }
        return ls;
    }
}