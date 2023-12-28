# goose mysql "root:@/learn-auth-golang?parseTime=true" up
# goose mysql "root:@/learn-auth-golang?parseTime=true" down  
if [ -f ./.env ]; then
  export $(echo $(cat ./.env | sed 's/#.*//g'| xargs) | envsubst)
fi
# printenv
echo $DRIVER $DSN

cd migration 

# sh -c "GOOSE_DRIVER=\"$DRIVER\" GOOSE_DBSTRING=\"$DSN\" goose down"
sh -c "GOOSE_DRIVER=\"$DRIVER\" GOOSE_DBSTRING=\"$DSN\" goose up"
