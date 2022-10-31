# !/bin/bash

base=1665768990
cat run_log | grep mined | sed 's/[^0-9]/ /g' | awk -F ' ' '{print $2 " " $1}' > data.txt
