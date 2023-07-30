#!/bin/usr/env bash 

for (( i = 0; i < 1000; i++ )); 
do 
  printf "Sending request [%i]\n" $i
  printf "\n"
  curl -X POST http://localhost:8000/order \
    -H "Content-Type: application/json" \
    -d '{
      "size":"extra large", 
      "flavor":"premium chocolate", 
      "decoration":"none", 
      "package":"normal", 
      "delivery":"fast"
    }' 
  printf "\n"  
done
