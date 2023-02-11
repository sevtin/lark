#!/usr/bin/env bash

source ./../cfg/services.cfg

for ((i=1; i<=2; i++))
do
  gp=$[7300+($i * 10)]
  wp=$[gp+1]
  mp=$[gp+2]
  sid=$[$i+1]
  echo "run msg_gateway" ${gp} ${wp} ${mp} ${sid}
  nohup ./../bin/lark_msg_gateway -gp=${gp} -wp=${wp} -mp=${mp} -sid=${sid} > /var/log/lark/lark_msg_gateway_${sid}.log 2>&1 &
done

for ((i=1; i<2; i++))
do
  gp=$[7400+($i * 10)]
  mp=$[gp+1]
  sid=$[$i+1]
  echo "run dist" ${gp} ${mp} ${sid}
  nohup ./../bin/lark_dist -gp=${gp} -mp=${mp} -sid=${sid} > /var/log/lark/lark_dist_${sid}.log 2>&1 &
done

length=${#service_names}
for (( i=0; i <= $length; i++ )); do
  service=${service_names[$i]}
  if [ -z ${service} ];then
      continue
  fi
  echo "run "${service}
  nohup ./../bin/lark_${service} > /var/log/lark/${service}.log 2>&1 &
  sleep 1
done

echo "run success..."

sleep 10
# fixme prevents the suzaku service exit after execution in the docker container
tail -f /dev/null
# ping 8.8.8.8
# ./../bin/lark_msg_gateway cfg ./configs/msg_gateway.yaml gp 33000 wp 33001 -sid 2 /var/log/lark/lark_msg_gateway_2.log 2>&1