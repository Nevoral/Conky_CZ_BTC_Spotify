#!/bin/bash

killall conky
sleep 2s
		
conky -c ~/.config/conky/Graffias/Graffias.conf &> /dev/null &
