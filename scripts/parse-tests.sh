#! /bin/sh

read -r summary

echo '# Coverage Summary'
echo ''
echo '| package | coverage |'
echo '| --- | --- |'
echo "$summary" | while IFS= read -r line
do
  pkg=$(echo "$line" | cut -f2)
  coverage=$(echo "$line" | cut -f4 | cut -d' ' -f 2)
  echo "| $pkg | $coverage |"
done
