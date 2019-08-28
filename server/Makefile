.PHONY: pb, build, clean


pb:
	protoc -I ../proto events.proto --go_out=events
	sed -i '' 's_vo "."_vo "mlp/catalog/vo"_g' events/events.pb.go

	protoc -I ../proto db.proto --go_out=db
	protoc -I ../proto vo.proto --go_out=vo

	protoc -I ../proto mlp_api.proto --go_out=plugins=grpc:mlp_api
	sed -i '' 's_vo "."_vo "mlp/catalog/vo"_g' mlp_api/mlp_api.pb.go


	protoc -I ../proto test_api.proto --go_out=plugins=grpc:test_api

	sed -i '' 's_vo "."_vo "mlp/catalog/vo"_g' db/db.pb.go
	sed -i '' 's_events "."_events "mlp/catalog/events"_g' test_api/test_api.pb.go

clean:
	rm -rf dist || True
build:
	go build -o dist/mlp-server
