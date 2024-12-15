import java.util.*;

class Solution {
    public double maxAmount(String initialCurrency, List<List<String>> pairs1, double[] rates1, List<List<String>> pairs2, double[] rates2) {
        for (int i = 0; i < rates2.length; i++)
        {
            rates2[i] = 1.0 / rates2[i];
        }        
        Map<String, Double> m1 = calc(initialCurrency, pairs1, rates1);
        Map<String, Double> m2 = calc(initialCurrency, pairs2, rates2);        
        double ans = 1.0;        
        for (Map.Entry<String, Double> entry : m1.entrySet())
        {
            String currency = entry.getKey();
            if (m2.containsKey(currency)) {
                ans = Math.max(ans, m1.get(currency) * m2.get(currency));
            }
        }        
        return ans;
    }
    
    private Map<String, Double> calc(String s0, List<List<String>> pairs, double[] rates) {
        Map<String, List<Pair<String, Double>>> e = new HashMap<>();
        int n = pairs.size();        
        for (int i = 0; i < n; i++)
        {
            String s1 = pairs.get(i).get(0), s2 = pairs.get(i).get(1);
            e.computeIfAbsent(s1, k -> new ArrayList<>()).add(new Pair<>(s2, rates[i]));
            e.computeIfAbsent(s2, k -> new ArrayList<>()).add(new Pair<>(s1, 1.0 / rates[i]));
        }        
        Queue<String> q = new LinkedList<>();
        Map<String, Double> ret = new HashMap<>();
        q.offer(s0);
        ret.put(s0, 1.0);        
        while (!q.isEmpty())
        {
            String now = q.poll();
            for (Pair<String, Double> t : e.getOrDefault(now, new ArrayList<>()))
            {
                String nxt = t.getKey();
                double r = t.getValue();
                if (ret.containsKey(nxt)) continue;
                ret.put(nxt, ret.get(now) * r);
                q.offer(nxt);
            }
        }        
        return ret;
    }
    
    private static class Pair<T, U>
    {
        private final T key;
        private final U value;        
        public Pair(T key, U value)
        {
            this.key = key;
            this.value = value;
        }        
        public T getKey() 
        {
            return key;
        }        
        public U getValue()
        {
            return value;
        }
    }
}
