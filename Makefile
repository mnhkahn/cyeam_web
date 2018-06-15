branch := $(shell git branch | grep \* | cut -d ' ' -f2)

.PHNOY: push
push:
	git commit -am "SYNC BIN $DATE"
	git push heroku ${branch}

DATE             := $(shell date --rfc-3339=seconds)
