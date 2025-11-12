#!/bin/bash
TIMESTAMP=$(date +"%Y%m%d-%H%M")
mongodump --uri="mongodb://admin:admin123@localhost:27017/biblioteca" --out=./backup-$TIMESTAMP
