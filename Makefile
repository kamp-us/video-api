gen:
	# Auto-generate code
	protoc --twirp_out=. --go_out=. rpc/video-api/service.proto