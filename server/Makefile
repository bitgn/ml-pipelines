.PHONY: pb


pb:
	protoc -I ../proto events.proto --go_out=events
	protoc -I ../proto db.proto --go_out=db
	protoc -I ../proto api.proto --go_out=plugins=grpc:api
	sed -i '' 's_events "."_events "mlp/catalog/events"_g' db/db.pb.go

	sed -i '' 's_events "."_events "mlp/catalog/events"_g' api/api.pb.go
