#!/bin/bash

shortenertest -test.v -test.run=^TestIteration1$ -binary-path=cmd/shortener/shortener
shortenertest -test.v -test.run=^TestIteration2$ -binary-path=cmd/shortener/shortener