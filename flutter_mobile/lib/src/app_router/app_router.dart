import 'package:auto_route/auto_route.dart';
import 'package:flutter_mobile/src/app_router/app_router.gr.dart';
import 'package:flutter_mobile/src/app_router/auth_guard.dart';
import 'package:flutter_mobile/src/features/auth/service/auth_service.dart';

@AutoRouterConfig(replaceInRouteName: "Screen,Route")
class AppRouter extends RootStackRouter {
  final AuthService _authService;

  AppRouter({
    required AuthService authService,
  }) : _authService = authService;

  @override
  List<AutoRoute> get routes => [
        AutoRoute(
          path: '/app',
          initial: true,
          page: AppNavRoute.page,
          children: [
            _mapRoutes,
            _socialRoutes,
            _chatsRoutes,
            _profileRoutes,
          ],
          guards: [
            AuthGuard(authService: _authService),
          ],
        ),
        AutoRoute(
          path: "/auth",
          page: AuthRoute.page,
          guards: [
            UnauthGuard(authService: _authService),
          ],
        ),
        AutoRoute(
          path: "/login",
          page: LoginRoute.page,
          guards: [
            UnauthGuard(authService: _authService),
          ],
        ),
        AutoRoute(
          path: "/reg",
          page: RegRoute.page,
          guards: [
            UnauthGuard(authService: _authService),
          ],
        ),
      ];

  final _mapRoutes = AutoRoute(
    path: "map",
    page: MapFragmentRoute.page,
  );

  final _socialRoutes = AutoRoute(
    path: "social",
    page: SocialShellRoute.page,
    children: [
      AutoRoute(
        path: "",
        initial: true,
        page: FriendsListFragmentRoute.page,
      )
    ],
  );

  final _chatsRoutes = AutoRoute(
    path: "chats",
    page: ChatShellRoute.page,
    children: [
      AutoRoute(
        path: "",
        initial: true,
        page: ChatListFragmentRoute.page,
      ),
      AutoRoute(
        path: "chat",
        page: ChatFragmentRoute.page,
      ),
    ],
  );

  final _profileRoutes = AutoRoute(
    path: "profile",
    initial: true,
    page: ProfileFragmentRoute.page,
  );
}
