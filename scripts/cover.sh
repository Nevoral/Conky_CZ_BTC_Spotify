#!/bin/bash

id_current=$(cat ~/.config/conky/Graffias/current/current.txt)
id_new=`~/.config/conky/Graffias/scripts/id.sh`
cover=
imgurl=

if [ "$id_new" != "$id_current" ]; then

	cover=`ls ./covers | grep $id_new`

	if [ "$cover" == "" ]; then

	    imgurl=`~/.config/conky/Graffias/scripts/imgurl.sh $id_new`
	    wget -q -O ~/.config/conky/Graffias/covers/$id_new.jpg $imgurl &> /dev/null
	    touch ~/.config/conky/Graffias/covers/$id_new.jpg
		# clean up covers folder, keeping only the latest X amount, in below example it is 10
	    rm -f `ls -t ~/.config/conky/Graffias/covers/* | awk 'NR>100'`
	    rm -f wget-log #wget-logs are accumulated otherwise
	    cover=`ls ~/.config/conky/Graffias/covers | grep $id_new`
	fi

	if [ "$cover" != "" ]; then
		cp ~/.config/conky/Graffias/covers/$cover ~/.config/conky/Graffias/current/current.jpg
	else
		cp ~/.config/conky/Graffias/img/empty.jpg ~/.config/conky/Graffias/current/current.jpg
	fi

	echo $id_new > ~/.config/conky/Graffias/current/current.txt
fi
