compile-protos:
	$(foreach file, $(wildcard protos/*.proto), \
		dir=`echo $(basename $(file))`; \
		filename=`echo $$dir | sed 's:.*/::'`; \
		mkdir -p $$dir; \
		protoc --go_out=. --go_opt=paths=source_relative \
			--go-grpc_out=. --go-grpc_opt=paths=source_relative $(file); \
		mv protos/$${filename}.pb.go protos/$$filename/$${filename}.pb.go; \
		mv protos/$${filename}_grpc.pb.go protos/$$filename/$${filename}_grpc.pb.go;)
