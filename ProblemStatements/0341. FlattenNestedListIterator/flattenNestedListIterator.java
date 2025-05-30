/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * public interface NestedInteger {
 *
 *     // @return true if this NestedInteger holds a single integer, rather than a nested list.
 *     public boolean isInteger();
 *
 *     // @return the single integer that this NestedInteger holds, if it holds a single integer
 *     // Return null if this NestedInteger holds a nested list
 *     public Integer getInteger();
 *
 *     // @return the nested list that this NestedInteger holds, if it holds a nested list
 *     // Return empty list if this NestedInteger holds a single integer
 *     public List<NestedInteger> getList();
 * }
 */
public class NestedIterator implements Iterator<Integer> {
    List<Integer> f;
    int id;
    
    public NestedIterator(List<NestedInteger> nestedList) {
        f = new ArrayList<>();
        id = 0;
        f = flatten(nestedList);    
    }

    List<Integer> flatten(List<NestedInteger> nested){
        List<Integer> res = new ArrayList<>();
        for (NestedInteger ni : nested){
            if (ni.isInteger()){
                res.add(ni.getInteger());
            } else{
                res.addAll(flatten(ni.getList()));
            }
        }
        return res;
    }

    @Override
    public Integer next() {
        return f.get(id++);
    }

    @Override
    public boolean hasNext() {
        return id < f.size();
    }
}

/**
 * Your NestedIterator object will be instantiated and called as such:
 * NestedIterator i = new NestedIterator(nestedList);
 * while (i.hasNext()) v[f()] = i.next();
 */