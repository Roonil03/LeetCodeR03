class Solution {
    static final int M = 100000;
    static int[] d = new int[M];
    static List<Integer>[] g = new ArrayList[M];

    static {
        for (int i = 0; i < M; i++) {
            g[i] = new ArrayList<>();
        }
    }

    static int diam(int[][] e) {
        int n = e.length + 1;
        Arrays.fill(d, 0, n, 0);
        for (int i = 0; i < n; i++) {
            g[i].clear();
        }
        for (int[] ed : e) {
            int u = ed[0], v = ed[1];
            g[u].add(v);
            g[v].add(u);
            d[u]++;
            d[v]++;
        }
        int[] q = new int[M];
        int f = 0, b = 0;
        for (int i = 0; i < n; i++) {
            if (d[i] == 1) {
                q[b++] = i;
            }
        }
        int lvl = 0, rem = n;
        while (rem > 2) {
            int sz = b - f;
            rem -= sz;
            for (int i = 0; i < sz; i++) {
                int u = q[f++];
                for (int v : g[u]) {
                    if (--d[v] == 1) {
                        q[b++] = v;
                    }
                }
            }
            lvl++;
        }
        return rem == 2 ? 2 * lvl + 1 : 2 * lvl;
    }

    public static int minimumDiameterAfterMerge(int[][] e1, int[][] e2) {
        int d1 = diam(e1), d2 = diam(e2);
        return Math.max(Math.max(d1, d2), (d1 + 1) / 2 + (d2 + 1) / 2 + 1);
    }
}
