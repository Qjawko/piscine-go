curl "https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json" | jq -c '.[] | select( .id == 70) | .name'
