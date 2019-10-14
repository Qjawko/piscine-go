curl 'https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json' | jq -c -j ".[] | select( .id == $HERO_ID ) | .connections.relatives" \n

