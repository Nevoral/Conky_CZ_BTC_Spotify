#!/bin/bash

id_current=$(cat ./current/current.txt)
id_new=`./scripts/id.sh`
cover=
imgurl=

if [ "$id_new" != "$id_current" ]; then

	cover=`ls ./covers | grep $id_new`

	if [ "$cover" == "" ]; then

	    imgurl=`./scripts/imgurl.sh $id_new`
	    wget -q -O ./covers/$id_new.jpg $imgurl &> /dev/null
	    touch ./covers/$id_new.jpg
		# clean up covers folder, keeping only the latest X amount, in below example it is 10
	    rm -f `ls -t ./covers/* | awk 'NR>100'`
	    rm -f wget-log #wget-logs are accumulated otherwise
	    cover=`ls ./covers | grep $id_new`
	fi

	if [ "$cover" != "" ]; then
		cp ./covers/$cover ./current/current.jpg
	else
		cp ./img/empty.jpg ./current/current.jpg
	fi

	echo $id_new > ./current/current.txt
fi
