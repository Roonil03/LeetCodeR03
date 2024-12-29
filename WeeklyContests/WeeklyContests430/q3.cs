/* Could not solve it entirely during the time of the contest,
 * it has been solved post the time limit of the contest.
 */

 public class Solution
{
    private int[,] cnt = new int[1001, 1001];

    private int Gcd(int a, int b)
    {
        return b == 0 ? a : Gcd(b, a % b);
    }

    public long NumberOfSubsequences(int[] n)
    {
        long res = 0;
        Array.Clear(cnt, 0, cnt.Length);
        for (int q = n.Length - 5; q > 1; --q)
        {
            for (int s = q + 4, r = q + 2; s < n.Length; ++s)
            {
                int f = Gcd(n[r], n[s]);
                ++cnt[n[r] / f, n[s] / f];
            }
            for (int p = 0; p + 1 < q; ++p)
            {
                int f = Gcd(n[p], n[q]);
                res += cnt[n[q] / f, n[p] / f];
            }
        }
        return res;
    }
}

// The provided code utilizes several important Data Structures and Algorithms (DSA) concepts, including arrays, 
// greatest common divisor (GCD) computation, and nested loops for iteration. The code defines a class Solution with 
// a method NumberOfSubsequences that counts specific subsequences in an array of integers. It employs a 2D array cnt 
// to keep track of counts of pairs of numbers normalized by their GCD. The Gcd method is a recursive function to 
// calculate the greatest common divisor of two numbers using the Euclidean algorithm. The NumberOfSubsequences method 
// iterates through the array in a nested manner, updating the count array and using it to calculate the result based 
// on the previously stored counts. The main logic involves counting pairs of values divided by their GCD and using 
// this information to determine the number of specific subsequences that meet the criteria defined in the problem.