
# Serve docs
doc:
	(sleep 1 && open http://localhost:8000/pkg/github.com/kyoto-framework/zen/v2) &
	godoc -http=:8000
