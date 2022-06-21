#!/bin/bash

killall conky
sleep 2s
		
conky -c ./Graffias.conf &> /dev/null &
