# postcar

for i in `seq 1 9` 
do
     curl -X POST -d "{\"name\": \"$(($RANDOM % 10000))\", \"home\": \"edgeP\", \"violation\": {\"speedViolation\": \"$(($RANDOM % 10))\", \"ignoreSignal\": \"$(($RANDOM % 10))\"}}" http://localhost:3000/v1/service/car
done