import { lazy } from "react";

import { createBrowserRouter } from "react-router-dom";

//
// Web Frontend
//
const WebLayout = lazy(() => import("@/components/layout/WebLayout"));
const HomePage = lazy(() => import("@/features/web/home/HomePage"));
const SearchPage = lazy(() => import("@/features/web/search/SearchPage"));
const FriendLinkPage = lazy(
  () => import("@/features/web/friendLink/FriendLinkPage")
);
const AboutPage = lazy(() => import("@/features/web/about/AboutPage"));
const ArticlePage = lazy(() => import("@/features/web/article/ArticlePage"));

//
// Auth
//
const LoginPage = lazy(() => import("@/features/auth/login/LoginPage"));

//
// Dashboard
//
const DashboardLayout = lazy(
  () => import("@/features/dashboard/layout/DashboardLayout")
);
const DashboardHomePage = lazy(
  () => import("@/features/dashboard/home/DashboardHomePage")
);

// User center
const UserInfoPage = lazy(
  () => import("@/features/dashboard/userCenter/UserInfoPage")
);
const UserStarPage = lazy(
  () => import("@/features/dashboard/userCenter/UserStarPage")
);
const UserCommentPage = lazy(
  () => import("@/features/dashboard/userCenter/UserCommentPage")
);
const UserFeedbackPage = lazy(
  () => import("@/features/dashboard/userCenter/UserFeedbackPage")
);

// Users
const UserListPage = lazy(
  () => import("@/features/dashboard/users/UserListPage")
);

// Articles
const ArticlePublishPage = lazy(
  () => import("@/features/dashboard/articles/ArticlePublishPage")
);
const CommentListPage = lazy(
  () => import("@/features/dashboard/articles/CommentListPage")
);
const ArticleListPage = lazy(
  () => import("@/features/dashboard/articles/ArticleListPage")
);

// Images
const ImageListPage = lazy(
  () => import("@/features/dashboard/images/ImageListPage")
);

// System
const FeedbackListPage = lazy(
  () => import("@/features/dashboard/system/FeedbackListPage")
);
const AdvertisementListPage = lazy(
  () => import("@/features/dashboard/system/AdvertisementListPage")
);
const FriendLinkListPage = lazy(
  () => import("@/features/dashboard/system/FriendLinkListPage")
);
const LoginLogsPage = lazy(
  () => import("@/features/dashboard/system/LoginLogsPage")
);

// App Config section
const AppConfigLayout = lazy(
  () => import("@/features/dashboard/system/appConfig/AppConfigLayout")
);
const SiteConfigPage = lazy(
  () => import("@/features/dashboard/system/appConfig/SiteConfigPage")
);
const SystemConfigPage = lazy(
  () => import("@/features/dashboard/system/appConfig/SystemConfigPage")
);
const EmailConfigPage = lazy(
  () => import("@/features/dashboard/system/appConfig/EmailConfigPage")
);
const GithubConfigPage = lazy(
  () => import("@/features/dashboard/system/appConfig/GithubConfigPage")
);
const JwtConfigPage = lazy(
  () => import("@/features/dashboard/system/appConfig/JwtConfigPage")
);

//
// Errors
//
const NotFoundPage = lazy(() => import("@/features/error/NotFoundPage"));

export const routes = [
  {
    path: "/",
    element: <WebLayout />,
    children: [
      { index: true, element: <HomePage /> },
      { path: "article/:id", element: <ArticlePage /> },
      { path: "search", element: <SearchPage /> },
      {
        path: "friend-link",
        element: <FriendLinkPage />,
      },
      { path: "about", element: <AboutPage /> },
    ],
  },

  { path: "/login", element: <LoginPage /> },

  {
    path: "/dashboard",
    element: <DashboardLayout />,
    meta: { requiresAuth: true },
    children: [
      { index: true, element: <DashboardHomePage />, meta: { title: "主页" } },

      {
        path: "user-center",
        children: [
          { path: "user-info", element: <UserInfoPage /> },
          { path: "user-star", element: <UserStarPage /> },
          { path: "user-comment", element: <UserCommentPage /> },
          { path: "user-feedback", element: <UserFeedbackPage /> },
        ],
      },

      {
        path: "users",
        meta: { requiresAdmin: true },
        children: [{ path: "user-list", element: <UserListPage /> }],
      },

      {
        path: "articles",
        meta: { requiresAdmin: true },
        children: [
          { path: "article-publish", element: <ArticlePublishPage /> },
          { path: "comment-list", element: <CommentListPage /> },
          { path: "article-list", element: <ArticleListPage /> },
        ],
      },

      {
        path: "images",
        meta: { requiresAdmin: true },
        children: [{ path: "image-list", element: <ImageListPage /> }],
      },

      {
        path: "system",
        meta: { requiresAdmin: true },
        children: [
          { path: "feedback-list", element: <FeedbackListPage /> },
          { path: "advertisement-list", element: <AdvertisementListPage /> },
          { path: "friend-link-list", element: <FriendLinkListPage /> },
          { path: "login-logs", element: <LoginLogsPage /> },

          {
            path: "app-config",
            element: <AppConfigLayout />,
            children: [
              { path: "site-config", element: <SiteConfigPage /> },
              { path: "system-config", element: <SystemConfigPage /> },
              { path: "email-config", element: <EmailConfigPage /> },
              { path: "qq-config", element: <GithubConfigPage /> },
              { path: "jwt-config", element: <JwtConfigPage /> },
            ],
          },
        ],
      },
    ],
  },

  { path: "*", element: <NotFoundPage /> },
];

export const router = createBrowserRouter(routes);
