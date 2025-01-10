import 'package:auto_route/auto_route.dart';
import 'package:flutter_mobile/src/app_router/app_router.gr.dart';
import 'package:flutter_mobile/src/features/auth/service/auth_service.dart';

class AuthGuard extends AutoRouteGuard {
  final AuthService _authService;

  AuthGuard({
    required AuthService authService,
  }) : _authService = authService;

  @override
  void onNavigation(
    NavigationResolver resolver,
    StackRouter router,
  ) {
    print(router.current.path);
    print(router.current.path);
    print(router.current.path);
    print(_authService.isAuth.value);

    if (_authService.isAuth.value) {
      resolver.next(true);
    } else {
      // How redirect to auth by uri?
      resolver.redirect(const AuthRoute());
    }
  }
}

class UnauthGuard extends AutoRouteGuard {
  final AuthService _authService;

  UnauthGuard({
    required AuthService authService,
  }) : _authService = authService;

  @override
  void onNavigation(
    NavigationResolver resolver,
    StackRouter router,
  ) {
    print(router.current.path);
    print(router.current.path);
    print(router.current.path);
    print(_authService.isAuth.value);

    if (!_authService.isAuth.value) {
      resolver.next(true);
    } else {
      // How redire
      resolver.redirect(
        const AppNavRoute(),
        replace: true,
      );
    }
  }
}
