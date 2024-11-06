class Solution:
    def divide(self, dividend: int, divisor: int) -> int:
        if dividend == -2147483648 and divisor == -1:
            return 2147483647
        neg = (dividend < 0) != (divisor < 0)
        dividend, divisor = abs(dividend), abs(divisor)
        q = 0
        while dividend >= divisor:
            temp, mul = divisor, 1
            while dividend >= (temp << 1):
                temp <<= 1
                mul <<= 1
            dividend -= temp
            q += mul
        return -q if neg else q