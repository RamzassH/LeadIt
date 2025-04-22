package interceptors

import (
	"context"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type CtxKey string

const (
	CtxUserID CtxKey = "userID"
	CtxEmail  CtxKey = "email"
)

func JwtUnaryServerInterceptor(secret string) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, fmt.Errorf("missing metadata")
		}

		tokens := md.Get("authorization")
		if len(tokens) == 0 {
			return nil, fmt.Errorf("authorization token is not provided")
		}

		tokenStr := tokens[0]

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			return nil, fmt.Errorf("invalid token: %v", err)
		}

		claims := token.Claims.(jwt.MapClaims)

		userID, ok := claims["uid"].(float64)
		if !ok {
			return nil, fmt.Errorf("invalid uid in token claims")
		}
		email, _ := claims["email"].(string)

		newCtx := context.WithValue(ctx, CtxUserID, int64(userID))
		newCtx = context.WithValue(newCtx, CtxEmail, email)

		return handler(newCtx, req)
	}
}
