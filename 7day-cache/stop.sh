#!/bin/bash
pid=$(lsof -ti :8001)
kill -9 $pid

pid=$(lsof -ti :8002)
kill -9 $pid

pid=$(lsof -ti :8003)
kill -9 $pid

rm ./server