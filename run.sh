#!/bin/bash

docker run -p 8080:8080 --rm -v `pwd`/data:/app/data ramonberrutti/desafio-tegra-backend