package proto

//go:generate protoc --proto_path=. --go_out=plugins=grpc:. api.proto auth_external.proto auth_types.proto catalog_external.proto price_types.proto stock_types.proto label_types.proto task_types.proto task_external.proto process_task.proto
