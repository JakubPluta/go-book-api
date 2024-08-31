package health

import "net/http"

// Read godoc
//
//  @summary        Read health
//  @description    Read health
//  @tags           health
//  @success        200
//  @router         /../health [get]

func Read(w http.ResponseWriter, r *http.Request) {}
