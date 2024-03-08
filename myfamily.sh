curl -s https://learn.zone01oujda.ma/assets/superhero/all.json | jq --arg hero_id "$HERO_ID" '.[] | select(.id == ($hero_id | tonumber)) | .connections.relatives' | tr -d "\""
