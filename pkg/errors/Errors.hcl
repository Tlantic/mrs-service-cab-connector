settings{
    Name = "apierr"
    Package = "apierr"
    Service = "connector-api"
	Keys = {
    }
}

message "error" "ErrStoreNotFound"{
    StatusHTTP = 400
    Message = "Store not found"
    Kind = 1
}

message "error" "ErrInternal"{
    StatusHTTP = 500
    Message = "Internal server error"
    Kind = 2
}

message "error" "ErrValidation"{
    StatusHTTP = 400
    Message = "Validation error"
    Kind = 3
}

message "error" "ErrGrpcConnection"{
    StatusHTTP = 500
    Message = "Service connector comunication error"
    Kind = 4
}

message "error" "ErrNotImplemented"{
    StatusHTTP = 501
    Message = "Not Implemented"
    Kind = 5
}

message "error" "ErrSkuNotFound"{
    StatusHTTP = 400
    Message = "Sku not found"
    Kind = 6
}