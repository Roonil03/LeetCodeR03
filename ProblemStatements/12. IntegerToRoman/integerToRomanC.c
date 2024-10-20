#include <stdio.h>
#include <string.h>

void intToRoman(int num, char* result) {
    const int values[] = {1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1};
    const char* symbols[] = {"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"};
    
    result[0] = '\0'; // Initialize result as an empty string

    for (int i = 0; i < 13; i++) {
        while (num >= values[i]) {
            strcat(result, symbols[i]);
            num -= values[i];
        }
    }
}

int main() {
    int number = 1994; // Example input
    char roman[20]; // Buffer for the result
    intToRoman(number, roman);
    printf("Roman numeral: %s\n", roman);
    return 0;
}
