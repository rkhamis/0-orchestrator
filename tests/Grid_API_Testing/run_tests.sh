#!/bin/bash

point=$1
if [ "$TRAVIS_EVENT_TYPE" == "cron" ] || [ "$TRAVIS_EVENT_TYPE" == "api" ]
 then
   if [ "$point" == "before" ]
    then
      pip3 install -r tests/Grid_API_Testing/requirements.txt
      pip3 install git+https://github.com/gigforks/packet-python.git
      bash tests/Grid_API_Testing/install_env.py master $ZT_NET_ID $ZT_TOKEN
      cd tests/Grid_API_Testing; python3 orch_packet_machines.py create $PACKET_TOKEN $ZT_NET_ID
      export PYTHONPATH='./'
   elif [ "$point" == "run" ]
    then
      echo "Running tests .."
      cd tests/Grid_API_Testing/
   nosetests-3.4 -v -s api_testing/testcases/basic_tests/test01_nodeid_apis.py:TestNodeidAPI.test001_list_nodes --tc-file=api_testing/config.ini
   elif [ "$point" == "after" ]
    then
      cd tests/Grid_API_Testing/
      python3 orch_packet_machines.py delete $PACKET_TOKEN
   fi
 else
   echo "Not a cron job" 
fi

