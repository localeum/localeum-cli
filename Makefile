.PHONY: build, release

build:
	goreleaser release --snapshot --skip-publish --rm-dist

release:
	git tag ${tag}
	git push origin ${tag}
	goreleaser release --rm-dist
