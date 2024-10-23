#include <stdio.h>
#include <string.h>
#include <stdlib.h>

#define TABLE_SIZE 256

typedef struct HashMap {
    int table[TABLE_SIZE];
} HashMap;

void initHashMap(HashMap *map) {
    // Set default values to 0
    for (int i = 0; i < TABLE_SIZE; i++) {
        map->table[i] = 0;
    }
    map->table['I'] = 1;
    map->table['V'] = 5;
    map->table['X'] = 10;
    map->table['L'] = 50;
    map->table['C'] = 100;
    map->table['D'] = 500;
    map->table['M'] = 1000;
}

int romanToInt(char *s) {
    HashMap map;
    initHashMap(&map);

    int sum = 0;
    int len = strlen(s);

    for (int i = 0; i < len; i++) {
        if (i + 1 < len && map.table[s[i]] < map.table[s[i + 1]]) {
            sum -= map.table[s[i]];
        } else {
            sum += map.table[s[i]];
        }
    }

    return sum;
}
/*
int main() {
    printf("%d\n", romanToInt("MCMXCIV")); // Output: 1994
    return 0;
}
*/