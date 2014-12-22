#!/bin/bash

go tool yacc go.y
patch y.go y.go.patch
