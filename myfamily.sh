curl 'https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json' | jq -c -r ".[] | select( .id == $HERO_ID ) | .connections.relatives"

