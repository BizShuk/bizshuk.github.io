#!/bin/bash

README="README.md"
> $README
echo "# Guideline:" >> $README
guidefiles=`ls | grep -P "^[^20\d\d\.\d{1,2}\.\d{1,2}].*\.md$"`

for file in $guidefiles
do
    echo "- [${file:0:-3}](${file})" >> $README
done
echo "- [國家介紹](Region)" >> $README


echo "" >> $README
echo "# Blog:" >> $README
blogfiles=`ls | grep -P "^20\d\d\.\d{1,2}\.\d{1,2}.*\.md$"`

for file in $blogfiles
do
    echo "- [${file:0:-3}](${file})" >> $README
done
