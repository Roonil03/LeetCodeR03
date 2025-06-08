# Read from the file file.txt and print its transposed content to stdout.
awk '
{
    for (i = 1; i <= NF; i++)  {
        a[NR,i] = $i
    }
    max_col = NF
}
END {
    for (i = 1; i <= max_col; i++) {
        for (j = 1; j <= NR; j++) {
            printf "%s%s", a[j,i], (j==NR ? ORS : OFS)
        }
    }
}' file.txt