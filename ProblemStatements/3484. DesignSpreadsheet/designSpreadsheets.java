class Spreadsheet {
    Map<String, Integer> mat;

    public Spreadsheet(int rows) {
        this.mat = new HashMap<>();
    }
    
    public void setCell(String cell, int value) {
        mat.put(cell, value);
    }
    
    public void resetCell(String cell) {
        mat.remove(cell);
    }
    
    public int getValue(String formula) {
        int sum = 0;
        String[] ops = formula.substring(1).split("\\+");
        for(String i : ops){
            if(Character.isDigit(i.charAt(0))){
                sum += Integer.parseInt(i);
            } else{
                sum += mat.getOrDefault(i, 0)
                ;
            }
        }
        return sum;
    }
}

/**
 * Your Spreadsheet object will be instantiated and called as such:
 * Spreadsheet obj = new Spreadsheet(rows);
 * obj.setCell(cell,value);
 * obj.resetCell(cell);
 * int param_3 = obj.getValue(formula);
 */