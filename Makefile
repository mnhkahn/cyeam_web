branch := $(shell git branch | grep \* | cut -d ' ' -f2)

.PHNOY: run
run:
	go build && PORT=8090 ./cyeam

.PHNOY: push
push:
	git commit -am "SYNC BIN $DATE"
	git push origin ${branch}

DATE             := $(shell date --rfc-3339=seconds)

