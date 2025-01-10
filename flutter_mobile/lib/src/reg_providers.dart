import 'package:flutter/material.dart';
import 'package:flutter_mobile/src/features/auth/service/auth_service.dart';
import 'package:provider/provider.dart';

class RegProviders extends StatelessWidget {
  final Widget _app;

  final AuthService _authService;

  const RegProviders({
    super.key,
    required Widget app,
    required AuthService authService,
  })  : _app = app,
        _authService = authService;

  @override
  Widget build(BuildContext context) {
    return MultiProvider(
      providers: [
        Provider<AuthService>(create: (_) => _authService),
      ],
      child: _app,
    );
  }
}
