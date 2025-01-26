public class Solution {
    private List<int> Toposort(List<List<int>> graph) {
        var res = new List<int>();
        var found = new int[graph.Count];
        var stack = new Stack<int>();
        
        for (int i = 0; i < graph.Count; i++) {
            stack.Push(i);
        }
        
        while (stack.Count > 0) {
            int node = stack.Pop();
            if (node < 0) {
                res.Add(~node);
            }
            else if (found[node] == 0) {
                found[node] = 1;
                stack.Push(~node);
                foreach (var nei in graph[node]) {
                    stack.Push(nei);
                }
            }
        }
        
        foreach (var node in res) {
            if (graph[node].Any(nei => found[nei] == 1)) {
                return null;
            }
            found[node] = 0;
        }
        
        res.Reverse();
        return res;
    }
    
    private (List<int> res, int[] idx, bool valid) Kahn(List<List<int>> graph) {
        int n = graph.Count;
        var indeg = new int[n];
        var idx = new int[n];
        
        for (int i = 0; i < n; i++) {
            foreach (var e in graph[i]) {
                indeg[e]++;
            }
        }
        
        var q = new List<int>();
        var res = new List<int>();
        
        for (int i = 0; i < n; i++) {
            if (indeg[i] == 0) {
                q.Add(i);
            }
        }
        
        int nr = 0;
        while (q.Count > 0) {
            int node = q[q.Count - 1];
            q.RemoveAt(q.Count - 1);
            res.Add(node);
            idx[res[res.Count - 1]] = nr;
            nr++;
            
            foreach (var e in graph[res[res.Count - 1]]) {
                indeg[e]--;
                if (indeg[e] == 0) {
                    q.Add(e);
                }
            }
        }
        
        return (res, idx, nr == n);
    }
    
    public IList<IList<int>> Supersequences(IList<string> words) {
        var g = new List<List<int>>();
        var g2 = new List<List<int>>();
        
        for (int i = 0; i < 26; i++) {
            g.Add(new List<int>());
            g2.Add(new List<int>());
        }
        
        foreach (var word in words) {
            int x = word[0] - 'a';
            int y = word[1] - 'a';
            g[x].Add(y);
            g2[y].Add(x);
        }
        
        var L = new List<int>();
        for (int i = 0; i < 26; i++) {
            if (g[i].Count > 0 || g2[i].Count > 0) {
                L.Add(i);
            }
        }
        
        int m = L.Count;
        for (int k = 0; k <= m; k++) {
            var works = new List<HashSet<int>>();
            foreach (var doubled in GetCombinations(L, k)) {
                var D = new HashSet<int>(doubled);
                var tempG = new List<List<int>>();
                
                for (int i = 0; i < 26; i++) {
                    tempG.Add(new List<int>());
                }
                
                foreach (var i in L) {
                    if (!D.Contains(i)) {
                        foreach (var j in g[i]) {
                            if (!D.Contains(j)) {
                                tempG[i].Add(j);
                            }
                        }
                    }
                }
                
                if (Toposort(tempG) != null) {
                    works.Add(D);
                }
            }
            
            if (works.Count > 0) {
                var answer = new List<IList<int>>();
                foreach (var D in works) {
                    var entry = new int[26];
                    foreach (var i in L) {
                        entry[i] = 1;
                    }
                    foreach (var i in D) {
                        entry[i]++;
                    }
                    answer.Add(entry);
                }
                return answer;
            }
        }
        
        return new List<IList<int>>();
    }
    
    private IEnumerable<IList<int>> GetCombinations(List<int> arr, int k) {
        int n = arr.Count;
        var result = new List<IList<int>>();
        
        void Combine(int start, List<int> current) {
            if (current.Count == k) {
                result.Add(new List<int>(current));
                return;
            }            
            for (int i = start; i < n; i++) {
                current.Add(arr[i]);
                Combine(i + 1, current);
                current.RemoveAt(current.Count - 1);
            }
        }        
        Combine(0, new List<int>());
        return result;
    }
}

/*
The provided program leverages graph theory and topological sorting to solve a problem related to 
finding valid supersequences of given words. The program first constructs directed graphs (g and 
g2) representing dependencies between characters based on the input words. Two methods are 
implemented for topological sorting: a depth-first search based method (Toposort) and Kahn's 
algorithm (Kahn). The main function, Supersequences, generates all possible combinations of 
characters that can be doubled, and for each combination, it creates a temporary graph excluding 
those characters. It then checks if the temporary graph has a valid topological order, indicating 
that the supersequence can be formed. The function returns the valid combinations that allow 
forming supersequences. This approach combines combinatorial generation (using GetCombinations) 
with graph processing, ensuring that all constraints from the input words are respected while 
forming the supersequences.
*/