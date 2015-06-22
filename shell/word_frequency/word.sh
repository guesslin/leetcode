# Read from the file words.txt and output the word frequency list to stdout.
sed -r 's/[[:space:]]+/\n/g' words.txt | grep -v '^$' | sort | uniq -c|sort -nr | awk '{print $2" "$1}'
