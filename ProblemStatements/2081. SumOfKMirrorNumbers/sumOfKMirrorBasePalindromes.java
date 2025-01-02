class Solution {
    public long kMirror(int k, int n) {
        long sum = 0;
        int count = 0;
        for (int len = 1; count < n; len++) {
            List<Long> m = generateMirrors(len);
            for (long num : m) {
                if (isPalindrome(Long.toString(num, k))) {
                    sum += num;
                    count++;
                    if (count == n) {
                        break;
                    }
                }
            }
        }
        return sum;
    }

    List<Long> generateMirrors(int l) {
        List<Long> res = new ArrayList<>();
        if (l == 1) {
            for (int i = 1; i <= 9; i++) {
                res.add((long) i);
            }
        } else {
            backtrack(l, 0, new char[l], res);
        }
        return res;
    }

    void backtrack(int l, int pos, char[] cur, List<Long> res) {
        if (pos >= (l + 1) / 2) {
            res.add(Long.parseLong(new String(cur)));
            return;
        }
        char s = (pos == 0) ? '1' : '0';
        char e = '9';
        for (char c = s; c <= e; c++) {
            cur[pos] = c;
            cur[l - pos - 1] = c;
            backtrack(l, pos + 1, cur, res);
        }
    }

    boolean isPalindrome(String s) {
        int left = 0, right = s.length() - 1;
        while (left < right) {
            if (s.charAt(left++) != s.charAt(right--)){
                return false;
            }
        }
        return true;
    }
}