package onlinestore

//go:generate swagger generate server --exclude-main -A online-store-server -t gen -f ./api/swagger.yml --principal models.Principal
//go:generate swagger -q generate client -A online-store-service -f api/swagger.yml -c pkg/client -m gen/models --principal models.Principal
