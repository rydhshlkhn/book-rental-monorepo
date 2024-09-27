// Code generated by candi v1.17.15.

package configs

import (
	"context"

	"monorepo/services/payment-service/api"
	"monorepo/sdk"
	"monorepo/services/payment-service/pkg/shared"
	"monorepo/services/payment-service/pkg/shared/repository"
	"monorepo/services/payment-service/pkg/shared/usecase"

	"github.com/golangid/candi/broker"
	"github.com/golangid/candi/candihelper"
	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/candiutils"
	"github.com/golangid/candi/codebase/factory/dependency"
	"github.com/golangid/candi/codebase/interfaces"
	"github.com/golangid/candi/config"
	"github.com/golangid/candi/config/database"
	
	"github.com/golangid/candi/logger"
	"github.com/golangid/candi/middleware"
	"github.com/golangid/candi/tracer"
	"github.com/golangid/candi/validator"
)

// LoadServiceConfigs load selected dependency configuration in this service
func LoadServiceConfigs(baseCfg *config.Config) (deps dependency.Dependency) {
	var sharedEnv shared.Environment
	candihelper.MustParseEnv(&sharedEnv)
	shared.SetEnv(sharedEnv)

	logger.InitZap()
	// logger.SetMaskLog(logger.NewMasker()) // add this for mask sensitive information

	baseCfg.LoadFunc(func(ctx context.Context) []interfaces.Closer {
		jaeger := tracer.InitJaeger(baseCfg.ServiceName)
		// redisDeps := database.InitRedis()
		sqlDeps := database.InitSQLDatabase()
		// mongoDeps := database.InitMongoDB(ctx)

		sdk.SetGlobalSDK(
			// init service client sdk
		)

		locker := &candiutils.NoopLocker{}

		brokerDeps := broker.InitBrokers(
			broker.NewKafkaBroker(),
			// broker.NewRabbitMQBroker(),
			// broker.NewRedisBroker(redisDeps.WritePool()),
		)

		validatorDeps := validator.NewValidator()
		validatorDeps.JSONSchema.SchemaStorage = validator.NewFileSystemStorage(api.JSONSchema, "jsonschema")

		// inject all service dependencies
		// See all option in dependency package
		deps = dependency.InitDependency(
			dependency.SetValidator(validatorDeps),
			dependency.SetBrokers(brokerDeps.GetBrokers()),
			dependency.SetLocker(locker),
			// dependency.SetRedisPool(redisDeps),
			dependency.SetSQLDatabase(sqlDeps),
			// dependency.SetMongoDatabase(mongoDeps),
			// ... add more dependencies
		)
		return []interfaces.Closer{ // throw back to base config for close connection when application shutdown
			jaeger,
			brokerDeps,
			locker,
			// redisDeps,
			sqlDeps,
			// mongoDeps,
		}
	})

	repository.SetSharedRepository(deps)
	usecase.SetSharedUsecase(deps)

	deps.SetMiddleware(middleware.NewMiddlewareWithOption(
		middleware.SetTokenValidator(&shared.DefaultMiddleware{}),
		middleware.SetACLPermissionChecker(&shared.DefaultMiddleware{}),
		middleware.SetUserIDExtractor(func(tokenClaim *candishared.TokenClaim) (userID string) {
			return tokenClaim.Subject
		}),
		// middleware.SetCache(deps.GetRedisPool().Cache(), middleware.DefaultCacheAge),
	))

	return deps
}
