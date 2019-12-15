---
title: Introduction to Shell for Data Science
tags: shell,data-science
url: https://www.datacamp.com/courses/introduction-to-shell-for-data-science
---

# 1. Manipulating files and directories

# 2. Manipulating data
## How can I view a file's contents?
```sh
cat course.txt
```

## How can I view a file's contents piece by piece?
```sh
less seasonal/spring.csv seasonal/summer.csv
```

## How can I look at the start of a file?
```sh
head seasonal/summer.csv
```

## How can I type less?
```sh
head seasonal/autum.csv
head seasonal/spring.csv
```

## How can I control what commands do?
```sh
head -5 seasonal/winter.csv
```

## How can I list everything below a directory?
```sh
ls -RF
```

## How can I get help for a command
```sh
man tail
tail -n +7 seasonal/spring.csv
```

## How can I select columns from a file?
```sh
cut -d , -f 1 seasonal/spring.csv
```

## How can I repeat commands?
```sh
head summer.csv
cd seasonal
!head
history
!head
```

## How can I select lines containing specific values?
```sh
grep molar seasonal/autum.csv
grep -n -v molar seasonal/spring.csv
grep -c incisor seasonal/autum.csv seasonal/winter.csv
```

## How can I store a command's output in a file?
```sh
tail -5 seasonal/winter.csv > last.csv
```

# 3. Combining tools
## How can I use a command's output as an input?
```sh
tail -2 seasonal/winter.csv > bottom.csv
head -1 bottom.csv
```

## What's a better way to combine commands?
```sh
cut -d, -f2 seasonal/summer.csv | grep -v Tooth
```

## How can I combine many commands?
```sh
cut -d, -f2 seasonal/summer.csv | grep -v Tooth | head -n 1
```

## How can I count the records in a file?
```sh
grep 2017-07 seasonal/spring.csv | wc -l
```

## How can I specify many files at once?
```sh
head -n 3 seasonal/s*
```

## How can I sort lines of text?
```sh
cut -d , -f 2 seasonal/winter.csv | grep -v Tooth | sort -r
```

## How can I remove duplicate lines?
```sh
cut -d , -f 2 seasonal/winter.csv | grep -v Tooth | sort | uniq -c
```

## How can I save the output of a pipe?
```sh
> result.txt head -n 3 seasonal/winter.csv
```

## Wrapping up
```sh
wc -l seasonal/*
wc -l seasonal/* | grep -v total
wc -l seasonal/* | grep -v total | sort -n | head -n 1
```

# 4. Batch processing
## How does the shell store information?
```sh
echo $HISTFILESIZE
```

## How can I print a variable's value
```sh
echo $OSTYPE
```

## How else does the shell store information? 
```sh
testing=seasonal/winter.csv
head -n 1 $testing
```

## How can I repeat a command many times?
```sh
for f in docx odt pdf; do echo $f; done
```

## How can I repeat a command once for each file?
```sh
for f in people/*; do echo $f; done
```

## How can I run many commands in a single loop?
```sh
file=seasonal/*.csv
for f in $file; do grep 2017-07 $f | tail -1; done
```

# 5. Creating new tools
## How can I record what I just did?
```sh
cp seasonal/s*.csv .
grep -h -v Tooth s*.csv > temp.csv
history | tail -n 3 > steps.txt
```


