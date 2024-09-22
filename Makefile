# protoc
protoc :
	protoc --go_out=. --go-grpc_out=. calculator.proto
# --go_out= 메시지 구조체, 메서드
#   -> calculator.pb.go 생성
# --go-grpc_out= gRPC 서버스 정의, 클라이언트 및 서버 인터페이스 위치
#   -> calculator_grpc.pb.go 생성