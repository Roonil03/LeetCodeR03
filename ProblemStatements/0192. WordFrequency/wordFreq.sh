# Read from the file words.txt and output the word frequency list to stdout.
tr -s ' ' '\n' < words.txt | grep -v '^$' | sort | uniq -c | sort -k1,1nr -k2,2 | awk '{print $2, $1}'