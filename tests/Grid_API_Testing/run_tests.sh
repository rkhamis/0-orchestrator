#!/bin/bash
if [ "$TRAVIS_EVENT_TYPE" == "cron" ]
then
   cd tests/Grid_API_Testing/
   nosetests-3.4 -v -s api_testing/testcases/basic_tests/test01_nodeid_apis.py:TestNodeidAPI.test001_list_nodes --tc-file=api_testing/config.ini
else
   echo "Not a cron job" 
fi

