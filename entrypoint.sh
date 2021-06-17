#!/bin/sh -l

export region=$1
export environment=$2
export client_id=$3
export client_secret=$4
export service_id=$5
export private_key=$6

/iam-service-login
