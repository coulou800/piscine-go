curl -s https://learn.zone01dakar.sn/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{login:{_eq:\"bcoulibal\"}}){id}}"}' | cut -d ":" -f4 | sed 's/}//g' | sed 's/]//g'
