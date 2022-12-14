// Code generated by erygo
package apierr

import erygo "github.com/andrepinto/erygo"

// ErrStoreNotFound error
func ErrStoreNotFound(params ...func(*erygo.Err)) *erygo.Err {
	err := &erygo.Err{
		Details: []string{},
		Info: erygo.Info{
			Kind:    1,
			Service: "connector-api",
		},
		Message:    "Store not found",
		StatusHTTP: 400,
	}
	for _, param := range params {
		param(err)
	}
	for i, detail := range err.Details {
		det := renderTemplate(detail)
		err.Details[i] = det
	}
	return err
}

// ErrInternal error
func ErrInternal(params ...func(*erygo.Err)) *erygo.Err {
	err := &erygo.Err{
		Details: []string{},
		Info: erygo.Info{
			Kind:    2,
			Service: "connector-api",
		},
		Message:    "Internal server error",
		StatusHTTP: 500,
	}
	for _, param := range params {
		param(err)
	}
	for i, detail := range err.Details {
		det := renderTemplate(detail)
		err.Details[i] = det
	}
	return err
}

// ErrValidation error
func ErrValidation(params ...func(*erygo.Err)) *erygo.Err {
	err := &erygo.Err{
		Details: []string{},
		Info: erygo.Info{
			Kind:    3,
			Service: "connector-api",
		},
		Message:    "Validation error",
		StatusHTTP: 400,
	}
	for _, param := range params {
		param(err)
	}
	for i, detail := range err.Details {
		det := renderTemplate(detail)
		err.Details[i] = det
	}
	return err
}

// ErrGrpcConnection error
func ErrGrpcConnection(params ...func(*erygo.Err)) *erygo.Err {
	err := &erygo.Err{
		Details: []string{},
		Info: erygo.Info{
			Kind:    4,
			Service: "connector-api",
		},
		Message:    "Service connector comunication error",
		StatusHTTP: 500,
	}
	for _, param := range params {
		param(err)
	}
	for i, detail := range err.Details {
		det := renderTemplate(detail)
		err.Details[i] = det
	}
	return err
}

// ErrNotImplemented error
func ErrNotImplemented(params ...func(*erygo.Err)) *erygo.Err {
	err := &erygo.Err{
		Details: []string{},
		Info: erygo.Info{
			Kind:    5,
			Service: "connector-api",
		},
		Message:    "Not Implemented",
		StatusHTTP: 501,
	}
	for _, param := range params {
		param(err)
	}
	for i, detail := range err.Details {
		det := renderTemplate(detail)
		err.Details[i] = det
	}
	return err
}

// ErrSkuNotFound error
func ErrSkuNotFound(params ...func(*erygo.Err)) *erygo.Err {
	err := &erygo.Err{
		Details: []string{},
		Info: erygo.Info{
			Kind:    6,
			Service: "connector-api",
		},
		Message:    "Sku not found",
		StatusHTTP: 400,
	}
	for _, param := range params {
		param(err)
	}
	for i, detail := range err.Details {
		det := renderTemplate(detail)
		err.Details[i] = det
	}
	return err
}

// ErrUnauthorized error
func ErrUnauthorized(params ...func(*erygo.Err)) *erygo.Err {
	err := &erygo.Err{
		Details: []string{},
		Info: erygo.Info{
			Kind:    8,
			Service: "connector-api",
		},
		Message:    "Unauthorized external call",
		StatusHTTP: 401,
	}
	for _, param := range params {
		param(err)
	}
	for i, detail := range err.Details {
		det := renderTemplate(detail)
		err.Details[i] = det
	}
	return err
}
