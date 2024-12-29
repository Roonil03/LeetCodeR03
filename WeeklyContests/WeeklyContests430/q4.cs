/* Could not solve it entirely during the time of the contest,
 * it has been solved post the time limit of the contest.
 */

 public class Solution
{
    private const int MOD = 1000000007;

    private long ModPow(long baseValue, long exp, long mod)
    {
        long result = 1;
        while (exp > 0)
        {
            if ((exp & 1) > 0)
                result = (result * baseValue) % mod;
            baseValue = (baseValue * baseValue) % mod;
            exp >>= 1;
        }
        return result;
    }

    public int CountGoodArrays(int n, int m, int k)
    {
        if (k >= n) return 0;
        long total = (m * ModPow(m - 1, n - 1 - k, MOD)) % MOD;
        for (int i = 1; i <= k; i++)
        {
            total = (total * (n - i)) % MOD;
            total = (total * ModPow(i, MOD - 2, MOD)) % MOD;
        }
        return (int)total;
    }
}

//The provided code includes several key Data Structures and Algorithms (DSA) concepts such as 
// modular arithmetic, exponentiation by squaring, and combinatorial calculations. The Solution 
// class defines a method CountGoodArrays that counts the number of valid arrays given the constraints. 
// The ModPow method is a helper function that computes the modular exponentiation using the efficient 
// exponentiation by squaring algorithm, which is crucial for handling large powers under a modulus. 
// The CountGoodArrays method calculates the total number of valid arrays using combinatorial principles 
// and modular arithmetic to ensure results fit within standard integer ranges. The use of MOD ensures 
// that computations remain within bounds to prevent overflow. It iterates through the range to compute 
// combinations and applies Fermat's Little Theorem for modular inverses, which is used in the calculation 
// of binomial coefficients. This code is designed to handle large inputs efficiently by leveraging these 
// advanced mathematical techniques.