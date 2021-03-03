.PHONY: release

release:
#	git tag ${tag}
#	git push origin ${tag}
	goreleaser release  --snapshot --skip-publish --rm-dist
