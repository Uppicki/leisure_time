import 'package:flutter/material.dart';
import 'package:flutter_mobile/src/app_router/app_router.dart';
import 'package:flutter_mobile/src/features/auth/service/auth_service.dart';
import 'package:flutter_mobile/src/screens/app_nav_screen.dart';
import 'package:flutter_mobile/src/screens/auth_screen.dart';
import 'package:flutter_mobile/src/screens/content_screen.dart';
import 'package:provider/provider.dart';

/// The Widget that configures your application.
class MyApp extends StatelessWidget {
  final AppRouter _router;

  const MyApp({
    super.key,
    required AppRouter router,
  }) : _router = router;

  @override
  Widget build(BuildContext context) {
    final _authService = context.read<AuthService>();

    return MaterialApp.router(
      theme: ThemeData(
        useMaterial3: true,
      ),
      debugShowCheckedModeBanner: false,
      routerConfig: _router.config(
        reevaluateListenable: _authService.isAuth,
      ),
    );
  }
}
