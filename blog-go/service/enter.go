package service

type ServiceGroup struct {
	EsService
	BaseService
	JwtService
	UserService
	ImageService
	ArticleService
	CommentService
	AdvertisementService
	FriendLinkService
	FeedbackService
	WebsiteService
	ConfigService
}

var ServiceGroupApp = &ServiceGroup{}
