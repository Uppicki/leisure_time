import 'package:flutter/material.dart';
import 'package:flutter_mobile/src/app_router/app_router.dart';
import 'package:flutter_mobile/src/features/auth/service/auth_service.dart';
import 'package:flutter_mobile/src/reg_providers.dart';

import 'src/app.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized();

  final authService = AuthService();

  final router = AppRouter(
    authService: authService,
  );

  Widget app = MyApp(
    router: router,
  );

  app = RegProviders(
    app: app,
    authService: authService,
  );

  runApp(app);
}
