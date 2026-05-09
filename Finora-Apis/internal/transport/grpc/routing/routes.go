package routing

// Service names registered on the gRPC server.
const (
	FinoraServiceName = "finora.v1.FinoraService"
)

// Fully-qualified unary method names used by interceptors and policies.
const (
	MethodHealth = "/finora.v1.FinoraService/Health"
	MethodLogin  = "/finora.v1.FinoraService/Login"

	MethodGRPCHealthCheck = "/grpc.health.v1.Health/Check"
	MethodGRPCHealthWatch = "/grpc.health.v1.Health/Watch"

	MethodReflectionV1Alpha = "/grpc.reflection.v1alpha.ServerReflection/ServerReflectionInfo"
	MethodReflectionV1      = "/grpc.reflection.v1.ServerReflection/ServerReflectionInfo"
)

var publicMethods = map[string]struct{}{
	MethodGRPCHealthCheck: {},
	MethodGRPCHealthWatch: {},
	MethodReflectionV1:    {},
	MethodReflectionV1Alpha: {},
	MethodHealth:          {},
	MethodLogin:           {},
}

// IsPublicMethod reports whether a method should bypass JWT auth.
func IsPublicMethod(fullMethod string) bool {
	_, ok := publicMethods[fullMethod]
	return ok
}
