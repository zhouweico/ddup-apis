package router

import (
	"ddup-apis/docs"
	"ddup-apis/internal/config"
	"ddup-apis/internal/db"
	"ddup-apis/internal/handler"
	"ddup-apis/internal/middleware"
	"ddup-apis/internal/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 设置 Swagger 信息
	cfg := config.GetConfig()
	docs.SwaggerInfo.Host = cfg.Swagger.Host
	docs.SwaggerInfo.Schemes = cfg.Swagger.Schemes

	// 添加全局中间件
	r.Use(middleware.Logger())
	r.Use(middleware.Cors())
	r.Use(middleware.ErrorHandler())
	r.Use(gin.Recovery())

	// 初始化 services
	userService := service.NewUserService(db.DB)
	socialService := service.NewSocialService(db.DB)
	organizationService := service.NewOrganizationService(db.DB)

	// 初始化 handlers
	userHandler := handler.NewUserHandler(userService)
	healthHandler := handler.NewHealthHandler()
	socialHandler := handler.NewSocialHandler(socialService)
	organizationHandler := handler.NewOrganizationHandler(organizationService)

	// 健康检查路由（放在 API v1 路由组之外）
	r.GET("/health", healthHandler.Check)

	// API v1 路由组
	v1 := r.Group("/api/v1")
	{
		// 认证相关路由
		auth := v1.Group("/auth")
		{
			auth.POST("/register", userHandler.Register)
			auth.POST("/login", userHandler.Login)
			auth.POST("/logout", userHandler.Logout)
		}

		// 用户相关路由
		users := v1.Group("/users")
		users.Use(middleware.JWTAuth(userService))
		{
			// 用户个人操作
			users.GET("", userHandler.GetUser)                 // 获取个人信息
			users.PUT("", userHandler.UpdateUser)              // 更新个人信息
			users.DELETE("", userHandler.DeleteUser)           // 注销账号
			users.PUT("/password", userHandler.ChangePassword) // 修改密码

			// 用户社交媒体
			socials := users.Group("/socials")
			{
				socials.POST("", socialHandler.CreateSocial)
				socials.GET("", socialHandler.GetUserSocial)
				socials.PUT("/:id", socialHandler.UpdateSocial)
				socials.DELETE("/:id", socialHandler.DeleteSocial)
			}
		}

		// 组织相关路由
		orgs := v1.Group("/organizations")
		orgs.Use(middleware.JWTAuth(userService))
		{
			// 组织基本操作
			orgs.POST("", organizationHandler.CreateOrganization)
			orgs.GET("", organizationHandler.GetUserOrganization)
			orgs.PUT("/:id", organizationHandler.UpdateOrganization)
			orgs.DELETE("/:id", organizationHandler.DeleteOrganization)

			// 组织成员管理
			orgs.POST("/:id/join", organizationHandler.JoinOrganization)
			orgs.GET("/:id/members", organizationHandler.GetOrganizationMembers)

			// 组织管理员操作
			members := orgs.Group("/:id/members")
			{
				members.POST("", organizationHandler.AddOrganizationMember)
				members.PUT("/:userid", organizationHandler.UpdateOrganizationMember)
				members.DELETE("/:userid", organizationHandler.RemoveOrganizationMember)
				members.PUT("/:userid/role", organizationHandler.SetOrganizationMemberRole)
			}
		}
	}

	// Swagger API 文档路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
