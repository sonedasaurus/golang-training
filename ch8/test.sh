#!/bin/sh

go build -o clock
cd clockwall 
go build
cd ..

TZ=US/Eastern ./clock -port 8010 &
TZ=Asia/Tokyo ./clock -port 8020 &
TZ=Europe/London ./clock -port 8030 &

cd clockwall
NewYork=localhost:8010 Tokyo=localhost:8020 London=localhost:8030 ./clockwall
