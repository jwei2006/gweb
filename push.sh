#!/bin/bash
git add api/*
git add config/*
git add service/*
git add sys/*
git add util/*

git commit -m "up"

git push

rm -rf /d/code/go/src/github.com/jwei2006/gweb

go get github.com/jwei2006/gweb