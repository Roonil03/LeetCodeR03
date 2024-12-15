/* Could not solve it entirely during the time of the contest, 
 * it has been solved post the time limit of the contest.
 */

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

/*
1. Currency Exchange Rate Inversion:
    The code first inverts the exchange rates in rates2 by dividing 1.0 by each rate. This is because the rates2 array initially represents the cost to buy a unit of the second currency in terms of the first currency.

2. Currency Exchange Rate Calculation (calc function):
    Graph Representation:
        The code constructs a graph-like structure using a HashMap (e) where keys are currencies, and values are lists of Pair objects. Each Pair represents an exchange possibility between two currencies, along with the corresponding exchange rate.
    Breadth-First Search (BFS):
        It performs a BFS starting from the initialCurrency.
        During BFS:
            It maintains a Map (ret) to store the maximum achievable amount of each currency from the initialCurrency.
            It explores all possible exchange paths, updating the ret map with the maximum achievable amount for each reachable currency.

3. Maximum Profit Calculation:
    Calculate Exchange Rates for Both Sets of Pairs:
        The calc function is called twice:
            Once with pairs1 and rates1 to get the maximum achievable amounts for each currency using the first set of exchange pairs.
            Once with pairs2 and the inverted rates2 to get the maximum achievable amounts for each currency using the second set of exchange pairs.
    Find Maximum Profit:
        The code iterates through the currencies found in the first set of exchanges.
        For each currency, if it also exists in the second set of exchanges, it calculates the product of the achievable amounts from both sets.
        It keeps track of the maximum profit found among all such currencies.

4. Return Maximum Profit:
    The code returns the calculated maximum profit.

In essence, the algorithm calculates the maximum achievable amount of each currency using two different sets of exchange rates and then finds the maximum profit by combining the results from both sets.
 */